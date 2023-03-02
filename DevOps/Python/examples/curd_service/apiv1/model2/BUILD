package(default_visibility = ["//visibility:public"])
load("@rules_python//python:defs.bzl", "py_library")
load("@gpip//:requirements.bzl", "requirement")

py_library(
    name = "model2",
    srcs = ["__init__.py"],
    imports = [
        "__init__.py",
    ],
    deps = [
        requirement("flask"),
        requirement("Werkzeug"),
        requirement("flask_restx")
    ],
)

