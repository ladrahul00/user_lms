package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/wolf00/user_lms/handler"

	organization "github.com/wolf00/user_lms/proto/organization"
	source "github.com/wolf00/user_lms/proto/source"
	user "github.com/wolf00/user_lms/proto/user"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("0.1"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	user.RegisterUserHandler(service.Server(), new(handler.UserServiceHandler))
	organization.RegisterOrganizationHandler(service.Server(), new(handler.OrganizationServiceHandler))
	source.RegisterSourceHandler(service.Server(), new(handler.SourceServiceHandler))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
