package main

import (
	"fmt"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"grocery/basic/common"
	"grocery/basic/config"
	"grocery/order-service/handler"
	"grocery/order-service/model/orders"
	pb "grocery/order-service/proto"
	"grocery/order-service/subscriber"
)

func main() {
	micReg := etcd.NewRegistry(registryOptions)
	// Create service
	srv := micro.NewService(
		micro.Name("order-service"),
		micro.Version("latest"),
		micro.Registry(micReg),
	)
	srv.Init(
		micro.Action(func(context *cli.Context) error {
			orders.Init()
			subscriber.Init()
			handler.Init()
			return nil
		}),
	)

	// Register handler
	err:=pb.RegisterOrdersHandler(srv.Server(), new(handler.Orders))

	if err != nil {
		logger.Fatal(err)
	}
	err = micro.RegisterSubscriber(common.TopicPaymentDone, srv.Server(), subscriber.PayOrder)
	if err != nil {
		logger.Fatal(err)
	}

	// 注册服务

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}