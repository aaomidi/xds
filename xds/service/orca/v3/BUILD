load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("//bazel:api_build_system.bzl", "xds_proto_package")

licenses(["notice"])  # Apache 2

xds_proto_package(
    has_services = True,
    deps = [
        "//xds/data/orca/v3:pkg",
    ],
)

go_library(
    name = "orca",
    srcs = [
        "orca.pb.go",
        "orca.pb.validate.go",
        "orca_grpc.pb.go",
    ],
    importpath = "github.com/aaomidi/xds/xds/service/orca/v3",
    visibility = ["//visibility:public"],
    deps = [
        "//xds/data/orca/v3:orca",
        "@org_golang_google_grpc//:grpc",
        "@org_golang_google_grpc//codes",
        "@org_golang_google_grpc//status",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
        "@org_golang_google_protobuf//types/known/anypb",
        "@org_golang_google_protobuf//types/known/durationpb",
    ],
)
