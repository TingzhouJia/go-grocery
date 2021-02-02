package main

import (
	"fmt"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"grocery/basic"
	"grocery/basic/config"
	"grocery/inventory-srv/handler"
	"grocery/inventory-srv/model"
	pb "grocery/inventory-srv/proto"
)

func main() {

	reg:=etcd.NewRegistry(registryOptions)

	// Create service
	srv := micro.NewService(
		micro.Name("inventory-srv"),
		micro.Version("latest"),
		micro.Registry(reg),
	)

	basic.Init()

	srv.Init(
		micro.Action(func(context *cli.Context) error {
			model.Init()
			handler.Init()
			return nil
		}),
	)
	// Register handler
	pb.RegisterInventorySrvHandler(srv.Server(), new(handler.Service))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}