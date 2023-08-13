// Code generated by Kitex v0.6.1. DO NOT EDIT.

package publishrpcservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	publish "publish-service/kitex_gen/publish"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	PublishAction(ctx context.Context, data *publish.VideoData, callOptions ...callopt.Option) (r *publish.PublishActionResp, err error)
	GetPublishLish(ctx context.Context, req *publish.PublishListReq, callOptions ...callopt.Option) (r *publish.PublishListResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kPublishRpcServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kPublishRpcServiceClient struct {
	*kClient
}

func (p *kPublishRpcServiceClient) PublishAction(ctx context.Context, data *publish.VideoData, callOptions ...callopt.Option) (r *publish.PublishActionResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PublishAction(ctx, data)
}

func (p *kPublishRpcServiceClient) GetPublishLish(ctx context.Context, req *publish.PublishListReq, callOptions ...callopt.Option) (r *publish.PublishListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetPublishLish(ctx, req)
}
