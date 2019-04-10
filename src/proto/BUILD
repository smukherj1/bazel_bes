load("@build_stack_rules_proto//go:go_grpc_compile.bzl", "go_grpc_compile")
load("@build_stack_rules_proto//go:go_proto_compile.bzl", "go_proto_compile")

go_grpc_compile(
    name = "google_devtools_build_v1_publish_build_event_go_proto",
    has_services = True,
    #importpath = "github.com/example/project/src/proto",
    deps = [
        "@googleapis//:google_devtools_build_v1_publish_build_event_proto",
    ],
)

proto_library(
    name = "build_event_stream_proto",
    srcs = ["build_event_stream.proto"],
)


go_proto_compile(
    name = "build_event_stream_go_proto",
    deps = [":build_event_stream_proto"]
)