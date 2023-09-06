package main

import (
	"log"
	"net"
	"sync"
	"time"
	user_service "user-service"
	user "user-service/kitex_gen/user/userrpcservice"
	"utils"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatalln(err)
	}
	listen, _ := net.ResolveTCPAddr("tcp", ":8081")
	svr := user.NewServer(&user_service.UserRpcServiceImpl{
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
			ServiceName: "user-service",
		}))

	err = svr.Run()

	if err != nil {
		log.Fatalln(err.Error())
	}
}
