// Code generated by Kitex v0.6.1. DO NOT EDIT.

package feedrpcservice

import (
	comment "comment-service/kitex_gen/comment"
	"context"
	feed "feed-service/kitex_gen/feed"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetFeed(ctx context.Context, req *feed.FeedReq, callOptions ...callopt.Option) (r *feed.FeedResp, err error)
}

func (p *kCommentRpcServiceClient) CommentGet(ctx context.Context, req *comment.CommentListRequest, callOptions ...callopt.Option) (r *comment.CommentListResponse, err error) {

	return p.CommentGet(ctx,req)
}

func (p *kCommentRpcServiceClient) CommentAction(ctx context.Context, req *comment.CommentActionRequest, callOptions ...callopt.Option) (r *comment.CommentActionResponse, err error) {

	return p.CommentAction(ctx,req)
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
	return &kFeedRpcServiceClient{
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

type kFeedRpcServiceClient struct {
	*kClient
}
type kCommentRpcServiceClient struct {
	*comment.CommentRpcServiceClient
}

func (p *kFeedRpcServiceClient) GetFeed(ctx context.Context, req *feed.FeedReq, callOptions ...callopt.Option) (r *feed.FeedResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFeed(ctx, req)
}
