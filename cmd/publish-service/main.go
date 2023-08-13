package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
	"publish-service"
	publish "publish-service/kitex_gen/publish/publishrpcservice"
	"sync"
	"time"
	"utils"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatalln(err)
	}
	listen, _ := net.ResolveTCPAddr("tcp", ":8083")
	svr := publish.NewServer(&publish_service.PublishRpcServiceImpl{
		Snowflake: utils.Snowflake{
			Mutex:        sync.Mutex{},
			Timestamp:    time.Now().Unix(),
			Workerid:     0,
			Datacenterid: 0,
			Sequence:     0,
		},
	},
		server.WithServiceAddr(listen),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "publish-service",
		}))

	err = svr.Run()

	if err != nil {
		log.Fatalln(err.Error())
	}
}
