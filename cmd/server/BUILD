load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["bazel_bes.go"],
    importpath = "github.com/smukherj1/bazel_bes/cmd/server",
    visibility = ["//visibility:private"],
    deps = [
        "//proto:build_event_stream_go_proto",
        "@com_github_golang_protobuf//ptypes:go_default_library_gen",
        "@com_github_pkg_errors//:go_default_library",
        "@go_googleapis//google/devtools/build/v1:build_go_proto",
        "@io_bazel_rules_go//proto/wkt:empty_go_proto",
        "@org_golang_google_grpc//:go_default_library",
    ],
)

go_binary(
    name = "server",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
