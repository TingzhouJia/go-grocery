package main

import (
	"fmt"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"grocery/basic/config"
	"grocery/user-service/model"

	"grocery/auth/handler"
	pb "grocery/auth/proto"
)

func main() {


	Reg:=etcd.NewRegistry(registryOptions)
	// Create service
	srv := micro.NewService(
		micro.Name("auth"),
		micro.Registry(Reg),
		micro.Version("latest"),
	)
	srv.Init(
		micro.Action(func(context *cli.Context) error {
			model.Init()
			return nil

		}),
	)
	// Register handler
	pb.RegisterServiceHandler(srv.Server(), new(handler.Auth))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}