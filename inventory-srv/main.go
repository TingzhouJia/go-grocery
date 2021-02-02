package main

import (
	"inventory-srv/handler"
	pb "inventory-srv/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("inventory-srv"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterInventorySrvHandler(srv.Server(), new(handler.InventorySrv))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
