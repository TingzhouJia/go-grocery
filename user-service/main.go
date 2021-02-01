package main

import (
	"user-service/handler"
	pb "user-service/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("user-service"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterUserServiceHandler(srv.Server(), new(handler.UserService))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
