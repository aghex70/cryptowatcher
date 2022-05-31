// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package providers

import (
	context "context"
	"gapi-agp/internal/core/usecases/pb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FetcherServiceClient is the client API for FetcherService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FetcherServiceClient interface {
	FetchTrades(ctx context.Context, in *pb.FetchTradesRequest, opts ...grpc.CallOption) (*pb.FetchTradesResponse, error)
	StopFetchTrades(ctx context.Context, in *pb.Empty, opts ...grpc.CallOption) (*pb.StopFetchTradesResponse, error)
}

type fetcherServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFetcherServiceClient(cc grpc.ClientConnInterface) FetcherServiceClient {
	return &fetcherServiceClient{cc}
}

func (c *fetcherServiceClient) FetchTrades(ctx context.Context, in *pb.FetchTradesRequest, opts ...grpc.CallOption) (*pb.FetchTradesResponse, error) {
	out := new(pb.FetchTradesResponse)
	err := c.cc.Invoke(ctx, "/notcoolj.transaction.v1.FetcherService/FetchTrades", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fetcherServiceClient) StopFetchTrades(ctx context.Context, in *pb.Empty, opts ...grpc.CallOption) (*pb.StopFetchTradesResponse, error) {
	out := new(pb.StopFetchTradesResponse)
	err := c.cc.Invoke(ctx, "/notcoolj.transaction.v1.FetcherService/StopFetchTrades", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FetcherServiceServer is the server API for FetcherService service.
// All implementations must embed UnimplementedFetcherServiceServer
// for forward compatibility
type FetcherServiceServer interface {
	FetchTrades(context.Context, *pb.FetchTradesRequest) (*pb.FetchTradesResponse, error)
	StopFetchTrades(context.Context, *pb.Empty) (*pb.StopFetchTradesResponse, error)
	mustEmbedUnimplementedFetcherServiceServer()
}

// UnimplementedFetcherServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFetcherServiceServer struct {
}

func (UnimplementedFetcherServiceServer) FetchTrades(context.Context, *pb.FetchTradesRequest) (*pb.FetchTradesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchTrades not implemented")
}
func (UnimplementedFetcherServiceServer) StopFetchTrades(context.Context, *pb.Empty) (*pb.StopFetchTradesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopFetchTrades not implemented")
}
func (UnimplementedFetcherServiceServer) mustEmbedUnimplementedFetcherServiceServer() {}

// UnsafeFetcherServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FetcherServiceServer will
// result in compilation errors.
type UnsafeFetcherServiceServer interface {
	mustEmbedUnimplementedFetcherServiceServer()
}

func RegisterFetcherServiceServer(s grpc.ServiceRegistrar, srv FetcherServiceServer) {
	s.RegisterService(&FetcherService_ServiceDesc, srv)
}

func _FetcherService_FetchTrades_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.FetchTradesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FetcherServiceServer).FetchTrades(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notcoolj.transaction.v1.FetcherService/FetchTrades",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FetcherServiceServer).FetchTrades(ctx, req.(*pb.FetchTradesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FetcherService_StopFetchTrades_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FetcherServiceServer).StopFetchTrades(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notcoolj.transaction.v1.FetcherService/StopFetchTrades",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FetcherServiceServer).StopFetchTrades(ctx, req.(*pb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// FetcherService_ServiceDesc is the grpc.ServiceDesc for FetcherService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FetcherService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notcoolj.transaction.v1.FetcherService",
	HandlerType: (*FetcherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchTrades",
			Handler:    _FetcherService_FetchTrades_Handler,
		},
		{
			MethodName: "StopFetchTrades",
			Handler:    _FetcherService_StopFetchTrades_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "routeguide/fetcher.proto",
}
