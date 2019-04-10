load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["bazel_bes.go"],
    importpath = "github.com/example/project/src/server",
    visibility = ["//visibility:private"],
    deps = [
        "@go_googleapis//google/devtools/build/v1:build_go_proto",
        "@io_bazel_rules_go//proto/wkt:empty_go_proto",
        "@org_golang_google_grpc//:go_default_library",
        "//src/proto:build_event_stream_go_proto",
    ],
)

go_binary(
    name = "server",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
