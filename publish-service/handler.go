package publish_service

import (
	"context"
	publish "publish-service/kitex_gen/publish"
	"utils"
)

// PublishRpcServiceImpl implements the last service interface defined in the IDL.
type PublishRpcServiceImpl struct {
	utils.Snowflake
}

// PublishAction implements the PublishRpcServiceImpl interface.
func (s *PublishRpcServiceImpl) PublishAction(ctx context.Context, data *publish.VideoData) (resp *publish.PublishActionResp, err error) {
	// TODO: Your code here...
	return
}

// GetPublishLish implements the PublishRpcServiceImpl interface.
func (s *PublishRpcServiceImpl) GetPublishLish(ctx context.Context, req *publish.PublishListReq) (resp *publish.PublishListResp, err error) {
	// TODO: Your code here...
	return
}
