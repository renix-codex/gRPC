// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: organisation.proto

package protobufs

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

const (
	OrgService_GetOrganisationByID_FullMethodName = "/organisation.OrgService/GetOrganisationByID"
)

// OrgServiceClient is the client API for OrgService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrgServiceClient interface {
	GetOrganisationByID(ctx context.Context, in *OrganisationIDRequest, opts ...grpc.CallOption) (*OrganisationResponse, error)
}

type orgServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrgServiceClient(cc grpc.ClientConnInterface) OrgServiceClient {
	return &orgServiceClient{cc}
}

func (c *orgServiceClient) GetOrganisationByID(ctx context.Context, in *OrganisationIDRequest, opts ...grpc.CallOption) (*OrganisationResponse, error) {
	out := new(OrganisationResponse)
	err := c.cc.Invoke(ctx, OrgService_GetOrganisationByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrgServiceServer is the server API for OrgService service.
// All implementations must embed UnimplementedOrgServiceServer
// for forward compatibility
type OrgServiceServer interface {
	GetOrganisationByID(context.Context, *OrganisationIDRequest) (*OrganisationResponse, error)
	mustEmbedUnimplementedOrgServiceServer()
}

// UnimplementedOrgServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrgServiceServer struct {
}

func (UnimplementedOrgServiceServer) GetOrganisationByID(context.Context, *OrganisationIDRequest) (*OrganisationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrganisationByID not implemented")
}
func (UnimplementedOrgServiceServer) mustEmbedUnimplementedOrgServiceServer() {}

// UnsafeOrgServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrgServiceServer will
// result in compilation errors.
type UnsafeOrgServiceServer interface {
	mustEmbedUnimplementedOrgServiceServer()
}

func RegisterOrgServiceServer(s grpc.ServiceRegistrar, srv OrgServiceServer) {
	s.RegisterService(&OrgService_ServiceDesc, srv)
}

func _OrgService_GetOrganisationByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrganisationIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgServiceServer).GetOrganisationByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrgService_GetOrganisationByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgServiceServer).GetOrganisationByID(ctx, req.(*OrganisationIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrgService_ServiceDesc is the grpc.ServiceDesc for OrgService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrgService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "organisation.OrgService",
	HandlerType: (*OrgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrganisationByID",
			Handler:    _OrgService_GetOrganisationByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "organisation.proto",
}
