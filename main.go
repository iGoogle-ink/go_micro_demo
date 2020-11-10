package main

import (
	"context"
	"log"

	"go_micro_demo/api"

	"github.com/iGoogle-ink/gotil/xlog"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *api.Request, rsp *api.Response) error {
	rsp.Msg = "Hello " + req.Name
	return nil
}

func main() {

	etcdRegistry := etcd.NewRegistry(func(opt *registry.Options) {
		opt.Addrs = []string{"127.0.0.1:2379"}
	})

	// 新建服务
	service := micro.NewService(
		micro.Name("service.hello.world"),
		micro.Registry(etcdRegistry),
		micro.Version("latest"),
	)

	// 初始化
	service.Init(
		micro.BeforeStart(func() error {
			xlog.Debug("启动前的日志打印")
			return nil
		}),
		micro.AfterStart(func() error {
			xlog.Debug("启动后的日志打印")
			return nil
		}),
	)

	// 注册服务
	_ = api.RegisterGreeterHandler(service.Server(), new(Greeter))

	// 启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
