package handler

import (
	"context"

	user "github.com/wolf00/golang_lms/user/proto/user"
	"github.com/wolf00/golang_lms/user/services"
)

// UserServiceHandler blah
type UserServiceHandler struct {
	services.UserService
}

// Create is a single request handler called via client.Call or the generated client code
func (e *UserServiceHandler) Create(ctx context.Context, req *user.NewUserRequest, rsp *user.Response) error {
	return e.UserService.Create(ctx, req, rsp)
}

// GetByID is a single request handler called via client.Call or the generated client code
func (e *UserServiceHandler) GetByID(ctx context.Context, req *user.UserByIdRequest, rsp *user.UserDetails) error {
	return e.UserService.GetByID(ctx, req, rsp)
}

// GetByEmail is a single request handler called via client.Call or the generated client code
func (e *UserServiceHandler) GetByEmail(ctx context.Context, req *user.UserByEmailRequest, rsp *user.UserDetails) error {
	return e.UserService.GetByEmail(ctx, req, rsp)
}

// Update is a single request handler called via client.Call or the generated client code
func (e *UserServiceHandler) Update(ctx context.Context, req *user.UpdateUserRequest, rsp *user.Response) error {
	return e.UserService.Update(ctx, req, rsp)
}
