// Code generated by protoc-gen-go. DO NOT EDIT.
// source: interaction.proto

package interaction

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

type InteractionStatusCode int32

const (
	InteractionStatusCode_Unknown InteractionStatusCode = 0
	InteractionStatusCode_Ok      InteractionStatusCode = 1
	InteractionStatusCode_Failed  InteractionStatusCode = 2
)

var InteractionStatusCode_name = map[int32]string{
	0: "Unknown",
	1: "Ok",
	2: "Failed",
}

var InteractionStatusCode_value = map[string]int32{
	"Unknown": 0,
	"Ok":      1,
	"Failed":  2,
}

func (x InteractionStatusCode) String() string {
	return proto.EnumName(InteractionStatusCode_name, int32(x))
}

func (InteractionStatusCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_43fb9bc5eb13a2f0, []int{0}
}

type InteractionStatus struct {
	Code                 InteractionStatusCode `protobuf:"varint,1,opt,name=Code,proto3,enum=interaction.InteractionStatusCode" json:"Code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *InteractionStatus) Reset()         { *m = InteractionStatus{} }
func (m *InteractionStatus) String() string { return proto.CompactTextString(m) }
func (*InteractionStatus) ProtoMessage()    {}
func (*InteractionStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_43fb9bc5eb13a2f0, []int{0}
}

func (m *InteractionStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InteractionStatus.Unmarshal(m, b)
}
func (m *InteractionStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InteractionStatus.Marshal(b, m, deterministic)
}
func (m *InteractionStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InteractionStatus.Merge(m, src)
}
func (m *InteractionStatus) XXX_Size() int {
	return xxx_messageInfo_InteractionStatus.Size(m)
}
func (m *InteractionStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_InteractionStatus.DiscardUnknown(m)
}

var xxx_messageInfo_InteractionStatus proto.InternalMessageInfo

func (m *InteractionStatus) GetCode() InteractionStatusCode {
	if m != nil {
		return m.Code
	}
	return InteractionStatusCode_Unknown
}

type InteractionMessage struct {
	Frame                *Frame   `protobuf:"bytes,1,opt,name=frame,proto3" json:"frame,omitempty"`
	InteractionElements  []byte   `protobuf:"bytes,2,opt,name=interactionElements,proto3" json:"interactionElements,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InteractionMessage) Reset()         { *m = InteractionMessage{} }
func (m *InteractionMessage) String() string { return proto.CompactTextString(m) }
func (*InteractionMessage) ProtoMessage()    {}
func (*InteractionMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_43fb9bc5eb13a2f0, []int{1}
}

func (m *InteractionMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InteractionMessage.Unmarshal(m, b)
}
func (m *InteractionMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InteractionMessage.Marshal(b, m, deterministic)
}
func (m *InteractionMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InteractionMessage.Merge(m, src)
}
func (m *InteractionMessage) XXX_Size() int {
	return xxx_messageInfo_InteractionMessage.Size(m)
}
func (m *InteractionMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_InteractionMessage.DiscardUnknown(m)
}

var xxx_messageInfo_InteractionMessage proto.InternalMessageInfo

func (m *InteractionMessage) GetFrame() *Frame {
	if m != nil {
		return m.Frame
	}
	return nil
}

func (m *InteractionMessage) GetInteractionElements() []byte {
	if m != nil {
		return m.InteractionElements
	}
	return nil
}

type Frame struct {
	SemanticProtocol     string              `protobuf:"bytes,1,opt,name=semanticProtocol,proto3" json:"semanticProtocol,omitempty"`
	Type                 string              `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	MessageId            string              `protobuf:"bytes,3,opt,name=messageId,proto3" json:"messageId,omitempty"`
	ReplyBy              uint32              `protobuf:"varint,4,opt,name=replyBy,proto3" json:"replyBy,omitempty"`
	Receiver             *ConversationMember `protobuf:"bytes,5,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Sender               *ConversationMember `protobuf:"bytes,6,opt,name=sender,proto3" json:"sender,omitempty"`
	ConversationId       string              `protobuf:"bytes,7,opt,name=conversationId,proto3" json:"conversationId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Frame) Reset()         { *m = Frame{} }
func (m *Frame) String() string { return proto.CompactTextString(m) }
func (*Frame) ProtoMessage()    {}
func (*Frame) Descriptor() ([]byte, []int) {
	return fileDescriptor_43fb9bc5eb13a2f0, []int{2}
}

func (m *Frame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Frame.Unmarshal(m, b)
}
func (m *Frame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Frame.Marshal(b, m, deterministic)
}
func (m *Frame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Frame.Merge(m, src)
}
func (m *Frame) XXX_Size() int {
	return xxx_messageInfo_Frame.Size(m)
}
func (m *Frame) XXX_DiscardUnknown() {
	xxx_messageInfo_Frame.DiscardUnknown(m)
}

var xxx_messageInfo_Frame proto.InternalMessageInfo

func (m *Frame) GetSemanticProtocol() string {
	if m != nil {
		return m.SemanticProtocol
	}
	return ""
}

func (m *Frame) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Frame) GetMessageId() string {
	if m != nil {
		return m.MessageId
	}
	return ""
}

func (m *Frame) GetReplyBy() uint32 {
	if m != nil {
		return m.ReplyBy
	}
	return 0
}

func (m *Frame) GetReceiver() *ConversationMember {
	if m != nil {
		return m.Receiver
	}
	return nil
}

func (m *Frame) GetSender() *ConversationMember {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *Frame) GetConversationId() string {
	if m != nil {
		return m.ConversationId
	}
	return ""
}

type ConversationMember struct {
	Identifier           *Identifier `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Role                 *Role       `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ConversationMember) Reset()         { *m = ConversationMember{} }
func (m *ConversationMember) String() string { return proto.CompactTextString(m) }
func (*ConversationMember) ProtoMessage()    {}
func (*ConversationMember) Descriptor() ([]byte, []int) {
	return fileDescriptor_43fb9bc5eb13a2f0, []int{3}
}

func (m *ConversationMember) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConversationMember.Unmarshal(m, b)
}
func (m *ConversationMember) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConversationMember.Marshal(b, m, deterministic)
}
func (m *ConversationMember) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConversationMember.Merge(m, src)
}
func (m *ConversationMember) XXX_Size() int {
	return xxx_messageInfo_ConversationMember.Size(m)
}
func (m *ConversationMember) XXX_DiscardUnknown() {
	xxx_messageInfo_ConversationMember.DiscardUnknown(m)
}

var xxx_messageInfo_ConversationMember proto.InternalMessageInfo

func (m *ConversationMember) GetIdentifier() *Identifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *ConversationMember) GetRole() *Role {
	if m != nil {
		return m.Role
	}
	return nil
}

type Role struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Role) Reset()         { *m = Role{} }
func (m *Role) String() string { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()    {}
func (*Role) Descriptor() ([]byte, []int) {
	return fileDescriptor_43fb9bc5eb13a2f0, []int{4}
}

func (m *Role) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Role.Unmarshal(m, b)
}
func (m *Role) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Role.Marshal(b, m, deterministic)
}
func (m *Role) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Role.Merge(m, src)
}
func (m *Role) XXX_Size() int {
	return xxx_messageInfo_Role.Size(m)
}
func (m *Role) XXX_DiscardUnknown() {
	xxx_messageInfo_Role.DiscardUnknown(m)
}

var xxx_messageInfo_Role proto.InternalMessageInfo

func (m *Role) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Identifier struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	IdType               string   `protobuf:"bytes,2,opt,name=idType,proto3" json:"idType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Identifier) Reset()         { *m = Identifier{} }
func (m *Identifier) String() string { return proto.CompactTextString(m) }
func (*Identifier) ProtoMessage()    {}
func (*Identifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_43fb9bc5eb13a2f0, []int{5}
}

func (m *Identifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Identifier.Unmarshal(m, b)
}
func (m *Identifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Identifier.Marshal(b, m, deterministic)
}
func (m *Identifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Identifier.Merge(m, src)
}
func (m *Identifier) XXX_Size() int {
	return xxx_messageInfo_Identifier.Size(m)
}
func (m *Identifier) XXX_DiscardUnknown() {
	xxx_messageInfo_Identifier.DiscardUnknown(m)
}

var xxx_messageInfo_Identifier proto.InternalMessageInfo

func (m *Identifier) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Identifier) GetIdType() string {
	if m != nil {
		return m.IdType
	}
	return ""
}

func init() {
	proto.RegisterEnum("interaction.InteractionStatusCode", InteractionStatusCode_name, InteractionStatusCode_value)
	proto.RegisterType((*InteractionStatus)(nil), "interaction.InteractionStatus")
	proto.RegisterType((*InteractionMessage)(nil), "interaction.InteractionMessage")
	proto.RegisterType((*Frame)(nil), "interaction.Frame")
	proto.RegisterType((*ConversationMember)(nil), "interaction.ConversationMember")
	proto.RegisterType((*Role)(nil), "interaction.Role")
	proto.RegisterType((*Identifier)(nil), "interaction.Identifier")
}

func init() { proto.RegisterFile("interaction.proto", fileDescriptor_43fb9bc5eb13a2f0) }

var fileDescriptor_43fb9bc5eb13a2f0 = []byte{
	// 451 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xae, 0xfd, 0x73, 0x9c, 0x5f, 0x26, 0x10, 0x25, 0x83, 0x80, 0x55, 0x85, 0xda, 0xc8, 0x12,
	0xc8, 0xea, 0xa1, 0x42, 0x01, 0x51, 0x24, 0x6e, 0x54, 0x54, 0x8a, 0x10, 0x02, 0x6d, 0xe8, 0x91,
	0xc3, 0xd6, 0x3b, 0x45, 0xab, 0xda, 0xbb, 0xd6, 0x7a, 0x09, 0xca, 0xc3, 0xf0, 0x62, 0x3c, 0x0d,
	0xf2, 0xda, 0x6d, 0x6c, 0xd2, 0x22, 0x0e, 0xdc, 0x76, 0xe6, 0xfb, 0xe6, 0x9b, 0x9d, 0x7f, 0x30,
	0x53, 0xda, 0x91, 0x15, 0x99, 0x53, 0x46, 0x1f, 0x97, 0xd6, 0x38, 0x83, 0xe3, 0x8e, 0x2b, 0x79,
	0x0f, 0xb3, 0xe5, 0xd6, 0x5c, 0x39, 0xe1, 0xbe, 0x55, 0xf8, 0x0a, 0xa2, 0x53, 0x23, 0x89, 0x05,
	0xf3, 0x20, 0x9d, 0x2c, 0x92, 0xe3, 0xae, 0xc6, 0x0e, 0xbb, 0x66, 0x72, 0xcf, 0x4f, 0x4a, 0xc0,
	0x0e, 0xfc, 0x81, 0xaa, 0x4a, 0x7c, 0x25, 0x4c, 0x61, 0x70, 0x69, 0x45, 0xd1, 0xc8, 0x8d, 0x17,
	0xd8, 0x93, 0x3b, 0xab, 0x11, 0xde, 0x10, 0xf0, 0x39, 0x3c, 0xe8, 0x60, 0xef, 0x72, 0x2a, 0x48,
	0xbb, 0x8a, 0x85, 0xf3, 0x20, 0xbd, 0xc7, 0x6f, 0x83, 0x92, 0x1f, 0x21, 0x0c, 0xbc, 0x04, 0x1e,
	0xc1, 0xb4, 0xa2, 0x42, 0x68, 0xa7, 0xb2, 0x4f, 0x75, 0x99, 0x99, 0xc9, 0x7d, 0xc2, 0x11, 0xdf,
	0xf1, 0x23, 0x42, 0xe4, 0x36, 0x25, 0x79, 0xe1, 0x11, 0xf7, 0x6f, 0x7c, 0x02, 0xa3, 0xa2, 0xf9,
	0xf0, 0x52, 0xb2, 0xff, 0x3c, 0xb0, 0x75, 0x20, 0x83, 0xa1, 0xa5, 0x32, 0xdf, 0xbc, 0xdd, 0xb0,
	0x68, 0x1e, 0xa4, 0xf7, 0xf9, 0xb5, 0x89, 0x6f, 0xe0, 0x7f, 0x4b, 0x19, 0xa9, 0x35, 0x59, 0x36,
	0xf0, 0x05, 0x1e, 0xf6, 0x0a, 0x3c, 0x35, 0x7a, 0x4d, 0xb6, 0x12, 0x4d, 0x47, 0x8a, 0x0b, 0xb2,
	0xfc, 0x26, 0x00, 0x4f, 0x20, 0xae, 0x48, 0x4b, 0xb2, 0x2c, 0xfe, 0xbb, 0xd0, 0x96, 0x8e, 0xcf,
	0x60, 0x92, 0x75, 0xd0, 0xa5, 0x64, 0x43, 0xff, 0xe5, 0xdf, 0xbc, 0x89, 0x03, 0xdc, 0x55, 0xc1,
	0x13, 0x00, 0x25, 0x49, 0x3b, 0x75, 0xa9, 0xc8, 0xb6, 0x63, 0x79, 0xdc, 0x9f, 0xf2, 0x0d, 0xcc,
	0x3b, 0x54, 0x7c, 0x0a, 0x91, 0x35, 0x79, 0xd3, 0xb8, 0xf1, 0x62, 0xd6, 0x0b, 0xe1, 0x26, 0x27,
	0xee, 0xe1, 0x64, 0x1f, 0xa2, 0xda, 0xaa, 0xfb, 0xac, 0xaf, 0x07, 0x3f, 0xe2, 0xfe, 0x9d, 0xbc,
	0x04, 0xd8, 0x8a, 0xe3, 0x04, 0x42, 0x25, 0x5b, 0x3c, 0x54, 0x12, 0x1f, 0x41, 0xac, 0xe4, 0xe7,
	0xed, 0x6c, 0x5a, 0xeb, 0xe8, 0x35, 0x3c, 0xbc, 0x75, 0xf1, 0x70, 0x0c, 0xc3, 0x73, 0x7d, 0xa5,
	0xcd, 0x77, 0x3d, 0xdd, 0xc3, 0x18, 0xc2, 0x8f, 0x57, 0xd3, 0x00, 0x01, 0xe2, 0x33, 0xa1, 0x72,
	0x92, 0xd3, 0x70, 0xf1, 0x33, 0xe8, 0x2d, 0xe5, 0x8a, 0xec, 0x5a, 0x65, 0x84, 0x04, 0x07, 0xe7,
	0x65, 0x6e, 0x84, 0xdc, 0x5d, 0xd8, 0x95, 0xb3, 0x24, 0x0a, 0x3c, 0xbc, 0x6b, 0xed, 0x5b, 0xda,
	0xfe, 0xc1, 0x9f, 0xef, 0x22, 0xd9, 0x4b, 0x03, 0xfc, 0x02, 0xec, 0xae, 0x34, 0xff, 0x20, 0xc1,
	0x45, 0xec, 0x2f, 0xfa, 0xc5, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x3f, 0xab, 0xba, 0xe6,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// InteractionServiceClient is the client API for InteractionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InteractionServiceClient interface {
	UploadInteractionMessageStream(ctx context.Context, opts ...grpc.CallOption) (InteractionService_UploadInteractionMessageStreamClient, error)
	UploadInteractionMessage(ctx context.Context, in *InteractionMessage, opts ...grpc.CallOption) (*InteractionStatus, error)
}

type interactionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInteractionServiceClient(cc grpc.ClientConnInterface) InteractionServiceClient {
	return &interactionServiceClient{cc}
}

func (c *interactionServiceClient) UploadInteractionMessageStream(ctx context.Context, opts ...grpc.CallOption) (InteractionService_UploadInteractionMessageStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_InteractionService_serviceDesc.Streams[0], "/interaction.InteractionService/UploadInteractionMessageStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &interactionServiceUploadInteractionMessageStreamClient{stream}
	return x, nil
}

type InteractionService_UploadInteractionMessageStreamClient interface {
	Send(*InteractionMessage) error
	CloseAndRecv() (*InteractionStatus, error)
	grpc.ClientStream
}

type interactionServiceUploadInteractionMessageStreamClient struct {
	grpc.ClientStream
}

func (x *interactionServiceUploadInteractionMessageStreamClient) Send(m *InteractionMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *interactionServiceUploadInteractionMessageStreamClient) CloseAndRecv() (*InteractionStatus, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(InteractionStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *interactionServiceClient) UploadInteractionMessage(ctx context.Context, in *InteractionMessage, opts ...grpc.CallOption) (*InteractionStatus, error) {
	out := new(InteractionStatus)
	err := c.cc.Invoke(ctx, "/interaction.InteractionService/UploadInteractionMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InteractionServiceServer is the server API for InteractionService service.
type InteractionServiceServer interface {
	UploadInteractionMessageStream(InteractionService_UploadInteractionMessageStreamServer) error
	UploadInteractionMessage(context.Context, *InteractionMessage) (*InteractionStatus, error)
}

// UnimplementedInteractionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedInteractionServiceServer struct {
}

func (*UnimplementedInteractionServiceServer) UploadInteractionMessageStream(srv InteractionService_UploadInteractionMessageStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadInteractionMessageStream not implemented")
}
func (*UnimplementedInteractionServiceServer) UploadInteractionMessage(ctx context.Context, req *InteractionMessage) (*InteractionStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadInteractionMessage not implemented")
}

func RegisterInteractionServiceServer(s *grpc.Server, srv InteractionServiceServer) {
	s.RegisterService(&_InteractionService_serviceDesc, srv)
}

func _InteractionService_UploadInteractionMessageStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(InteractionServiceServer).UploadInteractionMessageStream(&interactionServiceUploadInteractionMessageStreamServer{stream})
}

type InteractionService_UploadInteractionMessageStreamServer interface {
	SendAndClose(*InteractionStatus) error
	Recv() (*InteractionMessage, error)
	grpc.ServerStream
}

type interactionServiceUploadInteractionMessageStreamServer struct {
	grpc.ServerStream
}

func (x *interactionServiceUploadInteractionMessageStreamServer) SendAndClose(m *InteractionStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *interactionServiceUploadInteractionMessageStreamServer) Recv() (*InteractionMessage, error) {
	m := new(InteractionMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _InteractionService_UploadInteractionMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InteractionMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionServiceServer).UploadInteractionMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interaction.InteractionService/UploadInteractionMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionServiceServer).UploadInteractionMessage(ctx, req.(*InteractionMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _InteractionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "interaction.InteractionService",
	HandlerType: (*InteractionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadInteractionMessage",
			Handler:    _InteractionService_UploadInteractionMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadInteractionMessageStream",
			Handler:       _InteractionService_UploadInteractionMessageStream_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "interaction.proto",
}
