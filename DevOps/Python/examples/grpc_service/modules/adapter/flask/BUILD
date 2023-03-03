package(default_visibility = ["//visibility:public"])
load("@rules_python//python:defs.bzl", "py_library")
load("@gpip//:requirements.bzl", "requirement")

py_library(
    name = "http",
    srcs = ["__init__.py"],
    imports = [
        "__init__.py",
    ],
    # visibility = ["@//tests:__pkg__"],
    deps = [
        "//python/libs/adapter/base",
        requirement("flask"),
        requirement("Werkzeug"),
        requirement("flask_restx")
    ],
)

