load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("//bazel:api_build_system.bzl", "xds_proto_package")

licenses(["notice"])  # Apache 2

xds_proto_package()

go_library(
    name = "annotations",
    srcs = [
        "migrate.pb.go",
        "migrate.pb.validate.go",
        "security.pb.go",
        "security.pb.validate.go",
        "sensitive.pb.go",
        "sensitive.pb.validate.go",
        "status.pb.go",
        "status.pb.validate.go",
        "versioning.pb.go",
        "versioning.pb.validate.go",
    ],
    importpath = "github.com/aaomidi/xds/udpa/annotations",
    visibility = ["//visibility:public"],
    deps = [
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
        "@org_golang_google_protobuf//types/descriptorpb",
        "@org_golang_google_protobuf//types/known/anypb",
    ],
)
