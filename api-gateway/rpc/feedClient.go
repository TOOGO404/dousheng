package rpc

import (
	"feed-service/kitex_gen/feed/feedrpcservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/connpool"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"time"
)

var FeedRPCClient feedrpcservice.Client

func init() {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	client := feedrpcservice.MustNewClient("feed-service",
		client.WithResolver(r),
		client.WithLongConnection(connpool.IdleConfig{
			MaxIdlePerAddress: 10,
			MaxIdleGlobal:     100,
			MaxIdleTimeout:    time.Minute,
			MinIdlePerAddress: 2,
		}))
	FeedRPCClient = client
}
