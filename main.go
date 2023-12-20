package main

import (
	"fmt"

	"github.com/lenny-mo/emall-utils/tracer"
	"github.com/lenny-mo/router/global"
	"github.com/lenny-mo/router/router"
	"github.com/opentracing/opentracing-go"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/registry/consul/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
)

func main() {
	// New Service
	// 注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	rpcServiceName := "go.micro.service.router.globalRPC"

	// 链路追踪
	err := tracer.InitTracer(rpcServiceName, "127.0.0.1:6831")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tracer.Closer.Close()
	opentracing.SetGlobalTracer(tracer.Tracer)

	// 创建rpc服务
	global.GlobalRPCService = micro.NewService(
		micro.Name(rpcServiceName),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8098"),
		micro.Registry(consulRegistry),
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
	)
	global.GlobalRPCService.Init()

	// 启动global rpc service
	go func() {
		if err := global.GlobalRPCService.Run(); err != nil {
			fmt.Println(err)
		}
	}()

	// 添加router web service
	// 所有的路由规则都在这里添加
	ginRouter := router.InitRouters()
	webService := web.NewService(
		web.Name("go.micro.service.router"),
		web.Address("127.0.0.1:8099"),
		web.Handler(ginRouter),
		web.Registry(consulRegistry),
	)

	webService.Init() // 启动web server 服务器

	tracer.InitTracer("router", "127.0.0.1:6831")

	if err := webService.Run(); err != nil {
		fmt.Println(err)
	}
}
