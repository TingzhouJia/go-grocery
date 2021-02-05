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
	"grocery/user-service/model"

	"grocery/auth/handler"
	pb "grocery/auth/proto"
)
var (
	appName = "auth_srv"
	cfg     = &authCfg{}
)
type authCfg struct {
	common.AppCfg
}
func main() {
	initCfg()

	Reg:=etcd.NewRegistry(registryOptions)
	// Create service
	srv := micro.NewService(
		micro.Name(cfg.Name),
		micro.Registry(Reg),
		micro.Version(cfg.Version),
		micro.Address(cfg.Addr()),
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
