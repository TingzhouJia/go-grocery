package payment

import (
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"grocery/basic/common"
	inventory "grocery/inventory-srv/proto"
	order "grocery/order-service/proto"
	"sync"
)

var (
	s            *service
	invClient    inventory.InventorySrvService
	ordSClient   order.OrdersService
	m            sync.RWMutex
	payPublisher micro.Publisher
)
type service struct {
}



// Service 服务类
type Service interface {
	// PayOrder 支付订单
	PayOrder(orderId int64) (err error)
}
func GetService() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("[GetService] GetService 未初始化")
	}
	return s, nil
}

// Init 初始化库存服务层
func Init() {
	m.Lock()
	defer m.Unlock()

	if s != nil {
		return
	}

	invClient = inventory.NewInventorySrvService("service.inventory", client.DefaultClient)
	ordSClient = order.NewOrdersService("service.orders", client.DefaultClient)
	payPublisher = micro.NewPublisher(common.TopicPaymentDone, client.DefaultClient)
	s = &service{}
}