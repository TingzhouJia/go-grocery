package main

import (
	"fmt"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-plugins/config/source/grpc/v2"
	"grocery/basic"
	"grocery/basic/common"
	"grocery/basic/config"
	"grocery/order-service/handler"
	"grocery/order-service/model/orders"
	pb "grocery/order-service/proto"
	"grocery/order-service/subscriber"
	"time"
)
var (
	appName = "orders_srv"
	cfg     = &appCfg{}
)

type appCfg struct {
	common.AppCfg
}
func main() {
	initCfg()
	micReg := etcd.NewRegistry(registryOptions)
	// Create service
	srv := micro.NewService(
		micro.Name("order-service"),
		micro.Version("latest"),
		micro.Registry(micReg),
		micro.Address(cfg.Addr()),
		//进行健康检查
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*20),
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
	etcdCfg := &common.Etcd{}
	err := config.C().App("etcd", etcdCfg)
	if err != nil {
		panic(err)
	}

	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.Host, etcdCfg.Port)}
}

func initCfg() {
	source := grpc.NewSource(
		grpc.WithAddress("127.0.0.1:9600"),
		grpc.WithPath("micro"),
	)

	basic.Init(config.WithSource(source))

	err := config.C().App(appName, cfg)
	if err != nil {
		panic(err)
	}


	return
}
