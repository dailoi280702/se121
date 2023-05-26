import grpc
import logging
import asyncio

# from concurrent import futures

from protos import text_generate_pb2
from protos import text_generate_pb2_grpc


# class HelloServicer(hello_pb2_grpc.HelloServicer):
#     async def SayHello(
#         self, request: hello_pb2.HelloRequest, context: grpc.aio.ServicerContext
#     ) -> hello_pb2.HelloResponse:
#         return hello_pb2.HelloResponse(
#             message="hello %s from python!" % request.message
#         )


class TextGenerateServicer:
    async def GenerateCarReview(
        self,
        request: text_generate_pb2.GenerateReviewReq,
        context: grpc.aio.ServicerContext,
    ) -> text_generate_pb2.ResString:
        return text_generate_pb2.ResString(text="fugg")
        # context.set_code(grpc.StatusCode.INTERNAL)
        # context.set_details("Text generate server error")

    async def GenerateBlogSummarization(
        self,
        request: text_generate_pb2.GenerateBlogSummarizationReq,
        context: grpc.aio.ServicerContext,
    ) -> text_generate_pb2.ResString:
        return text_generate_pb2.ResString(text="fugg")
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
    logging.basicConfig(level=logging.INFO)
    asyncio.get_event_loop().run_until_complete(serveGrpc())


if __name__ == "__main__":
    main()
