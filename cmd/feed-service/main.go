package main

import (
	feed_service "feed-service"
	feed "feed-service/kitex_gen/feed/feedrpcservice"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func main() {

	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatalln(err)
	}
	listen, _ := net.ResolveTCPAddr("tcp", ":8082")
	svr := feed.NewServer(new(feed_service.FeedRpcServiceImpl),
		server.WithServiceAddr(listen),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "feed-service",
		}))
	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
