package main

import (
	"fmt"
	"github.com/micro/cli/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
	"grocery/basic"
	"grocery/basic/config"
	 "grocery/order-web/handler"
	"net/http"
)

func main() {
	basic.Init()
	micReg := etcd.NewRegistry(registryOptions)

	service := web.NewService(
		web.Name("mu.micro.book.web.orders"),
		web.Version("latest"),
		web.Registry(micReg),
		web.Address(":8091"),
	)
	if err := service.Init(
		web.Action(
			func(c *cli.Context) {
				// 初始化handler
				handler.Init()
			}),
	); err != nil {
		log.Fatal(err)
	}
	authHandler := http.HandlerFunc(handler.New)
	service.Handle("/orders/new", handler.AuthWrapper(authHandler))
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}


func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}
