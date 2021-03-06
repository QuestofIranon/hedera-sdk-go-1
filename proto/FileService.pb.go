// Code generated by protoc-gen-go. DO NOT EDIT.
// source: FileService.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("FileService.proto", fileDescriptor_5003a56c7e235889) }

var fileDescriptor_5003a56c7e235889 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x90, 0xcd, 0x4a, 0xc4, 0x30,
	0x14, 0x85, 0x17, 0xa2, 0x8b, 0xdb, 0xb1, 0x32, 0xd9, 0x19, 0x7c, 0x04, 0x89, 0xa0, 0x5b, 0x51,
	0x51, 0x11, 0x5c, 0xfa, 0xf7, 0x00, 0x31, 0xb9, 0x4e, 0x0a, 0xd3, 0x9b, 0x90, 0xa4, 0x42, 0x1f,
	0xd9, 0xb7, 0x90, 0x26, 0xa9, 0x14, 0x5d, 0x48, 0x3b, 0xab, 0x4b, 0xce, 0x3d, 0xdf, 0x39, 0x49,
	0x60, 0xfd, 0xd0, 0x6c, 0xf1, 0x05, 0xfd, 0x67, 0xa3, 0x50, 0x38, 0x6f, 0xa3, 0x65, 0xfb, 0x69,
	0xf0, 0xea, 0xa9, 0x43, 0xdf, 0x67, 0x8d, 0xd7, 0xcf, 0x18, 0x9c, 0xa5, 0x50, 0x3c, 0xfc, 0xf8,
	0xd5, 0x4b, 0x0a, 0x52, 0xc5, 0xc6, 0xd2, 0xaf, 0xd5, 0x7a, 0xb2, 0xca, 0xd2, 0xf9, 0xd7, 0x1e,
	0x54, 0x93, 0x1e, 0x76, 0x09, 0xa0, 0x3c, 0xca, 0x88, 0x83, 0xc8, 0x58, 0x76, 0x89, 0x09, 0xc7,
	0xf9, 0x5f, 0x6d, 0xac, 0x19, 0xe8, 0xce, 0xe9, 0x1d, 0x68, 0x8d, 0x5b, 0x5c, 0x48, 0x5f, 0xc3,
	0xa1, 0x74, 0x0e, 0x49, 0xdf, 0x59, 0x8a, 0x48, 0x71, 0x76, 0xc0, 0x19, 0xd4, 0x1b, 0x8c, 0x43,
	0xf7, 0x98, 0xb0, 0x2a, 0xee, 0xf4, 0xdd, 0xfc, 0xa8, 0x9c, 0x7e, 0x80, 0x53, 0xa8, 0x0a, 0xf0,
	0x48, 0x1f, 0xf6, 0x3f, 0xf7, 0x15, 0xac, 0x42, 0x1f, 0x22, 0xb6, 0xf7, 0xe9, 0x8d, 0xb3, 0xaf,
	0x77, 0x03, 0x75, 0xe6, 0xdf, 0x48, 0x2f, 0x4a, 0xb8, 0x3d, 0x01, 0xae, 0x6c, 0x2b, 0x0c, 0x6a,
	0xf4, 0x52, 0x18, 0x19, 0xcc, 0xc6, 0x4b, 0x67, 0x32, 0xf1, 0x7e, 0x90, 0xc6, 0xc5, 0x77, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x84, 0x02, 0xc2, 0x4d, 0x77, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FileServiceClient is the client API for FileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FileServiceClient interface {
	CreateFile(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	UpdateFile(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	DeleteFile(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	AppendContent(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	GetFileContent(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
	GetFileInfo(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
	SystemDelete(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	SystemUndelete(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
}

type fileServiceClient struct {
	cc *grpc.ClientConn
}

func NewFileServiceClient(cc *grpc.ClientConn) FileServiceClient {
	return &fileServiceClient{cc}
}

func (c *fileServiceClient) CreateFile(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.FileService/createFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) UpdateFile(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.FileService/updateFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) DeleteFile(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.FileService/deleteFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) AppendContent(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.FileService/appendContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) GetFileContent(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.FileService/getFileContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) GetFileInfo(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.FileService/getFileInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) SystemDelete(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.FileService/systemDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) SystemUndelete(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.FileService/systemUndelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileServiceServer is the server API for FileService service.
type FileServiceServer interface {
	CreateFile(context.Context, *Transaction) (*TransactionResponse, error)
	UpdateFile(context.Context, *Transaction) (*TransactionResponse, error)
	DeleteFile(context.Context, *Transaction) (*TransactionResponse, error)
	AppendContent(context.Context, *Transaction) (*TransactionResponse, error)
	GetFileContent(context.Context, *Query) (*Response, error)
	GetFileInfo(context.Context, *Query) (*Response, error)
	SystemDelete(context.Context, *Transaction) (*TransactionResponse, error)
	SystemUndelete(context.Context, *Transaction) (*TransactionResponse, error)
}

// UnimplementedFileServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFileServiceServer struct {
}

func (*UnimplementedFileServiceServer) CreateFile(ctx context.Context, req *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFile not implemented")
}
func (*UnimplementedFileServiceServer) UpdateFile(ctx context.Context, req *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFile not implemented")
}
func (*UnimplementedFileServiceServer) DeleteFile(ctx context.Context, req *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFile not implemented")
}
func (*UnimplementedFileServiceServer) AppendContent(ctx context.Context, req *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppendContent not implemented")
}
func (*UnimplementedFileServiceServer) GetFileContent(ctx context.Context, req *Query) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFileContent not implemented")
}
func (*UnimplementedFileServiceServer) GetFileInfo(ctx context.Context, req *Query) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFileInfo not implemented")
}
func (*UnimplementedFileServiceServer) SystemDelete(ctx context.Context, req *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemDelete not implemented")
}
func (*UnimplementedFileServiceServer) SystemUndelete(ctx context.Context, req *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemUndelete not implemented")
}

func RegisterFileServiceServer(s *grpc.Server, srv FileServiceServer) {
	s.RegisterService(&_FileService_serviceDesc, srv)
}

func _FileService_CreateFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).CreateFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FileService/CreateFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).CreateFile(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_UpdateFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).UpdateFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FileService/UpdateFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).UpdateFile(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_DeleteFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).DeleteFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FileService/DeleteFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).DeleteFile(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_AppendContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).AppendContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FileService/AppendContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).AppendContent(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_GetFileContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).GetFileContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FileService/GetFileContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).GetFileContent(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_GetFileInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).GetFileInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FileService/GetFileInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).GetFileInfo(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_SystemDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).SystemDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FileService/SystemDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).SystemDelete(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_SystemUndelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).SystemUndelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FileService/SystemUndelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).SystemUndelete(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

var _FileService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.FileService",
	HandlerType: (*FileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createFile",
			Handler:    _FileService_CreateFile_Handler,
		},
		{
			MethodName: "updateFile",
			Handler:    _FileService_UpdateFile_Handler,
		},
		{
			MethodName: "deleteFile",
			Handler:    _FileService_DeleteFile_Handler,
		},
		{
			MethodName: "appendContent",
			Handler:    _FileService_AppendContent_Handler,
		},
		{
			MethodName: "getFileContent",
			Handler:    _FileService_GetFileContent_Handler,
		},
		{
			MethodName: "getFileInfo",
			Handler:    _FileService_GetFileInfo_Handler,
		},
		{
			MethodName: "systemDelete",
			Handler:    _FileService_SystemDelete_Handler,
		},
		{
			MethodName: "systemUndelete",
			Handler:    _FileService_SystemUndelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "FileService.proto",
}
