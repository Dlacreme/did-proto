// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: demail.proto

package demail_proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DemailClient is the client API for Demail service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DemailClient interface {
	PushLog(ctx context.Context, in *Log, opts ...grpc.CallOption) (*Status, error)
}

type demailClient struct {
	cc grpc.ClientConnInterface
}

func NewDemailClient(cc grpc.ClientConnInterface) DemailClient {
	return &demailClient{cc}
}

func (c *demailClient) PushLog(ctx context.Context, in *Log, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/Demail/PushLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DemailServer is the server API for Demail service.
// All implementations must embed UnimplementedDemailServer
// for forward compatibility
type DemailServer interface {
	PushLog(context.Context, *Log) (*Status, error)
	mustEmbedUnimplementedDemailServer()
}

// UnimplementedDemailServer must be embedded to have forward compatible implementations.
type UnimplementedDemailServer struct {
}

func (UnimplementedDemailServer) PushLog(context.Context, *Log) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushLog not implemented")
}
func (UnimplementedDemailServer) mustEmbedUnimplementedDemailServer() {}

// UnsafeDemailServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DemailServer will
// result in compilation errors.
type UnsafeDemailServer interface {
	mustEmbedUnimplementedDemailServer()
}

func RegisterDemailServer(s grpc.ServiceRegistrar, srv DemailServer) {
	s.RegisterService(&Demail_ServiceDesc, srv)
}

func _Demail_PushLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Log)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemailServer).PushLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Demail/PushLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemailServer).PushLog(ctx, req.(*Log))
	}
	return interceptor(ctx, in, info, handler)
}

// Demail_ServiceDesc is the grpc.ServiceDesc for Demail service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Demail_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Demail",
	HandlerType: (*DemailServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PushLog",
			Handler:    _Demail_PushLog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demail.proto",
}
