load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("//bazel:api_build_system.bzl", "xds_proto_package")

licenses(["notice"])  # Apache 2

xds_proto_package(deps = ["//xds/annotations/v3:pkg"])

go_library(
    name = "core",
    srcs = [
        "authority.pb.go",
        "authority.pb.validate.go",
        "cidr.pb.go",
        "cidr.pb.validate.go",
        "collection_entry.pb.go",
        "collection_entry.pb.validate.go",
        "context_params.pb.go",
        "context_params.pb.validate.go",
        "extension.pb.go",
        "extension.pb.validate.go",
        "resource.pb.go",
        "resource.pb.validate.go",
        "resource_locator.pb.go",
        "resource_locator.pb.validate.go",
        "resource_name.pb.go",
        "resource_name.pb.validate.go",
    ],
    importpath = "github.com/aaomidi/xds/xds/core/v3",
    visibility = ["//visibility:public"],
    deps = [
        "//xds/annotations/v3:annotations",
        "@com_envoyproxy_protoc_gen_validate//validate:go_default_library",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
        "@org_golang_google_protobuf//types/known/anypb",
        "@org_golang_google_protobuf//types/known/wrapperspb",
    ],
)
