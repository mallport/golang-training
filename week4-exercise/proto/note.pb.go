// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/note.proto

/*
Package note is a generated protocol buffer package.

It is generated from these files:
	proto/note.proto

It has these top-level messages:
	Note
	NoteReq
	NoteFindReq
	NoteDelReq
	NoteDelRes
*/
package note

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/gogo/protobuf/gogoproto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Note struct {
	Id        int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Title     string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Completed bool   `protobuf:"varint,3,opt,name=completed" json:"completed,omitempty"`
	// bool complete = 3; // Only change the name
	// int32 completed = 3; // Change data type
	CreatedAt   *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt   *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
	Description string                     `protobuf:"bytes,6,opt,name=description" json:"description,omitempty"`
}

func (m *Note) Reset()                    { *m = Note{} }
func (m *Note) String() string            { return proto.CompactTextString(m) }
func (*Note) ProtoMessage()               {}
func (*Note) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Note) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Note) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Note) GetCompleted() bool {
	if m != nil {
		return m.Completed
	}
	return false
}

func (m *Note) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Note) GetUpdatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Note) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type NoteReq struct {
	Title     string `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Completed bool   `protobuf:"varint,2,opt,name=completed" json:"completed,omitempty"`
}

func (m *NoteReq) Reset()                    { *m = NoteReq{} }
func (m *NoteReq) String() string            { return proto.CompactTextString(m) }
func (*NoteReq) ProtoMessage()               {}
func (*NoteReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *NoteReq) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *NoteReq) GetCompleted() bool {
	if m != nil {
		return m.Completed
	}
	return false
}

type NoteFindReq struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *NoteFindReq) Reset()                    { *m = NoteFindReq{} }
func (m *NoteFindReq) String() string            { return proto.CompactTextString(m) }
func (*NoteFindReq) ProtoMessage()               {}
func (*NoteFindReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *NoteFindReq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type NoteDelReq struct {
	Id int32 `protobuf:"varint,5,opt,name=id" json:"id,omitempty"`
}

func (m *NoteDelReq) Reset()                    { *m = NoteDelReq{} }
func (m *NoteDelReq) String() string            { return proto.CompactTextString(m) }
func (*NoteDelReq) ProtoMessage()               {}
func (*NoteDelReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *NoteDelReq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type NoteDelRes struct {
	Success      bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	ErrorMessage string `protobuf:"bytes,2,opt,name=error_message,json=errorMessage" json:"error_message,omitempty"`
}

func (m *NoteDelRes) Reset()                    { *m = NoteDelRes{} }
func (m *NoteDelRes) String() string            { return proto.CompactTextString(m) }
func (*NoteDelRes) ProtoMessage()               {}
func (*NoteDelRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *NoteDelRes) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *NoteDelRes) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*Note)(nil), "note.Note")
	proto.RegisterType((*NoteReq)(nil), "note.NoteReq")
	proto.RegisterType((*NoteFindReq)(nil), "note.NoteFindReq")
	proto.RegisterType((*NoteDelReq)(nil), "note.NoteDelReq")
	proto.RegisterType((*NoteDelRes)(nil), "note.NoteDelRes")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for NoteService service

type NoteServiceClient interface {
	Create(ctx context.Context, in *NoteReq, opts ...grpc.CallOption) (*Note, error)
	Find(ctx context.Context, in *NoteFindReq, opts ...grpc.CallOption) (*Note, error)
	// rpc Find(NoteFindReq) returns (Note) {
	//   option (google.api.http) = {
	//     get: "/v1/note/{id}"
	//   };
	// }
	Delete(ctx context.Context, in *NoteDelReq, opts ...grpc.CallOption) (*NoteDelRes, error)
}

type noteServiceClient struct {
	cc *grpc.ClientConn
}

func NewNoteServiceClient(cc *grpc.ClientConn) NoteServiceClient {
	return &noteServiceClient{cc}
}

func (c *noteServiceClient) Create(ctx context.Context, in *NoteReq, opts ...grpc.CallOption) (*Note, error) {
	out := new(Note)
	err := grpc.Invoke(ctx, "/note.NoteService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noteServiceClient) Find(ctx context.Context, in *NoteFindReq, opts ...grpc.CallOption) (*Note, error) {
	out := new(Note)
	err := grpc.Invoke(ctx, "/note.NoteService/Find", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noteServiceClient) Delete(ctx context.Context, in *NoteDelReq, opts ...grpc.CallOption) (*NoteDelRes, error) {
	out := new(NoteDelRes)
	err := grpc.Invoke(ctx, "/note.NoteService/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NoteService service

type NoteServiceServer interface {
	Create(context.Context, *NoteReq) (*Note, error)
	Find(context.Context, *NoteFindReq) (*Note, error)
	// rpc Find(NoteFindReq) returns (Note) {
	//   option (google.api.http) = {
	//     get: "/v1/note/{id}"
	//   };
	// }
	Delete(context.Context, *NoteDelReq) (*NoteDelRes, error)
}

func RegisterNoteServiceServer(s *grpc.Server, srv NoteServiceServer) {
	s.RegisterService(&_NoteService_serviceDesc, srv)
}

func _NoteService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/note.NoteService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteServiceServer).Create(ctx, req.(*NoteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NoteService_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoteFindReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteServiceServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/note.NoteService/Find",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteServiceServer).Find(ctx, req.(*NoteFindReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NoteService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoteDelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/note.NoteService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteServiceServer).Delete(ctx, req.(*NoteDelReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _NoteService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "note.NoteService",
	HandlerType: (*NoteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _NoteService_Create_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _NoteService_Find_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _NoteService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/note.proto",
}

func init() { proto.RegisterFile("proto/note.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xc1, 0x8a, 0xd4, 0x40,
	0x10, 0x35, 0x71, 0x66, 0x76, 0xa7, 0x66, 0x76, 0x19, 0x0b, 0x0f, 0x21, 0x8c, 0x4c, 0x68, 0x51,
	0x07, 0xc1, 0x04, 0xd6, 0x93, 0x82, 0xe0, 0xae, 0x8b, 0x17, 0xd1, 0x43, 0xf4, 0x3e, 0x64, 0xd2,
	0x65, 0x6c, 0x98, 0xa4, 0x63, 0xba, 0xe3, 0x45, 0xbc, 0x78, 0xf2, 0xee, 0x07, 0xf8, 0x51, 0xde,
	0xf7, 0xe4, 0x17, 0xec, 0x17, 0x48, 0x77, 0x27, 0x66, 0x18, 0x84, 0xbd, 0x55, 0xbd, 0xf7, 0xaa,
	0x52, 0xef, 0xa5, 0x61, 0x51, 0x37, 0x52, 0xcb, 0xa4, 0x92, 0x9a, 0x62, 0x5b, 0xe2, 0xc8, 0xd4,
	0xe1, 0xaa, 0x90, 0xb2, 0xd8, 0x51, 0x62, 0xb1, 0x6d, 0xfb, 0x31, 0xd1, 0xa2, 0x24, 0xa5, 0xb3,
	0xb2, 0x76, 0xb2, 0x70, 0xd9, 0x09, 0xb2, 0x5a, 0x24, 0x59, 0x55, 0x49, 0x9d, 0x69, 0x21, 0x2b,
	0xd5, 0xb1, 0x4f, 0x0a, 0xa1, 0x3f, 0xb5, 0xdb, 0x38, 0x97, 0x65, 0x52, 0xc8, 0x42, 0x0e, 0x7b,
	0x4c, 0xe7, 0xbe, 0x69, 0x2a, 0x27, 0x67, 0x3f, 0x7c, 0x18, 0xbd, 0x93, 0x9a, 0x90, 0x81, 0x2f,
	0x78, 0xe0, 0x45, 0xde, 0x7a, 0x7c, 0x81, 0xd7, 0x57, 0xab, 0xd3, 0xad, 0x92, 0xd5, 0x73, 0x66,
	0x4e, 0xda, 0x08, 0xce, 0x52, 0x5f, 0x70, 0x7c, 0x08, 0x63, 0x2d, 0xf4, 0x8e, 0x02, 0x3f, 0xf2,
	0xd6, 0xd3, 0x8b, 0xc5, 0xf5, 0xd5, 0x6a, 0xee, 0x64, 0x16, 0x66, 0xa9, 0xa3, 0x71, 0x09, 0xd3,
	0x5c, 0x96, 0xf5, 0x8e, 0x34, 0xf1, 0xe0, 0x76, 0xe4, 0xad, 0x8f, 0xd3, 0x01, 0xc0, 0x67, 0x00,
	0x79, 0x43, 0x99, 0x26, 0xbe, 0xc9, 0x74, 0x30, 0x8a, 0xbc, 0xf5, 0xec, 0x2c, 0x8c, 0x9d, 0xa9,
	0xb8, 0xbf, 0x36, 0xfe, 0xd0, 0xbb, 0x4e, 0xa7, 0x9d, 0xfa, 0x5c, 0x9b, 0xd1, 0xb6, 0xe6, 0xfd,
	0xe8, 0xf8, 0xe6, 0xd1, 0x4e, 0x7d, 0xae, 0x31, 0x82, 0x19, 0x27, 0x95, 0x37, 0xa2, 0x36, 0x69,
	0x05, 0x13, 0xe3, 0x20, 0xdd, 0x87, 0xd8, 0x0b, 0x38, 0x32, 0x49, 0xa4, 0xf4, 0x19, 0xef, 0xf6,
	0x46, 0x3d, 0x2b, 0xfb, 0x9f, 0x2d, 0xff, 0xc0, 0x16, 0xbb, 0x07, 0x33, 0x33, 0xfe, 0x5a, 0x54,
	0xdc, 0xac, 0x38, 0x1d, 0xf2, 0x34, 0xd9, 0xb1, 0x25, 0x80, 0xa1, 0x2f, 0x69, 0x37, 0xb0, 0xe3,
	0x7f, 0xec, 0x9b, 0x3d, 0x56, 0x61, 0x00, 0x47, 0xaa, 0xcd, 0x73, 0x52, 0xca, 0x2e, 0x38, 0x4e,
	0xfb, 0x16, 0xef, 0xc3, 0x09, 0x35, 0x8d, 0x6c, 0x36, 0x25, 0x29, 0x95, 0x15, 0xdd, 0x9f, 0x48,
	0xe7, 0x16, 0x7c, 0xeb, 0xb0, 0xb3, 0x5f, 0x9e, 0x3b, 0xe5, 0x3d, 0x35, 0x5f, 0x44, 0x4e, 0xf8,
	0x00, 0x26, 0xaf, 0x6c, 0x84, 0x78, 0x12, 0xdb, 0xe7, 0xd6, 0xd9, 0x0c, 0x61, 0x68, 0xd9, 0x2d,
	0x7c, 0x04, 0x23, 0x73, 0x3c, 0xde, 0x19, 0xd0, 0xce, 0xcc, 0x81, 0xf0, 0x25, 0x4c, 0x2e, 0xc9,
	0x98, 0xc6, 0xc5, 0x80, 0x3b, 0x63, 0xe1, 0x21, 0xa2, 0x18, 0x7e, 0xff, 0xfd, 0xe7, 0xa7, 0x3f,
	0x7f, 0x0c, 0xf6, 0xa9, 0x27, 0x5f, 0x05, 0xff, 0xb6, 0x9d, 0xd8, 0x7f, 0xf5, 0xf4, 0x6f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x94, 0xb0, 0x51, 0xfc, 0x04, 0x03, 0x00, 0x00,
}