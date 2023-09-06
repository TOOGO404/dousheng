package main

import (
	comment_service "comment-service"
	comment "comment-service/kitex_gen/comment/commentrpcservice"
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
	listen, _ := net.ResolveTCPAddr("tcp", ":8084")
	svr := comment.NewServer(new(comment_service.CommentRpcServiceImpl),
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
