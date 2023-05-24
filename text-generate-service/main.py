import grpc
import logging
import asyncio
from concurrent import futures

from protos import hello_pb2_grpc
from protos import hello_pb2


class HelloServicer(hello_pb2_grpc.HelloServicer):
    async def SayHello(
        self, request: hello_pb2.HelloRequest, context: grpc.aio.ServicerContext
    ) -> hello_pb2.HelloResponse:
        return hello_pb2.HelloResponse(
            message="hello %s from python!" % request.message
        )


async def serveGrpc() -> None:
    server = grpc.aio.server()
    hello_pb2_grpc.add_HelloServicer_to_server(HelloServicer(), server)
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
