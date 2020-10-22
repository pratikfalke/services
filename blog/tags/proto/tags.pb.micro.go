// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: tags.proto

package tags

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Tags service

func NewTagsEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Tags service

type TagsService interface {
	// Add a tag to a resource
	Add(ctx context.Context, in *AddRequest, opts ...client.CallOption) (*AddResponse, error)
	// Remove a tag from a resource
	Remove(ctx context.Context, in *RemoveRequest, opts ...client.CallOption) (*RemoveResponse, error)
	// List tags by
	List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
	// Change properties of a tag, currently only the title and description
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
}

type tagsService struct {
	c    client.Client
	name string
}

func NewTagsService(name string, c client.Client) TagsService {
	return &tagsService{
		c:    c,
		name: name,
	}
}

func (c *tagsService) Add(ctx context.Context, in *AddRequest, opts ...client.CallOption) (*AddResponse, error) {
	req := c.c.NewRequest(c.name, "Tags.Add", in)
	out := new(AddResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagsService) Remove(ctx context.Context, in *RemoveRequest, opts ...client.CallOption) (*RemoveResponse, error) {
	req := c.c.NewRequest(c.name, "Tags.Remove", in)
	out := new(RemoveResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagsService) List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "Tags.List", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagsService) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.name, "Tags.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Tags service

type TagsHandler interface {
	// Add a tag to a resource
	Add(context.Context, *AddRequest, *AddResponse) error
	// Remove a tag from a resource
	Remove(context.Context, *RemoveRequest, *RemoveResponse) error
	// List tags by
	List(context.Context, *ListRequest, *ListResponse) error
	// Change properties of a tag, currently only the title and description
	Update(context.Context, *UpdateRequest, *UpdateResponse) error
}

func RegisterTagsHandler(s server.Server, hdlr TagsHandler, opts ...server.HandlerOption) error {
	type tags interface {
		Add(ctx context.Context, in *AddRequest, out *AddResponse) error
		Remove(ctx context.Context, in *RemoveRequest, out *RemoveResponse) error
		List(ctx context.Context, in *ListRequest, out *ListResponse) error
		Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error
	}
	type Tags struct {
		tags
	}
	h := &tagsHandler{hdlr}
	return s.Handle(s.NewHandler(&Tags{h}, opts...))
}

type tagsHandler struct {
	TagsHandler
}

func (h *tagsHandler) Add(ctx context.Context, in *AddRequest, out *AddResponse) error {
	return h.TagsHandler.Add(ctx, in, out)
}

func (h *tagsHandler) Remove(ctx context.Context, in *RemoveRequest, out *RemoveResponse) error {
	return h.TagsHandler.Remove(ctx, in, out)
}

func (h *tagsHandler) List(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.TagsHandler.List(ctx, in, out)
}

func (h *tagsHandler) Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.TagsHandler.Update(ctx, in, out)
}
