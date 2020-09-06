package handler

import (
	"context"

	organization "github.com/wolf00/user_lms/proto/organization"
	"github.com/wolf00/user_lms/services"
)

// OrganizationServiceHandler blah
type OrganizationServiceHandler struct {
	services.OrganizationService
}

// Create is a single request handler called via client.Call or the generated client code
func (e *OrganizationServiceHandler) Create(ctx context.Context, req *organization.NewOrganizationRequest, rsp *organization.OrganizationResponse) error {
	return e.OrganizationService.Create(ctx, req, rsp)
}

// GetByID is a single request handler called via client.Call or the generated client code
func (e *OrganizationServiceHandler) GetByID(ctx context.Context, req *organization.OrganizationByIdRequest, rsp *organization.OrganizationDetails) error {
	return e.OrganizationService.GetByID(ctx, req, rsp)
}

// Update is a single request handler called via client.Call or the generated client code
func (e *OrganizationServiceHandler) Update(ctx context.Context, req *organization.UpdateOrganizationRequest, rsp *organization.OrganizationResponse) error {
	return e.OrganizationService.Update(ctx, req, rsp)
}
