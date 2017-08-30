// Code generated by protoc-gen-go.
// source: provider.proto
// DO NOT EDIT!

package lumirpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"
import google_protobuf "github.com/golang/protobuf/ptypes/struct"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CheckRequest struct {
	Type       string                  `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Properties *google_protobuf.Struct `protobuf:"bytes,2,opt,name=properties" json:"properties,omitempty"`
}

func (m *CheckRequest) Reset()                    { *m = CheckRequest{} }
func (m *CheckRequest) String() string            { return proto.CompactTextString(m) }
func (*CheckRequest) ProtoMessage()               {}
func (*CheckRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *CheckRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *CheckRequest) GetProperties() *google_protobuf.Struct {
	if m != nil {
		return m.Properties
	}
	return nil
}

type CheckResponse struct {
	Defaults *google_protobuf.Struct `protobuf:"bytes,1,opt,name=defaults" json:"defaults,omitempty"`
	Failures []*CheckFailure         `protobuf:"bytes,2,rep,name=failures" json:"failures,omitempty"`
}

func (m *CheckResponse) Reset()                    { *m = CheckResponse{} }
func (m *CheckResponse) String() string            { return proto.CompactTextString(m) }
func (*CheckResponse) ProtoMessage()               {}
func (*CheckResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *CheckResponse) GetDefaults() *google_protobuf.Struct {
	if m != nil {
		return m.Defaults
	}
	return nil
}

func (m *CheckResponse) GetFailures() []*CheckFailure {
	if m != nil {
		return m.Failures
	}
	return nil
}

type CheckFailure struct {
	Property string `protobuf:"bytes,1,opt,name=property" json:"property,omitempty"`
	Reason   string `protobuf:"bytes,2,opt,name=reason" json:"reason,omitempty"`
}

func (m *CheckFailure) Reset()                    { *m = CheckFailure{} }
func (m *CheckFailure) String() string            { return proto.CompactTextString(m) }
func (*CheckFailure) ProtoMessage()               {}
func (*CheckFailure) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *CheckFailure) GetProperty() string {
	if m != nil {
		return m.Property
	}
	return ""
}

func (m *CheckFailure) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

type DiffRequest struct {
	Id   string                  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Type string                  `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Olds *google_protobuf.Struct `protobuf:"bytes,3,opt,name=olds" json:"olds,omitempty"`
	News *google_protobuf.Struct `protobuf:"bytes,4,opt,name=news" json:"news,omitempty"`
}

func (m *DiffRequest) Reset()                    { *m = DiffRequest{} }
func (m *DiffRequest) String() string            { return proto.CompactTextString(m) }
func (*DiffRequest) ProtoMessage()               {}
func (*DiffRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *DiffRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DiffRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *DiffRequest) GetOlds() *google_protobuf.Struct {
	if m != nil {
		return m.Olds
	}
	return nil
}

func (m *DiffRequest) GetNews() *google_protobuf.Struct {
	if m != nil {
		return m.News
	}
	return nil
}

type DiffResponse struct {
	Replaces []string `protobuf:"bytes,1,rep,name=replaces" json:"replaces,omitempty"`
}

func (m *DiffResponse) Reset()                    { *m = DiffResponse{} }
func (m *DiffResponse) String() string            { return proto.CompactTextString(m) }
func (*DiffResponse) ProtoMessage()               {}
func (*DiffResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *DiffResponse) GetReplaces() []string {
	if m != nil {
		return m.Replaces
	}
	return nil
}

