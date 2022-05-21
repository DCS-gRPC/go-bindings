// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: dcs/coalition/v0/coalition.proto

package coalition

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

// CoalitionServiceClient is the client API for CoalitionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoalitionServiceClient interface {
	// https://wiki.hoggitworld.com/view/DCS_func_addGroup
	AddGroup(ctx context.Context, in *AddGroupRequest, opts ...grpc.CallOption) (*AddGroupResponse, error)
	// Focussed on statics (linked statics - see `AddLinkedStatic`)
	// https://wiki.hoggitworld.com/view/DCS_func_addStaticObject
	AddStaticObject(ctx context.Context, in *AddStaticObjectRequest, opts ...grpc.CallOption) (*AddStaticObjectResponse, error)
	// Focussed on properties relevant to linked static objects
	// https://wiki.hoggitworld.com/view/DCS_func_addStaticObject
	AddLinkedStatic(ctx context.Context, in *AddLinkedStaticRequest, opts ...grpc.CallOption) (*AddLinkedStaticResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_getGroups
	GetGroups(ctx context.Context, in *GetGroupsRequest, opts ...grpc.CallOption) (*GetGroupsResponse, error)
	//
	// Get the Bullseye for the coalition
	//
	// This position is set at mission start and does not change for the duration
	// of the mission.
	//
	// See https://wiki.hoggitworld.com/view/DCS_func_getMainRefPoint for more
	// details
	GetBullseye(ctx context.Context, in *GetBullseyeRequest, opts ...grpc.CallOption) (*GetBullseyeResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_getPlayers
	GetPlayerUnits(ctx context.Context, in *GetPlayerUnitsRequest, opts ...grpc.CallOption) (*GetPlayerUnitsResponse, error)
}

type coalitionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCoalitionServiceClient(cc grpc.ClientConnInterface) CoalitionServiceClient {
	return &coalitionServiceClient{cc}
}

func (c *coalitionServiceClient) AddGroup(ctx context.Context, in *AddGroupRequest, opts ...grpc.CallOption) (*AddGroupResponse, error) {
	out := new(AddGroupResponse)
	err := c.cc.Invoke(ctx, "/dcs.coalition.v0.CoalitionService/AddGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coalitionServiceClient) AddStaticObject(ctx context.Context, in *AddStaticObjectRequest, opts ...grpc.CallOption) (*AddStaticObjectResponse, error) {
	out := new(AddStaticObjectResponse)
	err := c.cc.Invoke(ctx, "/dcs.coalition.v0.CoalitionService/AddStaticObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coalitionServiceClient) AddLinkedStatic(ctx context.Context, in *AddLinkedStaticRequest, opts ...grpc.CallOption) (*AddLinkedStaticResponse, error) {
	out := new(AddLinkedStaticResponse)
	err := c.cc.Invoke(ctx, "/dcs.coalition.v0.CoalitionService/AddLinkedStatic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coalitionServiceClient) GetGroups(ctx context.Context, in *GetGroupsRequest, opts ...grpc.CallOption) (*GetGroupsResponse, error) {
	out := new(GetGroupsResponse)
	err := c.cc.Invoke(ctx, "/dcs.coalition.v0.CoalitionService/GetGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coalitionServiceClient) GetBullseye(ctx context.Context, in *GetBullseyeRequest, opts ...grpc.CallOption) (*GetBullseyeResponse, error) {
	out := new(GetBullseyeResponse)
	err := c.cc.Invoke(ctx, "/dcs.coalition.v0.CoalitionService/GetBullseye", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coalitionServiceClient) GetPlayerUnits(ctx context.Context, in *GetPlayerUnitsRequest, opts ...grpc.CallOption) (*GetPlayerUnitsResponse, error) {
	out := new(GetPlayerUnitsResponse)
	err := c.cc.Invoke(ctx, "/dcs.coalition.v0.CoalitionService/GetPlayerUnits", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoalitionServiceServer is the server API for CoalitionService service.
// All implementations must embed UnimplementedCoalitionServiceServer
// for forward compatibility
type CoalitionServiceServer interface {
	// https://wiki.hoggitworld.com/view/DCS_func_addGroup
	AddGroup(context.Context, *AddGroupRequest) (*AddGroupResponse, error)
	// Focussed on statics (linked statics - see `AddLinkedStatic`)
	// https://wiki.hoggitworld.com/view/DCS_func_addStaticObject
	AddStaticObject(context.Context, *AddStaticObjectRequest) (*AddStaticObjectResponse, error)
	// Focussed on properties relevant to linked static objects
	// https://wiki.hoggitworld.com/view/DCS_func_addStaticObject
	AddLinkedStatic(context.Context, *AddLinkedStaticRequest) (*AddLinkedStaticResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_getGroups
	GetGroups(context.Context, *GetGroupsRequest) (*GetGroupsResponse, error)
	//
	// Get the Bullseye for the coalition
	//
	// This position is set at mission start and does not change for the duration
	// of the mission.
	//
	// See https://wiki.hoggitworld.com/view/DCS_func_getMainRefPoint for more
	// details
	GetBullseye(context.Context, *GetBullseyeRequest) (*GetBullseyeResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_getPlayers
	GetPlayerUnits(context.Context, *GetPlayerUnitsRequest) (*GetPlayerUnitsResponse, error)
	mustEmbedUnimplementedCoalitionServiceServer()
}

// UnimplementedCoalitionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCoalitionServiceServer struct {
}

func (UnimplementedCoalitionServiceServer) AddGroup(context.Context, *AddGroupRequest) (*AddGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddGroup not implemented")
}
func (UnimplementedCoalitionServiceServer) AddStaticObject(context.Context, *AddStaticObjectRequest) (*AddStaticObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddStaticObject not implemented")
}
func (UnimplementedCoalitionServiceServer) AddLinkedStatic(context.Context, *AddLinkedStaticRequest) (*AddLinkedStaticResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddLinkedStatic not implemented")
}
func (UnimplementedCoalitionServiceServer) GetGroups(context.Context, *GetGroupsRequest) (*GetGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroups not implemented")
}
func (UnimplementedCoalitionServiceServer) GetBullseye(context.Context, *GetBullseyeRequest) (*GetBullseyeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBullseye not implemented")
}
func (UnimplementedCoalitionServiceServer) GetPlayerUnits(context.Context, *GetPlayerUnitsRequest) (*GetPlayerUnitsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayerUnits not implemented")
}
func (UnimplementedCoalitionServiceServer) mustEmbedUnimplementedCoalitionServiceServer() {}

// UnsafeCoalitionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoalitionServiceServer will
// result in compilation errors.
type UnsafeCoalitionServiceServer interface {
	mustEmbedUnimplementedCoalitionServiceServer()
}

func RegisterCoalitionServiceServer(s grpc.ServiceRegistrar, srv CoalitionServiceServer) {
	s.RegisterService(&CoalitionService_ServiceDesc, srv)
}

func _CoalitionService_AddGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoalitionServiceServer).AddGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.coalition.v0.CoalitionService/AddGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoalitionServiceServer).AddGroup(ctx, req.(*AddGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoalitionService_AddStaticObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddStaticObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoalitionServiceServer).AddStaticObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.coalition.v0.CoalitionService/AddStaticObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoalitionServiceServer).AddStaticObject(ctx, req.(*AddStaticObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoalitionService_AddLinkedStatic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddLinkedStaticRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoalitionServiceServer).AddLinkedStatic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.coalition.v0.CoalitionService/AddLinkedStatic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoalitionServiceServer).AddLinkedStatic(ctx, req.(*AddLinkedStaticRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoalitionService_GetGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoalitionServiceServer).GetGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.coalition.v0.CoalitionService/GetGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoalitionServiceServer).GetGroups(ctx, req.(*GetGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoalitionService_GetBullseye_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBullseyeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoalitionServiceServer).GetBullseye(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.coalition.v0.CoalitionService/GetBullseye",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoalitionServiceServer).GetBullseye(ctx, req.(*GetBullseyeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoalitionService_GetPlayerUnits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlayerUnitsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoalitionServiceServer).GetPlayerUnits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.coalition.v0.CoalitionService/GetPlayerUnits",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoalitionServiceServer).GetPlayerUnits(ctx, req.(*GetPlayerUnitsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CoalitionService_ServiceDesc is the grpc.ServiceDesc for CoalitionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CoalitionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dcs.coalition.v0.CoalitionService",
	HandlerType: (*CoalitionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddGroup",
			Handler:    _CoalitionService_AddGroup_Handler,
		},
		{
			MethodName: "AddStaticObject",
			Handler:    _CoalitionService_AddStaticObject_Handler,
		},
		{
			MethodName: "AddLinkedStatic",
			Handler:    _CoalitionService_AddLinkedStatic_Handler,
		},
		{
			MethodName: "GetGroups",
			Handler:    _CoalitionService_GetGroups_Handler,
		},
		{
			MethodName: "GetBullseye",
			Handler:    _CoalitionService_GetBullseye_Handler,
		},
		{
			MethodName: "GetPlayerUnits",
			Handler:    _CoalitionService_GetPlayerUnits_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dcs/coalition/v0/coalition.proto",
}
