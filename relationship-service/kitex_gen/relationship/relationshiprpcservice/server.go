// Code generated by Kitex v0.7.0. DO NOT EDIT.
package relationshiprpcservice

import (
	server "github.com/cloudwego/kitex/server"
	relationship "relationship-service/kitex_gen/relationship"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler relationship.RelationshipRpcService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
