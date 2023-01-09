// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// VTSClient is the client API for VTS service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VTSClient interface {
	// Return the summary state of the service.
	GetServiceState(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ServiceState, error)
	// Returns attestation information -- evidences, endorsed claims, trust
	// vector, etc -- for the provided attestation token data.
	GetAttestation(ctx context.Context, in *AttestationToken, opts ...grpc.CallOption) (*AppraisalContext, error)
	GetSupportedVerificationMediaTypes(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*MediaTypeList, error)
	// Service endpoints that are used to store Reference Values
	// and Trust Anchors to the endorsement store
	AddRefValues(ctx context.Context, in *AddRefValuesRequest, opts ...grpc.CallOption) (*AddRefValuesResponse, error)
	AddTrustAnchor(ctx context.Context, in *AddTrustAnchorRequest, opts ...grpc.CallOption) (*AddTrustAnchorResponse, error)
}

type vTSClient struct {
	cc grpc.ClientConnInterface
}

func NewVTSClient(cc grpc.ClientConnInterface) VTSClient {
	return &vTSClient{cc}
}

func (c *vTSClient) GetServiceState(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ServiceState, error) {
	out := new(ServiceState)
	err := c.cc.Invoke(ctx, "/proto.VTS/GetServiceState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vTSClient) GetAttestation(ctx context.Context, in *AttestationToken, opts ...grpc.CallOption) (*AppraisalContext, error) {
	out := new(AppraisalContext)
	err := c.cc.Invoke(ctx, "/proto.VTS/GetAttestation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vTSClient) GetSupportedVerificationMediaTypes(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*MediaTypeList, error) {
	out := new(MediaTypeList)
	err := c.cc.Invoke(ctx, "/proto.VTS/GetSupportedVerificationMediaTypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vTSClient) AddRefValues(ctx context.Context, in *AddRefValuesRequest, opts ...grpc.CallOption) (*AddRefValuesResponse, error) {
	out := new(AddRefValuesResponse)
	err := c.cc.Invoke(ctx, "/proto.VTS/AddRefValues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vTSClient) AddTrustAnchor(ctx context.Context, in *AddTrustAnchorRequest, opts ...grpc.CallOption) (*AddTrustAnchorResponse, error) {
	out := new(AddTrustAnchorResponse)
	err := c.cc.Invoke(ctx, "/proto.VTS/AddTrustAnchor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VTSServer is the server API for VTS service.
// All implementations must embed UnimplementedVTSServer
// for forward compatibility
type VTSServer interface {
	// Return the summary state of the service.
	GetServiceState(context.Context, *empty.Empty) (*ServiceState, error)
	// Returns attestation information -- evidences, endorsed claims, trust
	// vector, etc -- for the provided attestation token data.
	GetAttestation(context.Context, *AttestationToken) (*AppraisalContext, error)
	GetSupportedVerificationMediaTypes(context.Context, *empty.Empty) (*MediaTypeList, error)
	// Service endpoints that are used to store Reference Values
	// and Trust Anchors to the endorsement store
	AddRefValues(context.Context, *AddRefValuesRequest) (*AddRefValuesResponse, error)
	AddTrustAnchor(context.Context, *AddTrustAnchorRequest) (*AddTrustAnchorResponse, error)
	mustEmbedUnimplementedVTSServer()
}

// UnimplementedVTSServer must be embedded to have forward compatible implementations.
type UnimplementedVTSServer struct {
}

func (UnimplementedVTSServer) GetServiceState(context.Context, *empty.Empty) (*ServiceState, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServiceState not implemented")
}
func (UnimplementedVTSServer) GetAttestation(context.Context, *AttestationToken) (*AppraisalContext, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAttestation not implemented")
}
func (UnimplementedVTSServer) GetSupportedVerificationMediaTypes(context.Context, *empty.Empty) (*MediaTypeList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSupportedVerificationMediaTypes not implemented")
}
func (UnimplementedVTSServer) AddRefValues(context.Context, *AddRefValuesRequest) (*AddRefValuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRefValues not implemented")
}
func (UnimplementedVTSServer) AddTrustAnchor(context.Context, *AddTrustAnchorRequest) (*AddTrustAnchorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTrustAnchor not implemented")
}
func (UnimplementedVTSServer) mustEmbedUnimplementedVTSServer() {}

// UnsafeVTSServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VTSServer will
// result in compilation errors.
type UnsafeVTSServer interface {
	mustEmbedUnimplementedVTSServer()
}

func RegisterVTSServer(s grpc.ServiceRegistrar, srv VTSServer) {
	s.RegisterService(&VTS_ServiceDesc, srv)
}

func _VTS_GetServiceState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VTSServer).GetServiceState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.VTS/GetServiceState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VTSServer).GetServiceState(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _VTS_GetAttestation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttestationToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VTSServer).GetAttestation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.VTS/GetAttestation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VTSServer).GetAttestation(ctx, req.(*AttestationToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _VTS_GetSupportedVerificationMediaTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VTSServer).GetSupportedVerificationMediaTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.VTS/GetSupportedVerificationMediaTypes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VTSServer).GetSupportedVerificationMediaTypes(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _VTS_AddRefValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRefValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VTSServer).AddRefValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.VTS/AddRefValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VTSServer).AddRefValues(ctx, req.(*AddRefValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VTS_AddTrustAnchor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTrustAnchorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VTSServer).AddTrustAnchor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.VTS/AddTrustAnchor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VTSServer).AddTrustAnchor(ctx, req.(*AddTrustAnchorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VTS_ServiceDesc is the grpc.ServiceDesc for VTS service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VTS_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.VTS",
	HandlerType: (*VTSServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetServiceState",
			Handler:    _VTS_GetServiceState_Handler,
		},
		{
			MethodName: "GetAttestation",
			Handler:    _VTS_GetAttestation_Handler,
		},
		{
			MethodName: "GetSupportedVerificationMediaTypes",
			Handler:    _VTS_GetSupportedVerificationMediaTypes_Handler,
		},
		{
			MethodName: "AddRefValues",
			Handler:    _VTS_AddRefValues_Handler,
		},
		{
			MethodName: "AddTrustAnchor",
			Handler:    _VTS_AddTrustAnchor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vts.proto",
}
