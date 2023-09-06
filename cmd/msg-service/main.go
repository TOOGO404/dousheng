package main

import (
	"log"
	service "msg-service"
	message "msg-service/kitex_gen/message/meassgerpcservice"
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
	listen, _ := net.ResolveTCPAddr("tcp", ":8087")

	svr := message.NewServer(new(service.MeassgeRpcServiceImpl),
		server.WithServiceAddr(listen),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "msg-service",
		}))
	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
