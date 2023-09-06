package main

import (
	favorite_service "favorite-service"
	service "favorite-service/kitex_gen/favorite/favoriterpcservice"
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {

	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatalln(err)
	}
	listen, _ := net.ResolveTCPAddr("tcp", ":8085")
	svr := service.NewServer(new(favorite_service.FavoriteRpcServiceImpl),
		server.WithServiceAddr(listen),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "favorite-service",
		}))
	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
