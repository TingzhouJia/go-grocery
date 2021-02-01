package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	userservice "user-service/proto"
)

type UserService struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *UserService) Call(ctx context.Context, req *userservice.Request, rsp *userservice.Response) error {
	log.Info("Received UserService.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *UserService) Stream(ctx context.Context, req *userservice.StreamingRequest, stream userservice.UserService_StreamStream) error {
	log.Infof("Received UserService.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&userservice.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *UserService) PingPong(ctx context.Context, stream userservice.UserService_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&userservice.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
