workspace(name="bazel_bes")

load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")
# https://github.com/bazelbuild/bazel/blob/master/third_party/googleapis/BUILD.bazel#L220
git_repository(
  name = "org_pubref_rules_protobuf",
  remote = "https://github.com/pubref/rules_protobuf",
  tag = "v0.8.2",
)

load("@org_pubref_rules_protobuf//java:rules.bzl", "java_proto_repositories")
java_proto_repositories()

load("@org_pubref_rules_protobuf//cpp:rules.bzl", "cpp_proto_repositories")
cpp_proto_repositories()

load("@org_pubref_rules_protobuf//go:rules.bzl", "go_proto_repositories")
go_proto_repositories()

git_repository(
  name = "@org_googleapis_googleapis",
  remote = "https://github.com/googleapis/googleapis",
  commit = "9a02c5acecb43f38fae4fa52c6420f21c335b888",
)
