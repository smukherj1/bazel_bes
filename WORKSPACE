workspace(name="bazel_bes")

# https://github.com/bazelbuild/bazel/blob/master/third_party/googleapis/BUILD.bazel#L220

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "build_stack_rules_proto",
    urls = ["https://github.com/stackb/rules_proto/archive/609362dd9b08110b7a95bfa26b5e3aac3cd06905.tar.gz"],
    #sha256 = "5474d1b83e24ec1a6db371033a27ff7aff412f2b23abba86fedd902330b61ee6",
    strip_prefix = "rules_proto-609362dd9b08110b7a95bfa26b5e3aac3cd06905",
)

http_archive(
    name = "googleapis",
    sha256 = "7b6ea252f0b8fb5cd722f45feb83e115b689909bbb6a393a873b6cbad4ceae1d",
    url = "https://github.com/googleapis/googleapis/archive/143084a2624b6591ee1f9d23e7f5241856642f4d.zip",
    strip_prefix = "googleapis-143084a2624b6591ee1f9d23e7f5241856642f4d",
    build_file = "@bazel_bes//:BUILD.googleapis",
)

http_archive(
    name = "bazel_gazelle",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.17.0/bazel-gazelle-0.17.0.tar.gz"],
    sha256 = "3c681998538231a2d24d0c07ed5a7658cb72bfb5fd4bf9911157c0e9ac6a2687",
)

load("@build_stack_rules_proto//go:deps.bzl", "go_proto_compile")
go_proto_compile()

load("@io_bazel_rules_go//go:def.bzl", "go_register_toolchains", "go_rules_dependencies")
go_rules_dependencies()
go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
gazelle_dependencies()
