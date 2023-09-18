// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: dcs/net/v0/net.proto

package net

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

// NetServiceClient is the client API for NetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NetServiceClient interface {
	// https://wiki.hoggitworld.com/view/DCS_func_send_chat_to
	SendChatTo(ctx context.Context, in *SendChatToRequest, opts ...grpc.CallOption) (*SendChatToResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_send_chat
	SendChat(ctx context.Context, in *SendChatRequest, opts ...grpc.CallOption) (*SendChatResponse, error)
	// returns a list of all connected players.
	// https://wiki.hoggitworld.com/view/DCS_func_get_player_info
	GetPlayers(ctx context.Context, in *GetPlayersRequest, opts ...grpc.CallOption) (*GetPlayersResponse, error)
	// Kick a specified player from the server with a message
	// https://wiki.hoggitworld.com/view/DCS_func_kick
	KickPlayer(ctx context.Context, in *KickPlayerRequest, opts ...grpc.CallOption) (*KickPlayerResponse, error)
	// Force a player into a slot / coalition.
	// To move the player back into spectators, use the following pseudo:
	// `ForcePlayerSlot({ player_id: ..., coalition: NEUTRAL, slot_id: "" })`
	ForcePlayerSlot(ctx context.Context, in *ForcePlayerSlotRequest, opts ...grpc.CallOption) (*ForcePlayerSlotResponse, error)
}

type netServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNetServiceClient(cc grpc.ClientConnInterface) NetServiceClient {
	return &netServiceClient{cc}
}

func (c *netServiceClient) SendChatTo(ctx context.Context, in *SendChatToRequest, opts ...grpc.CallOption) (*SendChatToResponse, error) {
	out := new(SendChatToResponse)
	err := c.cc.Invoke(ctx, "/dcs.net.v0.NetService/SendChatTo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *netServiceClient) SendChat(ctx context.Context, in *SendChatRequest, opts ...grpc.CallOption) (*SendChatResponse, error) {
	out := new(SendChatResponse)
	err := c.cc.Invoke(ctx, "/dcs.net.v0.NetService/SendChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *netServiceClient) GetPlayers(ctx context.Context, in *GetPlayersRequest, opts ...grpc.CallOption) (*GetPlayersResponse, error) {
	out := new(GetPlayersResponse)
	err := c.cc.Invoke(ctx, "/dcs.net.v0.NetService/GetPlayers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *netServiceClient) KickPlayer(ctx context.Context, in *KickPlayerRequest, opts ...grpc.CallOption) (*KickPlayerResponse, error) {
	out := new(KickPlayerResponse)
	err := c.cc.Invoke(ctx, "/dcs.net.v0.NetService/KickPlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *netServiceClient) ForcePlayerSlot(ctx context.Context, in *ForcePlayerSlotRequest, opts ...grpc.CallOption) (*ForcePlayerSlotResponse, error) {
	out := new(ForcePlayerSlotResponse)
	err := c.cc.Invoke(ctx, "/dcs.net.v0.NetService/ForcePlayerSlot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetServiceServer is the server API for NetService service.
// All implementations must embed UnimplementedNetServiceServer
// for forward compatibility
type NetServiceServer interface {
	// https://wiki.hoggitworld.com/view/DCS_func_send_chat_to
	SendChatTo(context.Context, *SendChatToRequest) (*SendChatToResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_send_chat
	SendChat(context.Context, *SendChatRequest) (*SendChatResponse, error)
	// returns a list of all connected players.
	// https://wiki.hoggitworld.com/view/DCS_func_get_player_info
	GetPlayers(context.Context, *GetPlayersRequest) (*GetPlayersResponse, error)
	// Kick a specified player from the server with a message
	// https://wiki.hoggitworld.com/view/DCS_func_kick
	KickPlayer(context.Context, *KickPlayerRequest) (*KickPlayerResponse, error)
	// Force a player into a slot / coalition.
	// To move the player back into spectators, use the following pseudo:
	// `ForcePlayerSlot({ player_id: ..., coalition: NEUTRAL, slot_id: "" })`
	ForcePlayerSlot(context.Context, *ForcePlayerSlotRequest) (*ForcePlayerSlotResponse, error)
	mustEmbedUnimplementedNetServiceServer()
}

// UnimplementedNetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNetServiceServer struct {
}

func (UnimplementedNetServiceServer) SendChatTo(context.Context, *SendChatToRequest) (*SendChatToResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendChatTo not implemented")
}
func (UnimplementedNetServiceServer) SendChat(context.Context, *SendChatRequest) (*SendChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendChat not implemented")
}
func (UnimplementedNetServiceServer) GetPlayers(context.Context, *GetPlayersRequest) (*GetPlayersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayers not implemented")
}
func (UnimplementedNetServiceServer) KickPlayer(context.Context, *KickPlayerRequest) (*KickPlayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KickPlayer not implemented")
}
func (UnimplementedNetServiceServer) ForcePlayerSlot(context.Context, *ForcePlayerSlotRequest) (*ForcePlayerSlotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForcePlayerSlot not implemented")
}
func (UnimplementedNetServiceServer) mustEmbedUnimplementedNetServiceServer() {}

// UnsafeNetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NetServiceServer will
// result in compilation errors.
type UnsafeNetServiceServer interface {
	mustEmbedUnimplementedNetServiceServer()
}

func RegisterNetServiceServer(s grpc.ServiceRegistrar, srv NetServiceServer) {
	s.RegisterService(&NetService_ServiceDesc, srv)
}

func _NetService_SendChatTo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendChatToRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetServiceServer).SendChatTo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.net.v0.NetService/SendChatTo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetServiceServer).SendChatTo(ctx, req.(*SendChatToRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetService_SendChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetServiceServer).SendChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.net.v0.NetService/SendChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetServiceServer).SendChat(ctx, req.(*SendChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetService_GetPlayers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlayersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetServiceServer).GetPlayers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.net.v0.NetService/GetPlayers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetServiceServer).GetPlayers(ctx, req.(*GetPlayersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetService_KickPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KickPlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetServiceServer).KickPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.net.v0.NetService/KickPlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetServiceServer).KickPlayer(ctx, req.(*KickPlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetService_ForcePlayerSlot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForcePlayerSlotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetServiceServer).ForcePlayerSlot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.net.v0.NetService/ForcePlayerSlot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetServiceServer).ForcePlayerSlot(ctx, req.(*ForcePlayerSlotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NetService_ServiceDesc is the grpc.ServiceDesc for NetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dcs.net.v0.NetService",
	HandlerType: (*NetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendChatTo",
			Handler:    _NetService_SendChatTo_Handler,
		},
		{
			MethodName: "SendChat",
			Handler:    _NetService_SendChat_Handler,
		},
		{
			MethodName: "GetPlayers",
			Handler:    _NetService_GetPlayers_Handler,
		},
		{
			MethodName: "KickPlayer",
			Handler:    _NetService_KickPlayer_Handler,
		},
		{
			MethodName: "ForcePlayerSlot",
			Handler:    _NetService_ForcePlayerSlot_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dcs/net/v0/net.proto",
}
