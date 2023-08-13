package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/connpool"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"publish-service/kitex_gen/publish/publishrpcservice"
	"time"
)

var PublishRPCClient publishrpcservice.Client

func init() {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	client := publishrpcservice.MustNewClient("publish-service",
		client.WithResolver(r),
		client.WithLongConnection(connpool.IdleConfig{
			MaxIdlePerAddress: 10,
			MaxIdleGlobal:     100,
			MaxIdleTimeout:    time.Minute,
			MinIdlePerAddress: 2,
		}))
	PublishRPCClient = client
}
