package main

import (
	favorite_service "favorite-service"
	favorite "favorite-service/kitex_gen/favorite/favoriterpcservice"
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
	listen, _ := net.ResolveTCPAddr("tcp", ":8084")
	svr := favorite.NewServer(new(favorite_service.FavoriteRpcServiceImpl),
		server.WithServiceAddr(listen),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "comment-service",
		}))
	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
