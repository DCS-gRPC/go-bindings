// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package custom

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

// CustomServiceClient is the client API for CustomService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomServiceClient interface {
	// DCT Function
	RequestMissionAssignment(ctx context.Context, in *RequestMissionAssignmentRequest, opts ...grpc.CallOption) (*RequestMissionAssignmentResponse, error)
	// DCT Function
	JoinMission(ctx context.Context, in *JoinMissionRequest, opts ...grpc.CallOption) (*JoinMissionResponse, error)
	// DCT Function
	AbortMission(ctx context.Context, in *AbortMissionRequest, opts ...grpc.CallOption) (*AbortMissionResponse, error)
	// DCT Function
	GetMissionStatus(ctx context.Context, in *GetMissionStatusRequest, opts ...grpc.CallOption) (*GetMissionStatusResponse, error)
	// Evaluate some Lua inside of the mission and return the result as a JSON
	// string. Disabled by default.
	Eval(ctx context.Context, in *EvalRequest, opts ...grpc.CallOption) (*EvalResponse, error)
	//*
	// Calculates the magnetic declination at the given position using the
	// International Geomagnetic Reference Field (IGRF) model. The result is not
	// always exactly the same as what DCS seem to use, but it is very close (DCS
	// doesn't expose its declination).
	GetMagneticDeclination(ctx context.Context, in *GetMagneticDeclinationRequest, opts ...grpc.CallOption) (*GetMagneticDeclinationResponse, error)
}

type customServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomServiceClient(cc grpc.ClientConnInterface) CustomServiceClient {
	return &customServiceClient{cc}
}

