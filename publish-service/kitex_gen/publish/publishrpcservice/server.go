// Code generated by Kitex v0.6.1. DO NOT EDIT.
package publishrpcservice

import (
	server "github.com/cloudwego/kitex/server"
	publish "publish-service/kitex_gen/publish"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler publish.PublishRpcService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
