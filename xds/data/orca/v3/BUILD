load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("//bazel:api_build_system.bzl", "xds_proto_package")

licenses(["notice"])  # Apache 2

xds_proto_package()

go_library(
    name = "orca",
    srcs = [
        "orca_load_report.pb.go",
        "orca_load_report.pb.validate.go",
    ],
    importpath = "github.com/aaomidi/xds/xds/data/orca/v3",
    visibility = ["//visibility:public"],
    deps = [
        "@com_envoyproxy_protoc_gen_validate//validate:go_default_library",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
        "@org_golang_google_protobuf//types/known/anypb",
    ],
)
