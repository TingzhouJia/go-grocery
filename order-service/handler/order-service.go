package handler

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"
	"grocery/order-service/model/orders"
	proto "grocery/order-service/proto"
)

var (
	ordersService orders.Service
)

type Orders struct {
}

// Init 初始化handler
func Init() {
	ordersService, _ = orders.GetService()
}

// New 新增订单
func (e *Orders) New(ctx context.Context, req *proto.Request, rsp *proto.Response) (err error) {
	orderId, err := ordersService.New(req.BookId, req.UserId)
	if err != nil {
		rsp.Success = false
		rsp.Error = &proto.Error{
			Detail: err.Error(),
		}
		return
	}

	rsp.Order = &proto.Order{
		Id: orderId,
	}
	return
}

// GetOrder 获取订单
func (e *Orders) GetOrder(ctx context.Context, req *proto.Request, rsp *proto.Response) (err error) {
	log.Logf(log.ErrorLevel,"[GetOrder] 收到获取订单请求，%d", req.OrderId)

	rsp.Order, err = ordersService.GetOrder(req.OrderId)
	if err != nil {
		rsp.Success = false
		rsp.Error = &proto.Error{
			Detail: err.Error(),
		}
		return
	}

	rsp.Success = true
	return
}
