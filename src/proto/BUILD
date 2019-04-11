load("@build_stack_rules_proto//go:go_proto_library.bzl", "go_proto_library")

proto_library(
    name = "build_event_stream_proto",
    srcs = ["build_event_stream.proto"],
)


go_proto_library(
    name = "build_event_stream_go_proto",
    deps = [
        ":build_event_stream_proto",
    ],
    has_services = True,
    importpath = "bazel_bes/proto",
    visibility = ["//visibility:public"],
)

