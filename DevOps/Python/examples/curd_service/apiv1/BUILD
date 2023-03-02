package(default_visibility = ["//visibility:public"])
load("@rules_python//python:defs.bzl", "py_library")
load("@gpip//:requirements.bzl", "requirement")

py_library(
    name = "apiv1",
    srcs = ["__init__.py"],
    imports = [
        "__init__.py",
    ],
    deps = [
        "//python/examples/http_service/apiv1/model1",
        "//python/examples/http_service/apiv1/model2",
        requirement("flask"),
        requirement("Werkzeug"),
        requirement("flask_restx")
    ],
)

