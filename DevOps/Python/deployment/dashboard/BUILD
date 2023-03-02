package(default_visibility = ["//visibility:public"])
load("@python//3.9:defs.bzl", py_binary_3_9 = "py_binary", py_test_3_9 = "py_test")
load("@io_bazel_rules_docker//python3:image.bzl", "py3_image")
load("@io_bazel_rules_docker//container:container.bzl", "container_image", "container_layer")

LIBS = [
    "//python/libs/glog",
    "//python/libs/adapter/http",
    "//python/libs/checker",
]

py_binary_3_9(
    name = "dashboard",
    srcs = [
        "service.py",
    ],
    main = "service.py",
    imports= [
        ".",
    ],
    deps = LIBS,
)

py3_image(
    name = "image",
    srcs_version = "PY3ONLY",
    imports= ["."],
    srcs = [
        "service.py",
    ],
    main = "service.py",
    # layers = LIBS,
    deps = LIBS,
    base = "@py_base//image",
)

container_image(
    name="test",
    base = ":image",
    workdir="/app",
)