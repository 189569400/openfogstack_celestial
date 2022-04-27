// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: celestial.proto

package celestial

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

// CelestialClient is the client API for Celestial service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CelestialClient interface {
	GetHostInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HostInfo, error)
	HostReady(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ReadyInfo, error)
	Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*Empty, error)
	InitRemotes(ctx context.Context, in *InitRemotesRequest, opts ...grpc.CallOption) (*Empty, error)
	StartPeering(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	CreateMachine(ctx context.Context, in *CreateMachineRequest, opts ...grpc.CallOption) (*Empty, error)
	ModifyMachine(ctx context.Context, in *ModifyMachineRequest, opts ...grpc.CallOption) (*Empty, error)
	ModifyLinks(ctx context.Context, in *ModifyLinksRequest, opts ...grpc.CallOption) (*Empty, error)
}

type celestialClient struct {
	cc grpc.ClientConnInterface
}

func NewCelestialClient(cc grpc.ClientConnInterface) CelestialClient {
	return &celestialClient{cc}
}

func (c *celestialClient) GetHostInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HostInfo, error) {
	out := new(HostInfo)
	err := c.cc.Invoke(ctx, "/openfogstack.celestial.celestial.Celestial/GetHostInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *celestialClient) HostReady(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ReadyInfo, error) {
	out := new(ReadyInfo)
	err := c.cc.Invoke(ctx, "/openfogstack.celestial.celestial.Celestial/HostReady", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *celestialClient) Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/openfogstack.celestial.celestial.Celestial/Init", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *celestialClient) InitRemotes(ctx context.Context, in *InitRemotesRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/openfogstack.celestial.celestial.Celestial/InitRemotes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *celestialClient) StartPeering(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/openfogstack.celestial.celestial.Celestial/StartPeering", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *celestialClient) CreateMachine(ctx context.Context, in *CreateMachineRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/openfogstack.celestial.celestial.Celestial/CreateMachine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *celestialClient) ModifyMachine(ctx context.Context, in *ModifyMachineRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/openfogstack.celestial.celestial.Celestial/ModifyMachine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *celestialClient) ModifyLinks(ctx context.Context, in *ModifyLinksRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/openfogstack.celestial.celestial.Celestial/ModifyLinks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CelestialServer is the server API for Celestial service.
// All implementations should embed UnimplementedCelestialServer
// for forward compatibility
type CelestialServer interface {
	GetHostInfo(context.Context, *Empty) (*HostInfo, error)
	HostReady(context.Context, *Empty) (*ReadyInfo, error)
	Init(context.Context, *InitRequest) (*Empty, error)
	InitRemotes(context.Context, *InitRemotesRequest) (*Empty, error)
	StartPeering(context.Context, *Empty) (*Empty, error)
	CreateMachine(context.Context, *CreateMachineRequest) (*Empty, error)
	ModifyMachine(context.Context, *ModifyMachineRequest) (*Empty, error)
	ModifyLinks(context.Context, *ModifyLinksRequest) (*Empty, error)
}

// UnimplementedCelestialServer should be embedded to have forward compatible implementations.
type UnimplementedCelestialServer struct {
}

func (UnimplementedCelestialServer) GetHostInfo(context.Context, *Empty) (*HostInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHostInfo not implemented")
}
func (UnimplementedCelestialServer) HostReady(context.Context, *Empty) (*ReadyInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HostReady not implemented")
}
func (UnimplementedCelestialServer) Init(context.Context, *InitRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (UnimplementedCelestialServer) InitRemotes(context.Context, *InitRemotesRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitRemotes not implemented")
}
func (UnimplementedCelestialServer) StartPeering(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartPeering not implemented")
}
func (UnimplementedCelestialServer) CreateMachine(context.Context, *CreateMachineRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMachine not implemented")
}
func (UnimplementedCelestialServer) ModifyMachine(context.Context, *ModifyMachineRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyMachine not implemented")
}
func (UnimplementedCelestialServer) ModifyLinks(context.Context, *ModifyLinksRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyLinks not implemented")
}

// UnsafeCelestialServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CelestialServer will
// result in compilation errors.
type UnsafeCelestialServer interface {
	mustEmbedUnimplementedCelestialServer()
}

func RegisterCelestialServer(s grpc.ServiceRegistrar, srv CelestialServer) {
	s.RegisterService(&Celestial_ServiceDesc, srv)
}

func _Celestial_GetHostInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CelestialServer).GetHostInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openfogstack.celestial.celestial.Celestial/GetHostInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CelestialServer).GetHostInfo(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Celestial_HostReady_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CelestialServer).HostReady(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openfogstack.celestial.celestial.Celestial/HostReady",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CelestialServer).HostReady(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Celestial_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CelestialServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openfogstack.celestial.celestial.Celestial/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CelestialServer).Init(ctx, req.(*InitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Celestial_InitRemotes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitRemotesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CelestialServer).InitRemotes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openfogstack.celestial.celestial.Celestial/InitRemotes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CelestialServer).InitRemotes(ctx, req.(*InitRemotesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Celestial_StartPeering_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CelestialServer).StartPeering(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openfogstack.celestial.celestial.Celestial/StartPeering",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CelestialServer).StartPeering(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Celestial_CreateMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMachineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CelestialServer).CreateMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openfogstack.celestial.celestial.Celestial/CreateMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CelestialServer).CreateMachine(ctx, req.(*CreateMachineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Celestial_ModifyMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyMachineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CelestialServer).ModifyMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openfogstack.celestial.celestial.Celestial/ModifyMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CelestialServer).ModifyMachine(ctx, req.(*ModifyMachineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Celestial_ModifyLinks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyLinksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CelestialServer).ModifyLinks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openfogstack.celestial.celestial.Celestial/ModifyLinks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CelestialServer).ModifyLinks(ctx, req.(*ModifyLinksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Celestial_ServiceDesc is the grpc.ServiceDesc for Celestial service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Celestial_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "openfogstack.celestial.celestial.Celestial",
	HandlerType: (*CelestialServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHostInfo",
			Handler:    _Celestial_GetHostInfo_Handler,
		},
		{
			MethodName: "HostReady",
			Handler:    _Celestial_HostReady_Handler,
		},
		{
			MethodName: "Init",
			Handler:    _Celestial_Init_Handler,
		},
		{
			MethodName: "InitRemotes",
			Handler:    _Celestial_InitRemotes_Handler,
		},
		{
			MethodName: "StartPeering",
			Handler:    _Celestial_StartPeering_Handler,
		},
		{
			MethodName: "CreateMachine",
			Handler:    _Celestial_CreateMachine_Handler,
		},
		{
			MethodName: "ModifyMachine",
			Handler:    _Celestial_ModifyMachine_Handler,
		},
		{
			MethodName: "ModifyLinks",
			Handler:    _Celestial_ModifyLinks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "celestial.proto",
}
