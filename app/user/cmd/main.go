package main

import (
	"fmt"

	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"

	"micro-todolist/app/user/repository/dao"
	"micro-todolist/app/user/service"
	"micro-todolist/conf"
	"micro-todolist/idl/pb"
)

func main() {
	conf.Init()
	dao.Init()
	// etcd注册
	etcdReg := registry.NewRegistry(
		registry.Addrs(fmt.Sprintf("%s:%d", conf.Conf.EtcdConfig.Host, conf.Conf.EtcdConfig.Port)))

	// new 一个微服务实例
	microService := micro.NewService(
		micro.Name("rpcUserService"),
		micro.Address(conf.Conf.ServerConfig.UserAddr),
		micro.Registry(etcdReg))

	microService.Init()
	_ = pb.RegisterUserServiceHandler(microService.Server(), service.NewUserSrv())
	_ = microService.Run()
}
