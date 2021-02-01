package main

import (
	"fmt"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"grocery/basic"
	"grocery/basic/config"
	"grocery/user-service/handler"
	"grocery/user-service/model"
	pb "grocery/user-service/proto"
)

func main() {
	basic.Init()

	microReg:=etcd.NewRegistry(registryOptions)

	// Create service
	srv := micro.NewService(
		micro.Name("grocery.user-service"),
		micro.Registry(microReg),
		micro.Version("latest"),
	)

	srv.Init(
		micro.Action(func(context *cli.Context) error {
			model.Init()
			handler.Init()
			return nil
		}),
	)
	// Register handler
	pb.RegisterUserHandler(srv.Server(), new(handler.UserService))

	// Run service
	if err := srv.Run(); err != nil {

	}
}

func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}