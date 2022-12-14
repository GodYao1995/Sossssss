// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: proto/service_side_stream.proto

package serviceside

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

// ServiceSideStreamClient is the client API for ServiceSideStream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceSideStreamClient interface {
	GetStreamValueFromService(ctx context.Context, in *ClientRequest, opts ...grpc.CallOption) (ServiceSideStream_GetStreamValueFromServiceClient, error)
}

type serviceSideStreamClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceSideStreamClient(cc grpc.ClientConnInterface) ServiceSideStreamClient {
	return &serviceSideStreamClient{cc}
}

func (c *serviceSideStreamClient) GetStreamValueFromService(ctx context.Context, in *ClientRequest, opts ...grpc.CallOption) (ServiceSideStream_GetStreamValueFromServiceClient, error) {
	stream, err := c.cc.NewStream(ctx, &ServiceSideStream_ServiceDesc.Streams[0], "/stream.ServiceSideStream/GetStreamValueFromService", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceSideStreamGetStreamValueFromServiceClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ServiceSideStream_GetStreamValueFromServiceClient interface {
	Recv() (*ServiceReply, error)
	grpc.ClientStream
}

type serviceSideStreamGetStreamValueFromServiceClient struct {
	grpc.ClientStream
}

func (x *serviceSideStreamGetStreamValueFromServiceClient) Recv() (*ServiceReply, error) {
	m := new(ServiceReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServiceSideStreamServer is the server API for ServiceSideStream service.
// All implementations must embed UnimplementedServiceSideStreamServer
// for forward compatibility
type ServiceSideStreamServer interface {
	GetStreamValueFromService(*ClientRequest, ServiceSideStream_GetStreamValueFromServiceServer) error
	mustEmbedUnimplementedServiceSideStreamServer()
}

// UnimplementedServiceSideStreamServer must be embedded to have forward compatible implementations.
type UnimplementedServiceSideStreamServer struct {
}

func (UnimplementedServiceSideStreamServer) GetStreamValueFromService(*ClientRequest, ServiceSideStream_GetStreamValueFromServiceServer) error {
	return status.Errorf(codes.Unimplemented, "method GetStreamValueFromService not implemented")
}
func (UnimplementedServiceSideStreamServer) mustEmbedUnimplementedServiceSideStreamServer() {}

// UnsafeServiceSideStreamServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceSideStreamServer will
// result in compilation errors.
type UnsafeServiceSideStreamServer interface {
	mustEmbedUnimplementedServiceSideStreamServer()
}

func RegisterServiceSideStreamServer(s grpc.ServiceRegistrar, srv ServiceSideStreamServer) {
	s.RegisterService(&ServiceSideStream_ServiceDesc, srv)
}

func _ServiceSideStream_GetStreamValueFromService_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ClientRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServiceSideStreamServer).GetStreamValueFromService(m, &serviceSideStreamGetStreamValueFromServiceServer{stream})
}

type ServiceSideStream_GetStreamValueFromServiceServer interface {
	Send(*ServiceReply) error
	grpc.ServerStream
}

type serviceSideStreamGetStreamValueFromServiceServer struct {
	grpc.ServerStream
}

func (x *serviceSideStreamGetStreamValueFromServiceServer) Send(m *ServiceReply) error {
	return x.ServerStream.SendMsg(m)
}

// ServiceSideStream_ServiceDesc is the grpc.ServiceDesc for ServiceSideStream service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceSideStream_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stream.ServiceSideStream",
	HandlerType: (*ServiceSideStreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetStreamValueFromService",
			Handler:       _ServiceSideStream_GetStreamValueFromService_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/service_side_stream.proto",
}
