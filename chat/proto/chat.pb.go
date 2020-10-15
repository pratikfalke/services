// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/chat.proto

package chat

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// NewRequest contains the infromation needed to create a new chat
type NewRequest struct {
	UserIds              []string `protobuf:"bytes,1,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewRequest) Reset()         { *m = NewRequest{} }
func (m *NewRequest) String() string { return proto.CompactTextString(m) }
func (*NewRequest) ProtoMessage()    {}
func (*NewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed7e7dde45555b7d, []int{0}
}

func (m *NewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewRequest.Unmarshal(m, b)
}
func (m *NewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewRequest.Marshal(b, m, deterministic)
}
func (m *NewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewRequest.Merge(m, src)
}
func (m *NewRequest) XXX_Size() int {
	return xxx_messageInfo_NewRequest.Size(m)
}
func (m *NewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewRequest proto.InternalMessageInfo

func (m *NewRequest) GetUserIds() []string {
	if m != nil {
		return m.UserIds
	}
	return nil
}

// NewResponse contains the chat id for the users
type NewResponse struct {
	ChatId               string   `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewResponse) Reset()         { *m = NewResponse{} }
func (m *NewResponse) String() string { return proto.CompactTextString(m) }
func (*NewResponse) ProtoMessage()    {}
func (*NewResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed7e7dde45555b7d, []int{1}
}

func (m *NewResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewResponse.Unmarshal(m, b)
}
func (m *NewResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewResponse.Marshal(b, m, deterministic)
}
func (m *NewResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewResponse.Merge(m, src)
}
func (m *NewResponse) XXX_Size() int {
	return xxx_messageInfo_NewResponse.Size(m)
}
func (m *NewResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NewResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NewResponse proto.InternalMessageInfo

func (m *NewResponse) GetChatId() string {
	if m != nil {
		return m.ChatId
	}
	return ""
}

// HistoryRequest contains the id of the chat we want the history for. This RPC will return all
// historical messages, however in a real life application we'd introduce some form of pagination
// here, only loading the older messages when required.
type HistoryRequest struct {
	ChatId               string   `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HistoryRequest) Reset()         { *m = HistoryRequest{} }
func (m *HistoryRequest) String() string { return proto.CompactTextString(m) }
func (*HistoryRequest) ProtoMessage()    {}
func (*HistoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed7e7dde45555b7d, []int{2}
}

func (m *HistoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HistoryRequest.Unmarshal(m, b)
}
func (m *HistoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HistoryRequest.Marshal(b, m, deterministic)
}
func (m *HistoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HistoryRequest.Merge(m, src)
}
func (m *HistoryRequest) XXX_Size() int {
	return xxx_messageInfo_HistoryRequest.Size(m)
}
func (m *HistoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HistoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HistoryRequest proto.InternalMessageInfo

func (m *HistoryRequest) GetChatId() string {
	if m != nil {
		return m.ChatId
	}
	return ""
}

// HistoryResponse contains the historical messages in a chat
type HistoryResponse struct {
	Messages             []*Message `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *HistoryResponse) Reset()         { *m = HistoryResponse{} }
func (m *HistoryResponse) String() string { return proto.CompactTextString(m) }
func (*HistoryResponse) ProtoMessage()    {}
func (*HistoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed7e7dde45555b7d, []int{3}
}

func (m *HistoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HistoryResponse.Unmarshal(m, b)
}
func (m *HistoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HistoryResponse.Marshal(b, m, deterministic)
}
func (m *HistoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HistoryResponse.Merge(m, src)
}
func (m *HistoryResponse) XXX_Size() int {
	return xxx_messageInfo_HistoryResponse.Size(m)
}
func (m *HistoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HistoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HistoryResponse proto.InternalMessageInfo

func (m *HistoryResponse) GetMessages() []*Message {
	if m != nil {
		return m.Messages
	}
	return nil
}

// Message sent to a chat
type Message struct {
	// id of the message, allocated by the server
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// a client side id, should be validated by the server to make the request retry safe
	ClientId string `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// id of the chat the message is being sent to / from
	ChatId string `protobuf:"bytes,3,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	// id of the user who sent the message
	UserId string `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// time time the message was sent in unix format
	SentAt int32 `protobuf:"varint,5,opt,name=sent_at,json=sentAt,proto3" json:"sent_at,omitempty"`
	// subject of the message
	Subject string `protobuf:"bytes,6,opt,name=subject,proto3" json:"subject,omitempty"`
	// text of the message
	Text                 string   `protobuf:"bytes,7,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed7e7dde45555b7d, []int{4}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Message) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *Message) GetChatId() string {
	if m != nil {
		return m.ChatId
	}
	return ""
}

func (m *Message) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Message) GetSentAt() int32 {
	if m != nil {
		return m.SentAt
	}
	return 0
}

func (m *Message) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *Message) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*NewRequest)(nil), "chat.NewRequest")
	proto.RegisterType((*NewResponse)(nil), "chat.NewResponse")
	proto.RegisterType((*HistoryRequest)(nil), "chat.HistoryRequest")
	proto.RegisterType((*HistoryResponse)(nil), "chat.HistoryResponse")
	proto.RegisterType((*Message)(nil), "chat.Message")
}

