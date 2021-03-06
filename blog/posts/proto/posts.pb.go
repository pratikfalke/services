// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/posts.proto

package posts

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

type Post struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Slug                 string   `protobuf:"bytes,3,opt,name=slug,proto3" json:"slug,omitempty"`
	Content              string   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Created              int64    `protobuf:"varint,5,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64    `protobuf:"varint,6,opt,name=updated,proto3" json:"updated,omitempty"`
	Author               string   `protobuf:"bytes,7,opt,name=author,proto3" json:"author,omitempty"`
	Tags                 []string `protobuf:"bytes,8,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Post) Reset()         { *m = Post{} }
func (m *Post) String() string { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()    {}
func (*Post) Descriptor() ([]byte, []int) {
	return fileDescriptor_e93dc7d934d9dc10, []int{0}
}

func (m *Post) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Post.Unmarshal(m, b)
}
func (m *Post) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Post.Marshal(b, m, deterministic)
}
func (m *Post) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Post.Merge(m, src)
}
func (m *Post) XXX_Size() int {
	return xxx_messageInfo_Post.Size(m)
}
func (m *Post) XXX_DiscardUnknown() {
	xxx_messageInfo_Post.DiscardUnknown(m)
}

var xxx_messageInfo_Post proto.InternalMessageInfo

func (m *Post) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Post) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Post) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *Post) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Post) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Post) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *Post) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *Post) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

// Query posts. Acts as a listing when no id or slug provided.
// Gets a single post by id or slug if any of them provided.
type QueryRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Slug                 string   `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug,omitempty"`
	Tag                  string   `protobuf:"bytes,3,opt,name=tag,proto3" json:"tag,omitempty"`
	Offset               int64    `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                int64    `protobuf:"varint,5,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e93dc7d934d9dc10, []int{1}
}

