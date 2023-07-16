import grpc
import logging
import asyncio
import openai
import threading
import os
import random
from protos import text_generate_pb2
from protos import text_generate_pb2_grpc
from dotenv import load_dotenv

load_dotenv()


def generate_tldr(blog_title, blog_content):
    prompt = f"{blog_title}\n\n{blog_content}\n\nSummarization under 150 words:"
    response = openai.Completion.create(
        engine="text-davinci-003",
        prompt=prompt,
        max_tokens=200,
        temperature=0.5,
        top_p=1.0,
        frequency_penalty=0.0,
        presence_penalty=0.0,
        n=1,
        stop=None,
        echo=False,
    )
    tldr = (
        response.choices[0]
        .text.strip()
        .replace("Summarization under 150 words:", "")
        .strip()
    )
    return tldr


def chunk_and_call_tldr(blog_content, blog_title):
    chunks = []
    tokens = blog_content.split(" ")
    chunk_size = 4000
    for i in range(0, len(tokens), chunk_size):
        chunk = tokens[i:i + chunk_size]
        chunk = " ".join(chunk)
        chunks.append(chunk)

    tldrs = []
    threads = []
    for chunk in chunks:
        thread = threading.Thread(target=generate_tldr, args=(blog_title, chunk))
        threads.append(thread)
        thread.start()

    for thread in threads:
        thread.join()

    combined_tldr = " ".join(tldrs)
    if len(tldrs) > 1:
        tldr = generate_tldr(blog_title, combined_tldr)
        return tldr
    else:
        return combined_tldr


class TextGenerateServicer:
    async def GenerateCarReview(
        self,
        request: text_generate_pb2.GenerateReviewReq,
        context: grpc.aio.ServicerContext,
    ) -> text_generate_pb2.ResString:
        # Format the input prompt
        prompt = f"Write a review withot subject for the {request.brand} {request.series} {request.name}."

        info_list = []
        if request.fuelType != None:
            info_list.append(f"\n fuel type: {request.fuelType}")
        if request.transmission != None:
            info_list.append(f"\n transmission: {request.transmission}")
        if request.horsePower != None:
            info_list.append(f"\n horse powser: {request.horsePower}")
        if request.torque != None:
            info_list.append(f"\n torque: {request.torque}")

        # Randomly select and include an information in the prompt
        if info_list:
            random_info = random.choice(info_list)
            prompt += f" It has {random_info}."

        # Generate the review using OpenAI API
        response = openai.Completion.create(
            engine="text-davinci-003",
            prompt=prompt,
            max_tokens=100,
            n=1,
            stop=None,
            temperature=0.7,
        )

        # Extract the generated review from the API response
        generated_review = response["choices"][0]["text"].strip()

        return text_generate_pb2.ResString(text=generated_review)

        # context.set_code(grpc.StatusCode.INTERNAL)
        # context.set_details("Text generate server error")

    async def GenerateBlogSummarization(
        self,
        request: text_generate_pb2.GenerateBlogSummarizationReq,
        context: grpc.aio.ServicerContext,
    ) -> text_generate_pb2.ResString:
        generated_tldr = chunk_and_call_tldr(request.title, request.body)
        return text_generate_pb2.ResString(text=generated_tldr)
        # context.set_code(grpc.StatusCode.INTERNAL)
        # context.set_details("Text generate server error")
        # raise NotImplementedError("Method not implemented!")


async def serveGrpc() -> None:
    server = grpc.aio.server()
    # hello_pb2_grpc.add_HelloServicer_to_server(HelloServicer(), server)
    text_generate_pb2_grpc.add_TextGenerateServiceServicer_to_server(
        TextGenerateServicer(), server
    )
    listen_addr = "[::]:50051"
    server.add_insecure_port(listen_addr)
    logging.info("Starting server on %s", listen_addr)
    await server.start()
    await server.wait_for_termination()


def main():
    openai.api_key = os.getenv("OPENAI_API_KEY")
    logging.info(openai.api_key)
    logging.basicConfig(level=logging.INFO)
    asyncio.get_event_loop().run_until_complete(serveGrpc())


if __name__ == "__main__":
    main()
