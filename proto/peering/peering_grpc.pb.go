// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package peering

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

// PeeringClient is the client API for Peering service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PeeringClient interface {
	StartPeer(ctx context.Context, in *StartPeerRequest, opts ...grpc.CallOption) (*Empty, error)
	Route(ctx context.Context, in *RouteRequest, opts ...grpc.CallOption) (*Empty, error)
}

type peeringClient struct {
	cc grpc.ClientConnInterface
}

func NewPeeringClient(cc grpc.ClientConnInterface) PeeringClient {
	return &peeringClient{cc}
}

func (c *peeringClient) StartPeer(ctx context.Context, in *StartPeerRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/openfogstack.celestial.peering.Peering/StartPeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peeringClient) Route(ctx context.Context, in *RouteRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/openfogstack.celestial.peering.Peering/Route", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PeeringServer is the server API for Peering service.
// All implementations should embed UnimplementedPeeringServer
// for forward compatibility
type PeeringServer interface {
	StartPeer(context.Context, *StartPeerRequest) (*Empty, error)
	Route(context.Context, *RouteRequest) (*Empty, error)
}

// UnimplementedPeeringServer should be embedded to have forward compatible implementations.
type UnimplementedPeeringServer struct {
}

func (UnimplementedPeeringServer) StartPeer(context.Context, *StartPeerRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartPeer not implemented")
}
func (UnimplementedPeeringServer) Route(context.Context, *RouteRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Route not implemented")
}

// UnsafePeeringServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PeeringServer will
// result in compilation errors.
type UnsafePeeringServer interface {
	mustEmbedUnimplementedPeeringServer()
}

func RegisterPeeringServer(s grpc.ServiceRegistrar, srv PeeringServer) {
	s.RegisterService(&Peering_ServiceDesc, srv)
}

func _Peering_StartPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartPeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeeringServer).StartPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openfogstack.celestial.peering.Peering/StartPeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeeringServer).StartPeer(ctx, req.(*StartPeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Peering_Route_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeeringServer).Route(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openfogstack.celestial.peering.Peering/Route",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeeringServer).Route(ctx, req.(*RouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Peering_ServiceDesc is the grpc.ServiceDesc for Peering service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Peering_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "openfogstack.celestial.peering.Peering",
	HandlerType: (*PeeringServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartPeer",
			Handler:    _Peering_StartPeer_Handler,
		},
		{
			MethodName: "Route",
			Handler:    _Peering_Route_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "peering.proto",
}
