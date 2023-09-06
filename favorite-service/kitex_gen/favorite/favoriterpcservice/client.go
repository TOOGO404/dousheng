// Code generated by Kitex v0.7.0. DO NOT EDIT.

package favoriterpcservice

import (
	"context"
	favorite "favorite-service/kitex_gen/favorite"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	FavoriteAction(ctx context.Context, req *favorite.FavoriteActionReq, callOptions ...callopt.Option) (r *favorite.FavoriteActionResp, err error)
	GetFavoriteList(ctx context.Context, req *favorite.FavoriteListReq, callOptions ...callopt.Option) (r *favorite.FavoriteListResp, err error)
	GetTotalFavorited(ctx context.Context, uid int64, callOptions ...callopt.Option) (r int64, err error)
	IsFavorited(ctx context.Context, req *favorite.CheckFavoritedReq, callOptions ...callopt.Option) (r bool, err error)
	GetFavorCount(ctx context.Context, uid int64, callOptions ...callopt.Option) (r int64, err error)
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
	return &kFavoriteRpcServiceClient{
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

type kFavoriteRpcServiceClient struct {
	*kClient
}

func (p *kFavoriteRpcServiceClient) FavoriteAction(ctx context.Context, req *favorite.FavoriteActionReq, callOptions ...callopt.Option) (r *favorite.FavoriteActionResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FavoriteAction(ctx, req)
}

func (p *kFavoriteRpcServiceClient) GetFavoriteList(ctx context.Context, req *favorite.FavoriteListReq, callOptions ...callopt.Option) (r *favorite.FavoriteListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFavoriteList(ctx, req)
}

func (p *kFavoriteRpcServiceClient) GetTotalFavorited(ctx context.Context, uid int64, callOptions ...callopt.Option) (r int64, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetTotalFavorited(ctx, uid)
}

func (p *kFavoriteRpcServiceClient) IsFavorited(ctx context.Context, req *favorite.CheckFavoritedReq, callOptions ...callopt.Option) (r bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.IsFavorited(ctx, req)
}

func (p *kFavoriteRpcServiceClient) GetFavorCount(ctx context.Context, uid int64, callOptions ...callopt.Option) (r int64, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFavorCount(ctx, uid)
}