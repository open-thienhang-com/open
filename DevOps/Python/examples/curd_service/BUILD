package(default_visibility = ["//visibility:public"])

# load("@python//3.9:defs.bzl", py_binary_3_9 = "py_binary", py_test_3_9 = "py_test")
load("@gpip//:requirements.bzl", "requirement")
load("@io_bazel_rules_docker//python3:image.bzl", "py3_image")
load("@io_bazel_rules_docker//container:container.bzl", "container_image", "container_layer")

LIBS = [
    "//python/libs/glog",
    "//python/libs/adapter/http",
    "//python/examples/http_service/apiv1",
]

py_binary(
    name = "http_service",
    srcs = [
        "service.py",
    ],
    main = "service.py",
    imports= [
        "."
        # "apiv1/**",
        # "apiv2",
        # "apiv1/model1",
        # "apiv1/model2"
    ],
    deps = LIBS,
)

container_image(
    name="image",
    base = "@grpc_py_base//image", #TODO
    workdir="/app",
    symlinks={
        "/usr/bin/python": "/usr/local/bin/python3.8",  # To work as base for py3_image
        "/usr/bin/python3": "/usr/local/bin/python3.8"  # To work as base for py3_image
    },
)

py3_image(
    name = "http_service_deploy",
    srcs_version = "PY3ONLY",
    imports= ["."],
    srcs = ["service.py"],
    main = "service.py",
    # layers = LIBS,
    deps = LIBS,
    base=":image",
)

