package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	paymentservice "payment-service/proto"
)

type PaymentService struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *PaymentService) Call(ctx context.Context, req *paymentservice.Request, rsp *paymentservice.Response) error {
	log.Info("Received PaymentService.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *PaymentService) Stream(ctx context.Context, req *paymentservice.StreamingRequest, stream paymentservice.PaymentService_StreamStream) error {
	log.Infof("Received PaymentService.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&paymentservice.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *PaymentService) PingPong(ctx context.Context, stream paymentservice.PaymentService_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&paymentservice.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
