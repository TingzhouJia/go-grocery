package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	orderservice "order-service/proto"
)

type OrderService struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *OrderService) Call(ctx context.Context, req *orderservice.Request, rsp *orderservice.Response) error {
	log.Info("Received OrderService.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *OrderService) Stream(ctx context.Context, req *orderservice.StreamingRequest, stream orderservice.OrderService_StreamStream) error {
	log.Infof("Received OrderService.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&orderservice.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *OrderService) PingPong(ctx context.Context, stream orderservice.OrderService_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&orderservice.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
