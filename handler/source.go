package handler

import (
	"context"

	source "github.com/wolf00/user_lms/proto/source"
	"github.com/wolf00/user_lms/user/services"
)

// SourceServiceHandler blah
type SourceServiceHandler struct {
	services.SourceServce
}

// GetBySourceTag is a single request handler called via client.Call or the generated client code
func (e *SourceServiceHandler) GetBySourceTag(ctx context.Context, req *source.SourceRequest, rsp *source.SourceResponse) error {
	return e.SourceServce.GetBySourceTag(ctx, req, rsp)
}
