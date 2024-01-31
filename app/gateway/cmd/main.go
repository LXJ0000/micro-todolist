package main

import (
	"fmt"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/web"
	"micro-todolist/app/gateway/router"
	"micro-todolist/app/gateway/rpc"
	"micro-todolist/conf"
	"time"
)

func main() {
	conf.Init()
	rpc.Init()
	// etcd注册
	etcdReg := registry.NewRegistry(
		registry.Addrs(fmt.Sprintf("%s:%d", conf.Conf.EtcdConfig.Host, conf.Conf.EtcdConfig.Port)))

	// new 一个微服务实例
	webService := web.NewService(
		web.Name("httpService"),
		web.Address(conf.Conf.WebAddr),
		web.Registry(etcdReg),
		web.Handler(router.NewRouter()),
		web.RegisterTTL(time.Second*30),
		web.Metadata(map[string]string{"protocol": "http"}),
	)

	_ = webService.Init()
	_ = webService.Run()
}
