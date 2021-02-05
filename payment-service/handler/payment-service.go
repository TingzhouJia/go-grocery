package handler

import (
	"context"
	"grocery/payment-service/model/payment"
	proto "grocery/payment-service/proto"
)

var (
	paymentService payment.Service
)

type Service struct {
}

// Init 初始化handler
func Init() {
	paymentService, _ = payment.GetService()
}

// New 新增订单
func (e *Service) PayOrder(ctx context.Context, req *proto.Request, rsp *proto.Response) (err error) {

	err = paymentService.PayOrder(req.OrderId)
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