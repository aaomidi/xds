load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("//bazel:api_build_system.bzl", "xds_proto_package")

licenses(["notice"])  # Apache 2

xds_proto_package(
    has_services = True,
    deps = [
        "//udpa/data/orca/v1:pkg",
    ],
)

go_library(
    name = "orca",
    srcs = [
        "orca.pb.go",
        "orca.pb.validate.go",
        "orca_grpc.pb.go",
    ],
    importpath = "github.com/aaomidi/xds/udpa/service/orca/v1",
    visibility = ["//visibility:public"],
    deps = [
        "//udpa/data/orca/v1:orca",
        "@org_golang_google_grpc//:grpc",
        "@org_golang_google_grpc//codes",
        "@org_golang_google_grpc//status",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
        "@org_golang_google_protobuf//types/known/anypb",
        "@org_golang_google_protobuf//types/known/durationpb",
    ],
)
