load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("//bazel:api_build_system.bzl", "xds_proto_package")

licenses(["notice"])  # Apache 2

xds_proto_package()

go_library(
    name = "type",
    srcs = [
        "typed_struct.pb.go",
        "typed_struct.pb.validate.go",
    ],
    importpath = "github.com/aaomidi/xds/udpa/type/v1",
    visibility = ["//visibility:public"],
    deps = [
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
        "@org_golang_google_protobuf//types/known/anypb",
        "@org_golang_google_protobuf//types/known/structpb",
    ],
)