func (m *QueryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryRequest.Unmarshal(m, b)
}
func (m *QueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryRequest.Marshal(b, m, deterministic)
}
func (m *QueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRequest.Merge(m, src)
}
func (m *QueryRequest) XXX_Size() int {
	return xxx_messageInfo_QueryRequest.Size(m)
}
func (m *QueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRequest proto.InternalMessageInfo

func (m *QueryRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *QueryRequest) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *QueryRequest) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *QueryRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *QueryRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type QueryResponse struct {
	Posts                []*Post  `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryResponse) Reset()         { *m = QueryResponse{} }
func (m *QueryResponse) String() string { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()    {}
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e93dc7d934d9dc10, []int{2}
}

func (m *QueryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryResponse.Unmarshal(m, b)
}
func (m *QueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryResponse.Marshal(b, m, deterministic)
}
func (m *QueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryResponse.Merge(m, src)
}
func (m *QueryResponse) XXX_Size() int {
	return xxx_messageInfo_QueryResponse.Size(m)
}
func (m *QueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryResponse proto.InternalMessageInfo

func (m *QueryResponse) GetPosts() []*Post {
	if m != nil {
		return m.Posts
	}
	return nil
}

type SaveRequest struct {
	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title     string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Slug      string `protobuf:"bytes,3,opt,name=slug,proto3" json:"slug,omitempty"`
	Content   string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Timestamp int64  `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// When updating a post and wanting to delete all tags,
	// send a list of tags with only one member being an empty string [""]
	Tags                 []string `protobuf:"bytes,6,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveRequest) Reset()         { *m = SaveRequest{} }
func (m *SaveRequest) String() string { return proto.CompactTextString(m) }
func (*SaveRequest) ProtoMessage()    {}
func (*SaveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e93dc7d934d9dc10, []int{3}
}

func (m *SaveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveRequest.Unmarshal(m, b)
}
func (m *SaveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveRequest.Marshal(b, m, deterministic)
}
func (m *SaveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveRequest.Merge(m, src)
}
func (m *SaveRequest) XXX_Size() int {
	return xxx_messageInfo_SaveRequest.Size(m)
}
func (m *SaveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SaveRequest proto.InternalMessageInfo

func (m *SaveRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SaveRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SaveRequest) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *SaveRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *SaveRequest) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *SaveRequest) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type SaveResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveResponse) Reset()         { *m = SaveResponse{} }
func (m *SaveResponse) String() string { return proto.CompactTextString(m) }
func (*SaveResponse) ProtoMessage()    {}
func (*SaveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e93dc7d934d9dc10, []int{4}
}

func (m *SaveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveResponse.Unmarshal(m, b)
}
func (m *SaveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveResponse.Marshal(b, m, deterministic)
}
func (m *SaveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveResponse.Merge(m, src)
}
func (m *SaveResponse) XXX_Size() int {
	return xxx_messageInfo_SaveResponse.Size(m)
}
func (m *SaveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SaveResponse proto.InternalMessageInfo

func (m *SaveResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e93dc7d934d9dc10, []int{5}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e93dc7d934d9dc10, []int{6}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Post)(nil), "posts.Post")
	proto.RegisterType((*QueryRequest)(nil), "posts.QueryRequest")
	proto.RegisterType((*QueryResponse)(nil), "posts.QueryResponse")
	proto.RegisterType((*SaveRequest)(nil), "posts.SaveRequest")
	proto.RegisterType((*SaveResponse)(nil), "posts.SaveResponse")
	proto.RegisterType((*DeleteRequest)(nil), "posts.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "posts.DeleteResponse")
}

func init() {
	proto.RegisterFile("proto/posts.proto", fileDescriptor_e93dc7d934d9dc10)
}

var fileDescriptor_e93dc7d934d9dc10 = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x53, 0x5d, 0x4e, 0xf2, 0x50,
	0x10, 0xfd, 0x4a, 0x7f, 0xf8, 0x18, 0x7e, 0x82, 0x03, 0x9a, 0x1b, 0x62, 0x14, 0xfb, 0xc4, 0x13,
	0x46, 0x34, 0x71, 0x03, 0x2e, 0x40, 0xeb, 0x0a, 0xaa, 0x5c, 0xb0, 0x49, 0xe1, 0xd6, 0xde, 0xa9,
	0x89, 0xeb, 0x70, 0x15, 0x6e, 0xc1, 0xd5, 0xd9, 0xfb, 0x53, 0x6c, 0x89, 0xbc, 0xf9, 0x36, 0xe7,
	0x9c, 0xce, 0xcc, 0x99, 0xd3, 0x5c, 0x38, 0xca, 0x72, 0x41, 0xe2, 0x32, 0x13, 0x92, 0xe4, 0x5c,
	0xd7, 0xe8, 0x6b, 0x10, 0x7e, 0x39, 0xe0, 0xdd, 0x97, 0x15, 0x0e, 0xa0, 0x95, 0x2c, 0x99, 0x33,
	0x75, 0x66, 0x9d, 0xa8, 0xac, 0x70, 0x0c, 0x3e, 0x25, 0x94, 0x72, 0xd6, 0xd2, 0x94, 0x01, 0x88,
	0xe0, 0xc9, 0xb4, 0x58, 0x33, 0x57, 0x93, 0xba, 0x46, 0x06, 0xed, 0x67, 0xb1, 0x25, 0xbe, 0x25,
	0xe6, 0x69, 0xba, 0x82, 0x5a, 0xc9, 0x79, 0x4c, 0x7c, 0xc9, 0xfc, 0x52, 0x71, 0xa3, 0x0a, 0x2a,
	0xa5, 0xc8, 0x96, 0x5a, 0x09, 0x8c, 0x62, 0x21, 0x9e, 0x40, 0x10, 0x17, 0xf4, 0x22, 0x72, 0xd6,
	0xd6, 0xc3, 0x2c, 0x52, 0x9b, 0x29, 0x5e, 0x4b, 0xf6, 0x7f, 0xea, 0xaa, 0xcd, 0xaa, 0x0e, 0x73,
	0xe8, 0x3d, 0x14, 0x3c, 0x7f, 0x8f, 0xf8, 0x6b, 0xc1, 0x7f, 0xb9, 0xa1, 0x72, 0xdb, 0xaa, 0xb9,
	0x1d, 0x82, 0x5b, 0xf6, 0xda, 0x03, 0x54, 0xa9, 0x36, 0x8a, 0xd5, 0x4a, 0x72, 0x63, 0xdf, 0x8d,
	0x2c, 0x52, 0x09, 0xa4, 0xc9, 0x26, 0x21, 0xeb, 0xdd, 0x80, 0x70, 0x01, 0x7d, 0xbb, 0x53, 0x66,
	0x62, 0x2b, 0x39, 0x5e, 0x80, 0x89, 0xb2, 0xdc, 0xeb, 0xce, 0xba, 0x8b, 0xee, 0xdc, 0xa4, 0xac,
	0x42, 0x8d, 0x6c, 0xc8, 0x1f, 0x0e, 0x74, 0x1f, 0xe3, 0x37, 0x7e, 0xc8, 0xe7, 0x5f, 0x64, 0x7d,
	0x0a, 0x1d, 0x4a, 0x36, 0xe5, 0xf4, 0x78, 0x93, 0x59, 0xc7, 0x3f, 0xc4, 0x2e, 0xbd, 0xa0, 0x96,
	0xde, 0x19, 0xf4, 0x8c, 0x29, 0x7b, 0xc8, 0x9e, 0xab, 0xf0, 0x1c, 0xfa, 0x77, 0x3c, 0xe5, 0x74,
	0xc8, 0x76, 0x38, 0x84, 0x41, 0xf5, 0x81, 0x19, 0xb1, 0xf8, 0x74, 0xc0, 0x57, 0x87, 0x4b, 0xbc,
	0x01, 0x5f, 0xc7, 0x84, 0x23, 0x9b, 0x47, 0xfd, 0x47, 0x4d, 0xc6, 0x4d, 0xd2, 0x74, 0x87, 0xff,
	0xf0, 0x0a, 0x3c, 0x65, 0x09, 0xd1, 0xea, 0xb5, 0xd0, 0x26, 0xa3, 0x06, 0xb7, 0x6b, 0xb9, 0x85,
	0xc0, 0x98, 0xc0, 0x6a, 0x68, 0xc3, 0xf4, 0xe4, 0x78, 0x8f, 0xad, 0x1a, 0x9f, 0x02, 0xfd, 0x0e,
	0xae, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x8c, 0x67, 0xbb, 0x1c, 0x03, 0x00, 0x00,
}
