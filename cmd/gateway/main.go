package main

import (
	"api-gateway"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/logger/accesslog"
)

func main() {
	h := server.New(server.WithHostPorts(":8899"))
	//添加 accesslog打印访问日志功能
	h.Use(accesslog.New(
		accesslog.WithFormat(
			"[${time}] ${status} - ${latency} ${method} ${path} ${queryParams}")))
	api_gateway.Register(h)
	h.Spin()
}
