package(default_visibility = ["//visibility:public"])

load("@build_stack_rules_proto//go:go_proto_library.bzl", "go_proto_library")

proto_library(
    name = "option_filters_proto",
    srcs = ["option_filters.proto"],
)

proto_library(
    name = "command_line_proto",
    srcs = ["command_line.proto"],
    deps = [":option_filters_proto"],
)

proto_library(
    name = "invocation_policy_proto",
    srcs = ["invocation_policy.proto"],
)

proto_library(
    name = "build_event_stream_proto",
    srcs = ["build_event_stream.proto"],
    deps = [
        ":invocation_policy_proto",
        ":command_line_proto",
    ],
)

go_proto_library(
    name = "build_event_stream_go_proto",
    deps = [
        ":build_event_stream_proto",
    ],
    importpath = "github.com/smukherj1/bazel_bes/proto",
    visibility = ["//visibility:public"],
)
