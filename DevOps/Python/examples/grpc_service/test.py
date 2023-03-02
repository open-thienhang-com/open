
if __name__ == '__main__':
    # import grpc
    import sys
    import os
    print("Python version")
    print (sys.version)
    print("Version info.")
    print (sys.version_info)
    print("xin chao")
    print(sys.path)
    print('os.path.abspath(__file__) is:', os.path.abspath(__file__))
    print('os.path.dirname(os.path.abspath(__file__)) is:', os.path.dirname(os.path.abspath(__file__)))
    from importlib import metadata
    print(metadata.version("pip"))
    print(metadata.version("grpc"))
    # sys.path.append(os.path.join(os.path.dirname(__file__), ""))

    # from grpc_service_compile import hello_service_pb2, hello_service_pb2_grpc
    # from concurrent import futures
    # import logging
    # # import flask
    # # print("Flask ok")
    # import grpc
    # from grpc_service_compile import hello_service_pb2, hello_service_pb2_grpc