type CreateRequest struct {
	Type       string                  `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Properties *google_protobuf.Struct `protobuf:"bytes,2,opt,name=properties" json:"properties,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

func (m *CreateRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *CreateRequest) GetProperties() *google_protobuf.Struct {
	if m != nil {
		return m.Properties
	}
	return nil
}

type CreateResponse struct {
	Id         string                  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Properties *google_protobuf.Struct `protobuf:"bytes,2,opt,name=properties" json:"properties,omitempty"`
}

func (m *CreateResponse) Reset()                    { *m = CreateResponse{} }
func (m *CreateResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()               {}
func (*CreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6} }

func (m *CreateResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateResponse) GetProperties() *google_protobuf.Struct {
	if m != nil {
		return m.Properties
	}
	return nil
}

type GetRequest struct {
	Id   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{7} }

func (m *GetRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type GetResponse struct {
	Properties *google_protobuf.Struct `protobuf:"bytes,1,opt,name=properties" json:"properties,omitempty"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{8} }

func (m *GetResponse) GetProperties() *google_protobuf.Struct {
	if m != nil {
		return m.Properties
	}
	return nil
}

type UpdateRequest struct {
	Id   string                  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Type string                  `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Olds *google_protobuf.Struct `protobuf:"bytes,3,opt,name=olds" json:"olds,omitempty"`
	News *google_protobuf.Struct `protobuf:"bytes,4,opt,name=news" json:"news,omitempty"`
}

func (m *UpdateRequest) Reset()                    { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()               {}
func (*UpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{9} }

func (m *UpdateRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *UpdateRequest) GetOlds() *google_protobuf.Struct {
	if m != nil {
		return m.Olds
	}
	return nil
}

func (m *UpdateRequest) GetNews() *google_protobuf.Struct {
	if m != nil {
		return m.News
	}
	return nil
}

type UpdateResponse struct {
	Properties *google_protobuf.Struct `protobuf:"bytes,1,opt,name=properties" json:"properties,omitempty"`
}

func (m *UpdateResponse) Reset()                    { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()               {}
func (*UpdateResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{10} }

func (m *UpdateResponse) GetProperties() *google_protobuf.Struct {
	if m != nil {
		return m.Properties
	}
	return nil
}

type DeleteRequest struct {
	Id         string                  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Type       string                  `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Properties *google_protobuf.Struct `protobuf:"bytes,3,opt,name=properties" json:"properties,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{11} }

func (m *DeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DeleteRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *DeleteRequest) GetProperties() *google_protobuf.Struct {
	if m != nil {
		return m.Properties
	}
	return nil
}

func init() {
	proto.RegisterType((*CheckRequest)(nil), "lumirpc.CheckRequest")
	proto.RegisterType((*CheckResponse)(nil), "lumirpc.CheckResponse")
	proto.RegisterType((*CheckFailure)(nil), "lumirpc.CheckFailure")
	proto.RegisterType((*DiffRequest)(nil), "lumirpc.DiffRequest")
	proto.RegisterType((*DiffResponse)(nil), "lumirpc.DiffResponse")
	proto.RegisterType((*CreateRequest)(nil), "lumirpc.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "lumirpc.CreateResponse")
	proto.RegisterType((*GetRequest)(nil), "lumirpc.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "lumirpc.GetResponse")
	proto.RegisterType((*UpdateRequest)(nil), "lumirpc.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "lumirpc.UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "lumirpc.DeleteRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ResourceProvider service

type ResourceProviderClient interface {
	// Check validates that the given property bag is valid for a resource of the given type.
	Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error)
	// Diff checks what impacts a hypothetical update will have on the resource's properties.
	Diff(ctx context.Context, in *DiffRequest, opts ...grpc.CallOption) (*DiffResponse, error)
	// Create allocates a new instance of the provided resource and returns its unique ID afterwards.  (The input ID
	// must be blank.)  If this call fails, the resource must not have been created (i.e., it is "transacational").
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// Get reads the instance state identified by ID, returning a populated resource object, or nil if not found.
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	// Update updates an existing resource with new values.
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	// Delete tears down an existing resource with the given ID.  If it fails, the resource is assumed to still exist.
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
}

type resourceProviderClient struct {
	cc *grpc.ClientConn
}

func NewResourceProviderClient(cc *grpc.ClientConn) ResourceProviderClient {
	return &resourceProviderClient{cc}
}

func (c *resourceProviderClient) Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := grpc.Invoke(ctx, "/lumirpc.ResourceProvider/Check", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceProviderClient) Diff(ctx context.Context, in *DiffRequest, opts ...grpc.CallOption) (*DiffResponse, error) {
	out := new(DiffResponse)
	err := grpc.Invoke(ctx, "/lumirpc.ResourceProvider/Diff", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceProviderClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := grpc.Invoke(ctx, "/lumirpc.ResourceProvider/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceProviderClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc.Invoke(ctx, "/lumirpc.ResourceProvider/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceProviderClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := grpc.Invoke(ctx, "/lumirpc.ResourceProvider/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceProviderClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/lumirpc.ResourceProvider/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ResourceProvider service

type ResourceProviderServer interface {
	// Check validates that the given property bag is valid for a resource of the given type.
	Check(context.Context, *CheckRequest) (*CheckResponse, error)
	// Diff checks what impacts a hypothetical update will have on the resource's properties.
	Diff(context.Context, *DiffRequest) (*DiffResponse, error)
	// Create allocates a new instance of the provided resource and returns its unique ID afterwards.  (The input ID
	// must be blank.)  If this call fails, the resource must not have been created (i.e., it is "transacational").
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	// Get reads the instance state identified by ID, returning a populated resource object, or nil if not found.
	Get(context.Context, *GetRequest) (*GetResponse, error)
	// Update updates an existing resource with new values.
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	// Delete tears down an existing resource with the given ID.  If it fails, the resource is assumed to still exist.
	Delete(context.Context, *DeleteRequest) (*google_protobuf1.Empty, error)
}

func RegisterResourceProviderServer(s *grpc.Server, srv ResourceProviderServer) {
	s.RegisterService(&_ResourceProvider_serviceDesc, srv)
}

func _ResourceProvider_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceProviderServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lumirpc.ResourceProvider/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceProviderServer).Check(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceProvider_Diff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceProviderServer).Diff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lumirpc.ResourceProvider/Diff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceProviderServer).Diff(ctx, req.(*DiffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceProvider_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceProviderServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lumirpc.ResourceProvider/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceProviderServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceProvider_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceProviderServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lumirpc.ResourceProvider/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceProviderServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceProvider_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceProviderServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lumirpc.ResourceProvider/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceProviderServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceProvider_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceProviderServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lumirpc.ResourceProvider/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceProviderServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResourceProvider_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lumirpc.ResourceProvider",
	HandlerType: (*ResourceProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _ResourceProvider_Check_Handler,
		},
		{
			MethodName: "Diff",
			Handler:    _ResourceProvider_Diff_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ResourceProvider_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ResourceProvider_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ResourceProvider_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ResourceProvider_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "provider.proto",
}

func init() { proto.RegisterFile("provider.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 506 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xc4, 0x54, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x5d, 0x92, 0x52, 0xba, 0xdb, 0xb5, 0x42, 0x66, 0xeb, 0xaa, 0xc0, 0xc3, 0xe4, 0xa7, 0x0a,
	0xa4, 0x0c, 0x3a, 0x21, 0x10, 0xbc, 0xb1, 0xb1, 0x89, 0x37, 0x14, 0xc4, 0x03, 0x82, 0x97, 0x2c,
	0xb9, 0x19, 0x11, 0x59, 0x6d, 0x6c, 0x87, 0xa9, 0xdf, 0x80, 0xf8, 0x54, 0xfe, 0x01, 0xc5, 0x76,
	0x82, 0x93, 0x49, 0x63, 0x03, 0x21, 0xde, 0xe2, 0xeb, 0xe3, 0x73, 0xee, 0x39, 0xd7, 0x0e, 0x4c,
	0xb9, 0x60, 0x5f, 0x8b, 0x0c, 0x45, 0xc4, 0x05, 0x53, 0x8c, 0xdc, 0x2e, 0xab, 0xf3, 0x42, 0xf0,
	0x34, 0xbc, 0x77, 0xc6, 0xd8, 0x59, 0x89, 0xfb, 0xba, 0x7c, 0x5a, 0xe5, 0xfb, 0x78, 0xce, 0xd5,
	0xda, 0xa0, 0xc2, 0xfb, 0xfd, 0x4d, 0xa9, 0x44, 0x95, 0x2a, 0xb3, 0x4b, 0x3f, 0xc0, 0xd6, 0xe1,
	0x27, 0x4c, 0x3f, 0xc7, 0xf8, 0xa5, 0x42, 0xa9, 0x08, 0x81, 0x81, 0x5a, 0x73, 0x9c, 0x7b, 0x7b,
	0xde, 0x62, 0x33, 0xd6, 0xdf, 0xe4, 0x29, 0x00, 0x17, 0x8c, 0xa3, 0x50, 0x05, 0xca, 0xb9, 0xbf,
	0xe7, 0x2d, 0xc6, 0xcb, 0xdd, 0xc8, 0xd0, 0x46, 0x0d, 0x6d, 0xf4, 0x56, 0xd3, 0xc6, 0x0e, 0x94,
	0x5e, 0xc0, 0xc4, 0x92, 0x4b, 0xce, 0x56, 0x12, 0xc9, 0x01, 0x8c, 0x32, 0xcc, 0x93, 0xaa, 0x54,
	0x52, 0x2b, 0x5c, 0xc1, 0xd3, 0x02, 0xc9, 0x63, 0x18, 0xe5, 0x49, 0x51, 0x56, 0x42, 0x8b, 0x07,
	0x8b, 0xf1, 0x72, 0x27, 0xb2, 0xce, 0x23, 0x4d, 0x7f, 0x6c, 0x76, 0xe3, 0x16, 0x46, 0x5f, 0x5a,
	0x57, 0x76, 0x87, 0x84, 0x30, 0xb2, 0x6d, 0xad, 0xad, 0xb3, 0x76, 0x4d, 0x66, 0x30, 0x14, 0x98,
	0x48, 0xb6, 0xd2, 0xce, 0x36, 0x63, 0xbb, 0xa2, 0xdf, 0x3c, 0x18, 0x1f, 0x15, 0x79, 0xde, 0x24,
	0x33, 0x05, 0xbf, 0xc8, 0xec, 0x69, 0xbf, 0xc8, 0xda, 0xa4, 0x7c, 0x27, 0xa9, 0x87, 0x30, 0x60,
	0x65, 0x26, 0xe7, 0xc1, 0xd5, 0xde, 0x34, 0xa8, 0x06, 0xaf, 0xf0, 0x42, 0xce, 0x07, 0xbf, 0x01,
	0xd7, 0x20, 0xfa, 0x00, 0xb6, 0x4c, 0x33, 0x36, 0xc9, 0x10, 0x46, 0x02, 0x79, 0x99, 0xa4, 0x58,
	0x27, 0x19, 0xd4, 0x8e, 0x9a, 0x35, 0xfd, 0x08, 0x93, 0x43, 0x81, 0x89, 0xc2, 0x7f, 0x32, 0xd4,
	0xf7, 0x30, 0x6d, 0xd8, 0x6d, 0x2f, 0xfd, 0x64, 0xfe, 0x98, 0xfa, 0x11, 0xc0, 0x09, 0xaa, 0x1b,
	0x04, 0x4e, 0x8f, 0x61, 0xac, 0x4f, 0xd8, 0x4e, 0xba, 0xca, 0xde, 0xf5, 0x95, 0xbf, 0x7b, 0x30,
	0x79, 0xc7, 0x33, 0x27, 0xb3, 0xff, 0x3b, 0xee, 0xd7, 0x30, 0x6d, 0xda, 0xf9, 0x5b, 0x6b, 0x25,
	0x4c, 0x8e, 0xb0, 0xc4, 0x9b, 0x39, 0xeb, 0xaa, 0x05, 0xd7, 0x56, 0x5b, 0xfe, 0xf0, 0xe1, 0x4e,
	0x8c, 0x92, 0x55, 0x22, 0xc5, 0x37, 0xf6, 0x77, 0x45, 0x9e, 0xc1, 0x2d, 0xfd, 0x1c, 0x49, 0xef,
	0xe1, 0xda, 0x8e, 0xc2, 0x59, 0xbf, 0x6c, 0x3c, 0xd3, 0x0d, 0xf2, 0x04, 0x06, 0xf5, 0xb5, 0x27,
	0xdb, 0x2d, 0xc2, 0x79, 0x92, 0xe1, 0x4e, 0xaf, 0xda, 0x1e, 0x7b, 0x01, 0x43, 0x73, 0x47, 0x89,
	0x43, 0xed, 0x3e, 0x89, 0x70, 0xf7, 0x52, 0xbd, 0x3d, 0xbc, 0x84, 0xe0, 0x04, 0x15, 0xb9, 0xdb,
	0x22, 0x7e, 0xdd, 0xc9, 0x70, 0xbb, 0x5b, 0x74, 0x05, 0xcd, 0xbc, 0x1c, 0xc1, 0xce, 0x7d, 0x72,
	0x04, 0xbb, 0x83, 0xa5, 0x1b, 0xe4, 0x39, 0x0c, 0xcd, 0x84, 0x9c, 0xc3, 0x9d, 0x91, 0x85, 0xb3,
	0x4b, 0xd1, 0xbf, 0xaa, 0xff, 0xf0, 0x74, 0xe3, 0x74, 0xa8, 0x2b, 0x07, 0x3f, 0x03, 0x00, 0x00,
	0xff, 0xff, 0xe2, 0x60, 0x2e, 0x37, 0x1c, 0x06, 0x00, 0x00,
}
