// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/user/user.proto

package user_lms

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Api Endpoints for User service

func NewUserEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for User service

type UserService interface {
	Create(ctx context.Context, in *NewUserRequest, opts ...client.CallOption) (*UserResponse, error)
	GetByID(ctx context.Context, in *UserByIdRequest, opts ...client.CallOption) (*UserDetails, error)
	GetByEmail(ctx context.Context, in *UserByEmailRequest, opts ...client.CallOption) (*UserDetails, error)
	Update(ctx context.Context, in *UpdateUserRequest, opts ...client.CallOption) (*UserResponse, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) Create(ctx context.Context, in *NewUserRequest, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.name, "User.Create", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetByID(ctx context.Context, in *UserByIdRequest, opts ...client.CallOption) (*UserDetails, error) {
	req := c.c.NewRequest(c.name, "User.GetByID", in)
	out := new(UserDetails)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetByEmail(ctx context.Context, in *UserByEmailRequest, opts ...client.CallOption) (*UserDetails, error) {
	req := c.c.NewRequest(c.name, "User.GetByEmail", in)
	out := new(UserDetails)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Update(ctx context.Context, in *UpdateUserRequest, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.name, "User.Update", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserHandler interface {
	Create(context.Context, *NewUserRequest, *UserResponse) error
	GetByID(context.Context, *UserByIdRequest, *UserDetails) error
	GetByEmail(context.Context, *UserByEmailRequest, *UserDetails) error
	Update(context.Context, *UpdateUserRequest, *UserResponse) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) error {
	type user interface {
		Create(ctx context.Context, in *NewUserRequest, out *UserResponse) error
		GetByID(ctx context.Context, in *UserByIdRequest, out *UserDetails) error
		GetByEmail(ctx context.Context, in *UserByEmailRequest, out *UserDetails) error
		Update(ctx context.Context, in *UpdateUserRequest, out *UserResponse) error
	}
	type User struct {
		user
	}
	h := &userHandler{hdlr}
	return s.Handle(s.NewHandler(&User{h}, opts...))
}

type userHandler struct {
	UserHandler
}

func (h *userHandler) Create(ctx context.Context, in *NewUserRequest, out *UserResponse) error {
	return h.UserHandler.Create(ctx, in, out)
}

func (h *userHandler) GetByID(ctx context.Context, in *UserByIdRequest, out *UserDetails) error {
	return h.UserHandler.GetByID(ctx, in, out)
}

func (h *userHandler) GetByEmail(ctx context.Context, in *UserByEmailRequest, out *UserDetails) error {
	return h.UserHandler.GetByEmail(ctx, in, out)
}

func (h *userHandler) Update(ctx context.Context, in *UpdateUserRequest, out *UserResponse) error {
	return h.UserHandler.Update(ctx, in, out)
}
