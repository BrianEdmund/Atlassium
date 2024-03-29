// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// BrunoCoinClient is the client API for BrunoCoin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BrunoCoinClient interface {
	ForwardTransaction(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*Empty, error)
	ForwardBlock(ctx context.Context, in *Block, opts ...grpc.CallOption) (*Empty, error)
	// Establishes a one way connection to a node (may be reciprocated)
	Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*Empty, error)
	// Gets maximum 500 blocks past block with top hash
	GetBlocks(ctx context.Context, in *GetBlocksRequest, opts ...grpc.CallOption) (*GetBlocksResponse, error)
	// Get a single block
	GetData(ctx context.Context, in *GetDataRequest, opts ...grpc.CallOption) (*GetDataResponse, error)
	// Sends know addresses to neighbors, forwarded from node to node
	SendAddresses(ctx context.Context, in *Addresses, opts ...grpc.CallOption) (*Empty, error)
	// Gets neighbor addresses from node (can be multicast with static addr_me)
	GetAddresses(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Addresses, error)
}

type brunoCoinClient struct {
	cc grpc.ClientConnInterface
}

func NewBrunoCoinClient(cc grpc.ClientConnInterface) BrunoCoinClient {
	return &brunoCoinClient{cc}
}

func (c *brunoCoinClient) ForwardTransaction(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/BrunoCoin/ForwardTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brunoCoinClient) ForwardBlock(ctx context.Context, in *Block, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/BrunoCoin/ForwardBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brunoCoinClient) Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/BrunoCoin/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brunoCoinClient) GetBlocks(ctx context.Context, in *GetBlocksRequest, opts ...grpc.CallOption) (*GetBlocksResponse, error) {
	out := new(GetBlocksResponse)
	err := c.cc.Invoke(ctx, "/BrunoCoin/GetBlocks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brunoCoinClient) GetData(ctx context.Context, in *GetDataRequest, opts ...grpc.CallOption) (*GetDataResponse, error) {
	out := new(GetDataResponse)
	err := c.cc.Invoke(ctx, "/BrunoCoin/GetData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brunoCoinClient) SendAddresses(ctx context.Context, in *Addresses, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/BrunoCoin/SendAddresses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brunoCoinClient) GetAddresses(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Addresses, error) {
	out := new(Addresses)
	err := c.cc.Invoke(ctx, "/BrunoCoin/GetAddresses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BrunoCoinServer is the server API for BrunoCoin service.
// All implementations must embed UnimplementedBrunoCoinServer
// for forward compatibility
type BrunoCoinServer interface {
	ForwardTransaction(context.Context, *Transaction) (*Empty, error)
	ForwardBlock(context.Context, *Block) (*Empty, error)
	// Establishes a one way connection to a node (may be reciprocated)
	Version(context.Context, *VersionRequest) (*Empty, error)
	// Gets maximum 500 blocks past block with top hash
	GetBlocks(context.Context, *GetBlocksRequest) (*GetBlocksResponse, error)
	// Get a single block
	GetData(context.Context, *GetDataRequest) (*GetDataResponse, error)
	// Sends know addresses to neighbors, forwarded from node to node
	SendAddresses(context.Context, *Addresses) (*Empty, error)
	// Gets neighbor addresses from node (can be multicast with static addr_me)
	GetAddresses(context.Context, *Empty) (*Addresses, error)
	mustEmbedUnimplementedBrunoCoinServer()
}

// UnimplementedBrunoCoinServer must be embedded to have forward compatible implementations.
type UnimplementedBrunoCoinServer struct {
}

func (UnimplementedBrunoCoinServer) ForwardTransaction(context.Context, *Transaction) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForwardTransaction not implemented")
}
func (UnimplementedBrunoCoinServer) ForwardBlock(context.Context, *Block) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForwardBlock not implemented")
}
func (UnimplementedBrunoCoinServer) Version(context.Context, *VersionRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedBrunoCoinServer) GetBlocks(context.Context, *GetBlocksRequest) (*GetBlocksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlocks not implemented")
}
func (UnimplementedBrunoCoinServer) GetData(context.Context, *GetDataRequest) (*GetDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetData not implemented")
}
func (UnimplementedBrunoCoinServer) SendAddresses(context.Context, *Addresses) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendAddresses not implemented")
}
func (UnimplementedBrunoCoinServer) GetAddresses(context.Context, *Empty) (*Addresses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAddresses not implemented")
}
func (UnimplementedBrunoCoinServer) mustEmbedUnimplementedBrunoCoinServer() {}

// UnsafeBrunoCoinServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BrunoCoinServer will
// result in compilation errors.
type UnsafeBrunoCoinServer interface {
	mustEmbedUnimplementedBrunoCoinServer()
}

func RegisterBrunoCoinServer(s grpc.ServiceRegistrar, srv BrunoCoinServer) {
	s.RegisterService(&BrunoCoin_ServiceDesc, srv)
}

func _BrunoCoin_ForwardTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrunoCoinServer).ForwardTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BrunoCoin/ForwardTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrunoCoinServer).ForwardTransaction(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _BrunoCoin_ForwardBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Block)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrunoCoinServer).ForwardBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BrunoCoin/ForwardBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrunoCoinServer).ForwardBlock(ctx, req.(*Block))
	}
	return interceptor(ctx, in, info, handler)
}

func _BrunoCoin_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrunoCoinServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BrunoCoin/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrunoCoinServer).Version(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BrunoCoin_GetBlocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlocksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrunoCoinServer).GetBlocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BrunoCoin/GetBlocks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrunoCoinServer).GetBlocks(ctx, req.(*GetBlocksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BrunoCoin_GetData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrunoCoinServer).GetData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BrunoCoin/GetData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrunoCoinServer).GetData(ctx, req.(*GetDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BrunoCoin_SendAddresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Addresses)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrunoCoinServer).SendAddresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BrunoCoin/SendAddresses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrunoCoinServer).SendAddresses(ctx, req.(*Addresses))
	}
	return interceptor(ctx, in, info, handler)
}

func _BrunoCoin_GetAddresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrunoCoinServer).GetAddresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BrunoCoin/GetAddresses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrunoCoinServer).GetAddresses(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// BrunoCoin_ServiceDesc is the grpc.ServiceDesc for BrunoCoin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BrunoCoin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BrunoCoin",
	HandlerType: (*BrunoCoinServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ForwardTransaction",
			Handler:    _BrunoCoin_ForwardTransaction_Handler,
		},
		{
			MethodName: "ForwardBlock",
			Handler:    _BrunoCoin_ForwardBlock_Handler,
		},
		{
			MethodName: "Version",
			Handler:    _BrunoCoin_Version_Handler,
		},
		{
			MethodName: "GetBlocks",
			Handler:    _BrunoCoin_GetBlocks_Handler,
		},
		{
			MethodName: "GetData",
			Handler:    _BrunoCoin_GetData_Handler,
		},
		{
			MethodName: "SendAddresses",
			Handler:    _BrunoCoin_SendAddresses_Handler,
		},
		{
			MethodName: "GetAddresses",
			Handler:    _BrunoCoin_GetAddresses_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "advancedcoin.proto",
}
