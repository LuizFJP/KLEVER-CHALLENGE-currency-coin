// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: service.proto

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

// CurrencyCoinServiceClient is the client API for CurrencyCoinService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CurrencyCoinServiceClient interface {
	CreateCoin(ctx context.Context, in *CreateCoinRequest, opts ...grpc.CallOption) (*CoinResponse, error)
	ListCoins(ctx context.Context, in *ListCoinRequest, opts ...grpc.CallOption) (CurrencyCoinService_ListCoinsClient, error)
}

type currencyCoinServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCurrencyCoinServiceClient(cc grpc.ClientConnInterface) CurrencyCoinServiceClient {
	return &currencyCoinServiceClient{cc}
}

func (c *currencyCoinServiceClient) CreateCoin(ctx context.Context, in *CreateCoinRequest, opts ...grpc.CallOption) (*CoinResponse, error) {
	out := new(CoinResponse)
	err := c.cc.Invoke(ctx, "/service.CurrencyCoinService/createCoin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyCoinServiceClient) ListCoins(ctx context.Context, in *ListCoinRequest, opts ...grpc.CallOption) (CurrencyCoinService_ListCoinsClient, error) {
	stream, err := c.cc.NewStream(ctx, &CurrencyCoinService_ServiceDesc.Streams[0], "/service.CurrencyCoinService/ListCoins", opts...)
	if err != nil {
		return nil, err
	}
	x := &currencyCoinServiceListCoinsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CurrencyCoinService_ListCoinsClient interface {
	Recv() (*CoinResponse, error)
	grpc.ClientStream
}

type currencyCoinServiceListCoinsClient struct {
	grpc.ClientStream
}

func (x *currencyCoinServiceListCoinsClient) Recv() (*CoinResponse, error) {
	m := new(CoinResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CurrencyCoinServiceServer is the server API for CurrencyCoinService service.
// All implementations must embed UnimplementedCurrencyCoinServiceServer
// for forward compatibility
type CurrencyCoinServiceServer interface {
	CreateCoin(context.Context, *CreateCoinRequest) (*CoinResponse, error)
	ListCoins(*ListCoinRequest, CurrencyCoinService_ListCoinsServer) error
	mustEmbedUnimplementedCurrencyCoinServiceServer()
}

// UnimplementedCurrencyCoinServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCurrencyCoinServiceServer struct {
}

func (UnimplementedCurrencyCoinServiceServer) CreateCoin(context.Context, *CreateCoinRequest) (*CoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCoin not implemented")
}
func (UnimplementedCurrencyCoinServiceServer) ListCoins(*ListCoinRequest, CurrencyCoinService_ListCoinsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListCoins not implemented")
}
func (UnimplementedCurrencyCoinServiceServer) mustEmbedUnimplementedCurrencyCoinServiceServer() {}

// UnsafeCurrencyCoinServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CurrencyCoinServiceServer will
// result in compilation errors.
type UnsafeCurrencyCoinServiceServer interface {
	mustEmbedUnimplementedCurrencyCoinServiceServer()
}

func RegisterCurrencyCoinServiceServer(s grpc.ServiceRegistrar, srv CurrencyCoinServiceServer) {
	s.RegisterService(&CurrencyCoinService_ServiceDesc, srv)
}

func _CurrencyCoinService_CreateCoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyCoinServiceServer).CreateCoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.CurrencyCoinService/createCoin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyCoinServiceServer).CreateCoin(ctx, req.(*CreateCoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrencyCoinService_ListCoins_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListCoinRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CurrencyCoinServiceServer).ListCoins(m, &currencyCoinServiceListCoinsServer{stream})
}

type CurrencyCoinService_ListCoinsServer interface {
	Send(*CoinResponse) error
	grpc.ServerStream
}

type currencyCoinServiceListCoinsServer struct {
	grpc.ServerStream
}

func (x *currencyCoinServiceListCoinsServer) Send(m *CoinResponse) error {
	return x.ServerStream.SendMsg(m)
}

// CurrencyCoinService_ServiceDesc is the grpc.ServiceDesc for CurrencyCoinService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CurrencyCoinService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.CurrencyCoinService",
	HandlerType: (*CurrencyCoinServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createCoin",
			Handler:    _CurrencyCoinService_CreateCoin_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListCoins",
			Handler:       _CurrencyCoinService_ListCoins_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "service.proto",
}