func init() { proto.RegisterFile("proto/chat.proto", fileDescriptor_ed7e7dde45555b7d) }

var fileDescriptor_ed7e7dde45555b7d = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x41, 0x4f, 0xf2, 0x40,
	0x10, 0xcd, 0x42, 0x69, 0x61, 0xc8, 0xc7, 0x87, 0x1b, 0x8d, 0x2b, 0x5e, 0x48, 0x0f, 0x5a, 0x24,
	0xa1, 0x06, 0x13, 0x2f, 0x7a, 0x51, 0x2e, 0x72, 0x90, 0x43, 0x8f, 0x5e, 0x48, 0x69, 0x37, 0xb0,
	0x46, 0xba, 0xd8, 0xd9, 0x8a, 0xfe, 0x12, 0x7f, 0x86, 0x7f, 0xd1, 0xec, 0x6e, 0x41, 0x30, 0xf1,
	0x36, 0xef, 0xbd, 0x99, 0xd9, 0x37, 0x2f, 0x0b, 0xed, 0x55, 0x2e, 0x95, 0x0c, 0x93, 0x45, 0xac,
	0x06, 0xa6, 0xa4, 0x8e, 0xae, 0xfd, 0x73, 0x80, 0x09, 0x5f, 0x47, 0xfc, 0xb5, 0xe0, 0xa8, 0xe8,
	0x09, 0xd4, 0x0b, 0xe4, 0xf9, 0x54, 0xa4, 0xc8, 0x48, 0xb7, 0x1a, 0x34, 0x22, 0x4f, 0xe3, 0x71,
	0x8a, 0xfe, 0x19, 0x34, 0x4d, 0x23, 0xae, 0x64, 0x86, 0x9c, 0x1e, 0x83, 0xa7, 0xe7, 0xa7, 0x22,
	0x65, 0xa4, 0x4b, 0x82, 0x46, 0xe4, 0x6a, 0x38, 0x4e, 0xfd, 0x1e, 0xb4, 0x1e, 0x04, 0x2a, 0x99,
	0x7f, 0x6c, 0x96, 0xfe, 0xd9, 0x7a, 0x0b, 0xff, 0xb7, 0xad, 0xe5, 0xda, 0x1e, 0xd4, 0x97, 0x1c,
	0x31, 0x9e, 0x73, 0x6b, 0xa0, 0x39, 0xfc, 0x37, 0x30, 0x9e, 0x1f, 0x2d, 0x1b, 0x6d, 0x65, 0xff,
	0x8b, 0x80, 0x57, 0xb2, 0xb4, 0x05, 0x95, 0xed, 0xf6, 0x8a, 0x48, 0xe9, 0x29, 0x34, 0x92, 0x17,
	0xc1, 0x33, 0xf3, 0x68, 0xc5, 0xd0, 0x75, 0x4b, 0x8c, 0xd3, 0x5d, 0x3f, 0xd5, 0x5d, 0x3f, 0x5a,
	0x28, 0xaf, 0x67, 0x8e, 0x15, 0xec, 0xf1, 0x5a, 0x40, 0xbd, 0x2c, 0x56, 0xac, 0xd6, 0x25, 0x41,
	0x2d, 0x72, 0x35, 0xbc, 0x53, 0x94, 0x81, 0x87, 0xc5, 0xec, 0x99, 0x27, 0x8a, 0xb9, 0x66, 0x62,
	0x03, 0x29, 0x05, 0x47, 0xf1, 0x77, 0xc5, 0x3c, 0x43, 0x9b, 0x7a, 0xf8, 0x49, 0xc0, 0x19, 0x2d,
	0x62, 0x45, 0x2f, 0xa0, 0x3a, 0xe1, 0x6b, 0xda, 0xb6, 0xa7, 0xfd, 0xe4, 0xdf, 0x39, 0xd8, 0x61,
	0xca, 0x44, 0xae, 0xc1, 0x2b, 0x43, 0xa2, 0x87, 0x56, 0xdd, 0x8f, 0xb7, 0x73, 0xf4, 0x8b, 0x2d,
	0xe7, 0xfa, 0xe0, 0x8d, 0x64, 0x96, 0x69, 0x2f, 0xfb, 0x11, 0x76, 0xf6, 0x61, 0x40, 0x2e, 0xc9,
	0x7d, 0xff, 0xa9, 0x37, 0x17, 0x6a, 0x51, 0xcc, 0x06, 0x89, 0x5c, 0x86, 0x4b, 0x91, 0xe4, 0x32,
	0x44, 0x9e, 0xbf, 0x89, 0x84, 0xa3, 0xf9, 0x33, 0xa1, 0xf9, 0x33, 0x37, 0xba, 0x9c, 0xb9, 0xa6,
	0xbe, 0xfa, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x71, 0x1b, 0x44, 0xc0, 0x53, 0x02, 0x00, 0x00,
}
