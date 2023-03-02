# load("@pypi//:requirements.bzl", "requirement")
# load("@rules_python//python:defs.bzl", "py_library")
package(default_visibility = ["//visibility:public"])
py_library(
    name = "env",
    srcs = ["__init__.py"],
    imports = [
        "__init__.py",
    ],
)

