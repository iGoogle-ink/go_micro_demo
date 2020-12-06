package main

import (
	"context"

	"go_micro_demo/api/payment"

	"github.com/iGoogle-ink/gotil/xlog"
	"github.com/micro/go-micro/v2"
)

func main() {
	//etcdRegistry := etcd.NewRegistry(func(options *registry.Options) {
	//	options.Addrs = []string{"127.0.0.1:2379"}
	//})

	// 创建服务
	service := micro.NewService(
		micro.Name("client.hello.world"),
		//micro.Registry(etcdRegistry),
		micro.Version("latest"),
	)

	// 初始化
	service.Init()

	// 向Service注册服务
	greeterCli := payment.NewGreeterService("service.hello.world", service.Client())

	// 调用Service方法
	rsp, err := greeterCli.Hello(context.Background(), &payment.Request{
		Name: "Jerry",
	})
	if err != nil {
		return
	}
	xlog.Debug("rsp:", rsp)
}
