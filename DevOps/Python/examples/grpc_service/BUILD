package(default_visibility = ["//visibility:public"])

# load("@python//3.9:defs.bzl", py_binary_3_9 = "py_binary", py_test_3_9 = "py_test")
load("@gpip//:requirements.bzl", "requirement")
load("@io_bazel_rules_docker//python3:image.bzl", "py3_image")
load("@io_bazel_rules_docker//container:container.bzl", "container_image", "container_layer")
# proto + grpc 

load("@rules_proto_grpc//python:defs.bzl", "python_grpc_compile")

proto_library(
    name = "hello_service_proto",
    srcs = ["hello_service.proto"],
    strip_import_prefix="",
    deps = [
        "@com_google_protobuf//:any_proto",
    ],
)

python_grpc_compile(
    name = "grpc_service_compile",
    protos = [":hello_service_proto"],
)

py_library(
    name = "grpc_service_lib",
    srcs = [
        ":grpc_service_compile",
    ],
    imports = [
        "grpc_service_compile",
    ],
    deps = [
        requirement("grpcio"),
        requirement("protobuf"),
    ]
)

LIBS = [
    ":grpc_service_lib",
    requirement("grpcio"),
    requirement("protobuf"),
]

py_binary(
    name = "grpc_service",
    srcs = ["service.py"],
    main = "service.py",
    # data = [":grpc_service_lib"],
    deps = LIBS
)

container_image(
    name= "image",
    # visibility=["//visibility:public"],
    base = "@py_base//image",
    workdir="/app",
    # symlinks={
    #     "/usr/bin/python": "/usr/local/bin/python3.8",  # To work as base for py3_image
    #     "/usr/bin/python3": "/usr/local/bin/python3.8"  # To work as base for py3_image
    # },
    
)

py3_image(
    name = "grpc_service_deploy",
    srcs_version = "PY3ONLY",
    imports= ["."],
    srcs = [
        "test.py",
    ],
    main = "test.py",
    deps = LIBS,
    base=":image",
)
