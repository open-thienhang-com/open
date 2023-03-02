import sys
import os
sys.path.append(os.path.join(os.path.dirname(__file__), ""))


from grpc_service_compile import hello_service_pb2, hello_service_pb2_grpc
from concurrent import futures
import logging
import grpc


# print(dir(hello_service_pb2))
# print(dir(hello_service_pb2_grpc))


class Hello(hello_service_pb2_grpc.HelloService):

    def GetReply(self, request, context):
        return hello_service_pb2.HelloReply(message='Hello, %s!' % request.name)


def serve():
    port = '50051'
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    hello_service_pb2_grpc.add_HelloServiceServicer_to_server(Hello(), server)
    server.add_insecure_port('[::]:50051')
    server.start()

    print("Server started, listening on " + port)
    server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig()
    serve()
