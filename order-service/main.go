package main

import (
	"order-service/handler"
	pb "order-service/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("order-service"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterOrderServiceHandler(srv.Server(), new(handler.OrderService))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
