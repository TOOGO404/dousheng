package main

import (
	"log"
	"net"
	relationship_impl "relationship-service"
	relationship "relationship-service/kitex_gen/relationship/relationshiprpcservice"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatalln(err)
	}
	listen, _ := net.ResolveTCPAddr("tcp", ":8086")
	svr := relationship.NewServer(
		&relationship_impl.RelationshipRpcServiceImpl{},
		server.WithServiceAddr(listen),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "relationship-service",
		}))

	err = svr.Run()

	if err != nil {
		log.Fatalln(err.Error())
	}
}
