load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("//bazel:api_build_system.bzl", "xds_proto_package")

licenses(["notice"])  # Apache 2

xds_proto_package(
    deps = [
        "//xds/annotations/v3:pkg",
        "//xds/core/v3:pkg",
        "//xds/type/v3:pkg",
    ],
)

go_library(
    name = "matcher",
    srcs = [
        "cel.pb.go",
        "cel.pb.validate.go",
        "domain.pb.go",
        "domain.pb.validate.go",
        "http_inputs.pb.go",
        "http_inputs.pb.validate.go",
        "ip.pb.go",
        "ip.pb.validate.go",
        "matcher.pb.go",
        "matcher.pb.validate.go",
        "range.pb.go",
        "range.pb.validate.go",
        "regex.pb.go",
        "regex.pb.validate.go",
        "string.pb.go",
        "string.pb.validate.go",
    ],
    importpath = "github.com/aaomidi/xds/xds/type/matcher/v3",
    visibility = ["//visibility:public"],
    deps = [
        "//xds/annotations/v3:annotations",
        "//xds/core/v3:core",
        "//xds/type/v3:type",
        "@com_envoyproxy_protoc_gen_validate//validate:go_default_library",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
        "@org_golang_google_protobuf//types/known/anypb",
    ],
)
