// Code generated by Kitex v0.7.0. DO NOT EDIT.

package commentrpcservice

import (
	comment "comment-service/kitex_gen/comment"
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CommentGet(ctx context.Context, req *comment.CommentListRequest, callOptions ...callopt.Option) (r *comment.CommentListResponse, err error)
	CommentAction(ctx context.Context, req *comment.CommentActionRequest, callOptions ...callopt.Option) (r *comment.CommentActionResponse, err error)
	GetCommentCnt(ctx context.Context, vid int64, callOptions ...callopt.Option) (r int64, err error)
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
	return &kCommentRpcServiceClient{
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

type kCommentRpcServiceClient struct {
	*kClient
}

func (p *kCommentRpcServiceClient) CommentGet(ctx context.Context, req *comment.CommentListRequest, callOptions ...callopt.Option) (r *comment.CommentListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CommentGet(ctx, req)
}

func (p *kCommentRpcServiceClient) CommentAction(ctx context.Context, req *comment.CommentActionRequest, callOptions ...callopt.Option) (r *comment.CommentActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CommentAction(ctx, req)
}

func (p *kCommentRpcServiceClient) GetCommentCnt(ctx context.Context, vid int64, callOptions ...callopt.Option) (r int64, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetCommentCnt(ctx, vid)
}
