package main

import (
	"context"

	"go_micro_demo/proto"

	"github.com/iGoogle-ink/gotil/xlog"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

func main2() {
	etcdRegistry := etcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"127.0.0.1:2379"}
	})

	// 创建服务
	service := micro.NewService(
		micro.Name("client.hello.world"),
		micro.Registry(etcdRegistry),
		micro.Version("latest"),
	)

	// 初始化
	service.Init()

	// 向Service注册服务
	greeterCli := proto.NewGreeterService("service.hello.world", service.Client())

	// 调用Service方法
	rsp, err := greeterCli.Hello(context.Background(), &proto.Request{
		Name: "Jerry",
	})
	if err != nil {
		return
	}
	xlog.Debug("rsp:", rsp)
}