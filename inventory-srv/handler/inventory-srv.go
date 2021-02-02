package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	inventorysrv "inventory-srv/proto"
)

type InventorySrv struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *InventorySrv) Call(ctx context.Context, req *inventorysrv.Request, rsp *inventorysrv.Response) error {
	log.Info("Received InventorySrv.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *InventorySrv) Stream(ctx context.Context, req *inventorysrv.StreamingRequest, stream inventorysrv.InventorySrv_StreamStream) error {
	log.Infof("Received InventorySrv.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&inventorysrv.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *InventorySrv) PingPong(ctx context.Context, stream inventorysrv.InventorySrv_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&inventorysrv.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
