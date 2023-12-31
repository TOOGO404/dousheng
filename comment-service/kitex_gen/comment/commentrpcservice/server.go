// Code generated by Kitex v0.7.0. DO NOT EDIT.
package commentrpcservice

import (
	comment "comment-service/kitex_gen/comment"
	server "github.com/cloudwego/kitex/server"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler comment.CommentRpcService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
