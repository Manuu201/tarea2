// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: proto/cliente.proto

package proto

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

const (
	ClienteService_AccionCliente_FullMethodName = "/grpc.ClienteService/AccionCliente"
)

// ClienteServiceClient is the client API for ClienteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClienteServiceClient interface {
	AccionCliente(ctx context.Context, in *ClienteRequest, opts ...grpc.CallOption) (*ClienteResponse, error)
}

type clienteServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClienteServiceClient(cc grpc.ClientConnInterface) ClienteServiceClient {
	return &clienteServiceClient{cc}
}

func (c *clienteServiceClient) AccionCliente(ctx context.Context, in *ClienteRequest, opts ...grpc.CallOption) (*ClienteResponse, error) {
	out := new(ClienteResponse)
	err := c.cc.Invoke(ctx, ClienteService_AccionCliente_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClienteServiceServer is the server API for ClienteService service.
// All implementations must embed UnimplementedClienteServiceServer
// for forward compatibility
type ClienteServiceServer interface {
	AccionCliente(context.Context, *ClienteRequest) (*ClienteResponse, error)
	mustEmbedUnimplementedClienteServiceServer()
}

// UnimplementedClienteServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClienteServiceServer struct {
}

func (UnimplementedClienteServiceServer) AccionCliente(context.Context, *ClienteRequest) (*ClienteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccionCliente not implemented")
}
func (UnimplementedClienteServiceServer) mustEmbedUnimplementedClienteServiceServer() {}

// UnsafeClienteServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClienteServiceServer will
// result in compilation errors.
type UnsafeClienteServiceServer interface {
	mustEmbedUnimplementedClienteServiceServer()
}

func RegisterClienteServiceServer(s grpc.ServiceRegistrar, srv ClienteServiceServer) {
	s.RegisterService(&ClienteService_ServiceDesc, srv)
}

func _ClienteService_AccionCliente_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClienteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClienteServiceServer).AccionCliente(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClienteService_AccionCliente_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClienteServiceServer).AccionCliente(ctx, req.(*ClienteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClienteService_ServiceDesc is the grpc.ServiceDesc for ClienteService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClienteService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.ClienteService",
	HandlerType: (*ClienteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AccionCliente",
			Handler:    _ClienteService_AccionCliente_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/cliente.proto",
}
