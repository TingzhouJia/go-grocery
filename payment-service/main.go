package main

import (
	"payment-service/handler"
	pb "payment-service/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("payment-service"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterPaymentServiceHandler(srv.Server(), new(handler.PaymentService))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