func (c *customServiceClient) RequestMissionAssignment(ctx context.Context, in *RequestMissionAssignmentRequest, opts ...grpc.CallOption) (*RequestMissionAssignmentResponse, error) {
	out := new(RequestMissionAssignmentResponse)
	err := c.cc.Invoke(ctx, "/dcs.custom.v0.CustomService/RequestMissionAssignment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customServiceClient) JoinMission(ctx context.Context, in *JoinMissionRequest, opts ...grpc.CallOption) (*JoinMissionResponse, error) {
	out := new(JoinMissionResponse)
	err := c.cc.Invoke(ctx, "/dcs.custom.v0.CustomService/JoinMission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customServiceClient) AbortMission(ctx context.Context, in *AbortMissionRequest, opts ...grpc.CallOption) (*AbortMissionResponse, error) {
	out := new(AbortMissionResponse)
	err := c.cc.Invoke(ctx, "/dcs.custom.v0.CustomService/AbortMission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customServiceClient) GetMissionStatus(ctx context.Context, in *GetMissionStatusRequest, opts ...grpc.CallOption) (*GetMissionStatusResponse, error) {
	out := new(GetMissionStatusResponse)
	err := c.cc.Invoke(ctx, "/dcs.custom.v0.CustomService/GetMissionStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customServiceClient) Eval(ctx context.Context, in *EvalRequest, opts ...grpc.CallOption) (*EvalResponse, error) {
	out := new(EvalResponse)
	err := c.cc.Invoke(ctx, "/dcs.custom.v0.CustomService/Eval", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customServiceClient) GetMagneticDeclination(ctx context.Context, in *GetMagneticDeclinationRequest, opts ...grpc.CallOption) (*GetMagneticDeclinationResponse, error) {
	out := new(GetMagneticDeclinationResponse)
	err := c.cc.Invoke(ctx, "/dcs.custom.v0.CustomService/GetMagneticDeclination", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomServiceServer is the server API for CustomService service.
// All implementations must embed UnimplementedCustomServiceServer
// for forward compatibility
type CustomServiceServer interface {
	// DCT Function
	RequestMissionAssignment(context.Context, *RequestMissionAssignmentRequest) (*RequestMissionAssignmentResponse, error)
	// DCT Function
	JoinMission(context.Context, *JoinMissionRequest) (*JoinMissionResponse, error)
	// DCT Function
	AbortMission(context.Context, *AbortMissionRequest) (*AbortMissionResponse, error)
	// DCT Function
	GetMissionStatus(context.Context, *GetMissionStatusRequest) (*GetMissionStatusResponse, error)
	// Evaluate some Lua inside of the mission and return the result as a JSON
	// string. Disabled by default.
	Eval(context.Context, *EvalRequest) (*EvalResponse, error)
	//*
	// Calculates the magnetic declination at the given position using the
	// International Geomagnetic Reference Field (IGRF) model. The result is not
	// always exactly the same as what DCS seem to use, but it is very close (DCS
	// doesn't expose its declination).
	GetMagneticDeclination(context.Context, *GetMagneticDeclinationRequest) (*GetMagneticDeclinationResponse, error)
	mustEmbedUnimplementedCustomServiceServer()
}

// UnimplementedCustomServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCustomServiceServer struct {
}

func (UnimplementedCustomServiceServer) RequestMissionAssignment(context.Context, *RequestMissionAssignmentRequest) (*RequestMissionAssignmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestMissionAssignment not implemented")
}
func (UnimplementedCustomServiceServer) JoinMission(context.Context, *JoinMissionRequest) (*JoinMissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinMission not implemented")
}
func (UnimplementedCustomServiceServer) AbortMission(context.Context, *AbortMissionRequest) (*AbortMissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AbortMission not implemented")
}
func (UnimplementedCustomServiceServer) GetMissionStatus(context.Context, *GetMissionStatusRequest) (*GetMissionStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMissionStatus not implemented")
}
func (UnimplementedCustomServiceServer) Eval(context.Context, *EvalRequest) (*EvalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Eval not implemented")
}
func (UnimplementedCustomServiceServer) GetMagneticDeclination(context.Context, *GetMagneticDeclinationRequest) (*GetMagneticDeclinationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMagneticDeclination not implemented")
}
func (UnimplementedCustomServiceServer) mustEmbedUnimplementedCustomServiceServer() {}

// UnsafeCustomServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomServiceServer will
// result in compilation errors.
type UnsafeCustomServiceServer interface {
	mustEmbedUnimplementedCustomServiceServer()
}

func RegisterCustomServiceServer(s grpc.ServiceRegistrar, srv CustomServiceServer) {
	s.RegisterService(&CustomService_ServiceDesc, srv)
}

func _CustomService_RequestMissionAssignment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestMissionAssignmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomServiceServer).RequestMissionAssignment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.custom.v0.CustomService/RequestMissionAssignment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomServiceServer).RequestMissionAssignment(ctx, req.(*RequestMissionAssignmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomService_JoinMission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinMissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomServiceServer).JoinMission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.custom.v0.CustomService/JoinMission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomServiceServer).JoinMission(ctx, req.(*JoinMissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomService_AbortMission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AbortMissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomServiceServer).AbortMission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.custom.v0.CustomService/AbortMission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomServiceServer).AbortMission(ctx, req.(*AbortMissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomService_GetMissionStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMissionStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomServiceServer).GetMissionStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.custom.v0.CustomService/GetMissionStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomServiceServer).GetMissionStatus(ctx, req.(*GetMissionStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomService_Eval_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EvalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomServiceServer).Eval(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.custom.v0.CustomService/Eval",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomServiceServer).Eval(ctx, req.(*EvalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomService_GetMagneticDeclination_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMagneticDeclinationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomServiceServer).GetMagneticDeclination(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.custom.v0.CustomService/GetMagneticDeclination",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomServiceServer).GetMagneticDeclination(ctx, req.(*GetMagneticDeclinationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomService_ServiceDesc is the grpc.ServiceDesc for CustomService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dcs.custom.v0.CustomService",
	HandlerType: (*CustomServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestMissionAssignment",
			Handler:    _CustomService_RequestMissionAssignment_Handler,
		},
		{
			MethodName: "JoinMission",
			Handler:    _CustomService_JoinMission_Handler,
		},
		{
			MethodName: "AbortMission",
			Handler:    _CustomService_AbortMission_Handler,
		},
		{
			MethodName: "GetMissionStatus",
			Handler:    _CustomService_GetMissionStatus_Handler,
		},
		{
			MethodName: "Eval",
			Handler:    _CustomService_Eval_Handler,
		},
		{
			MethodName: "GetMagneticDeclination",
			Handler:    _CustomService_GetMagneticDeclination_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dcs/custom/v0/custom.proto",
}
