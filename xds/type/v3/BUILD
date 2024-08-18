load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("//bazel:api_build_system.bzl", "xds_proto_package")

licenses(["notice"])  # Apache 2

xds_proto_package(
    deps = [
        "//xds/annotations/v3:pkg",
        "@com_google_googleapis//google/api/expr/v1alpha1:checked_proto",
        "@com_google_googleapis//google/api/expr/v1alpha1:syntax_proto",
        "@dev_cel_expr//proto/cel/expr:checked_proto",
        "@dev_cel_expr//proto/cel/expr:syntax_proto",
    ],
)

go_library(
    name = "expr",
    srcs = [
        "@dev_cel_expr2//:checked.pb.go",
        "@dev_cel_expr2//:eval.pb.go",
        "@dev_cel_expr2//:explain.pb.go",
        "@dev_cel_expr2//:syntax.pb.go",
        "@dev_cel_expr2//:value.pb.go",
    ],
    importpath = "cel.dev/expr",
    visibility = ["//visibility:public"],
    deps = [
        "@dev_cel_expr2//proto/cel/expr:google_rpc_status_go_proto",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
        "@org_golang_google_protobuf//types/known/anypb",
        "@org_golang_google_protobuf//types/known/durationpb",
        "@org_golang_google_protobuf//types/known/emptypb",
        "@org_golang_google_protobuf//types/known/structpb",
        "@org_golang_google_protobuf//types/known/timestamppb",
    ],
)

go_library(
    name = "type",
    srcs = [
        "cel.pb.go",
        "cel.pb.validate.go",
        "range.pb.go",
        "range.pb.validate.go",
        "typed_struct.pb.go",
        "typed_struct.pb.validate.go",
    ],
    importpath = "github.com/aaomidi/xds/xds/type/v3",
    visibility = ["//visibility:public"],
    deps = [
        ":expr",
        "//xds/annotations/v3:annotations",
        "@com_envoyproxy_protoc_gen_validate//validate:go_default_library",
        "@org_golang_google_genproto_googleapis_api//expr/v1alpha1",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
        "@org_golang_google_protobuf//types/known/anypb",
        "@org_golang_google_protobuf//types/known/structpb",
        "@org_golang_google_protobuf//types/known/wrapperspb",
    ],
)
