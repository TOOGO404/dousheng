package publish_service

import (
	publish "publish-service/kitex_gen/publish";
	model "datasource/database/model";
	"utils";
	"context";
	"time";
)

// PublishRpcServiceImpl implements the last service interface defined in the IDL.
type PublishRpcServiceImpl struct {
	utils.Snowflake
}

// PublishAction implements the PublishRpcServiceImpl interface.
func (s *PublishRpcServiceImpl) PublishAction(ctx context.Context, data *publish.VideoData) (resp *publish.PublishActionResp, err error) {
	// TODO: Your code here...
	videoId := new(utils.Snowflake).NextVal()
	video := model.Video{
		Id		 : videoId,
		// Author   :
		Title	 : data.GetTitle(),
		// PlayUrl  : 
		// CoverUrl :
		Status   : 1,
		Updated  : time.Now().Unix(),
		Created  : time.Now().Unix(),
	}
	DB.Table("videos").Debug().Create(&video)
	return &publish.PublishActionResp{
		VideoId : videoId,
	}, nil
}

// GetPublishLish implements the PublishRpcServiceImpl interface.
func (s *PublishRpcServiceImpl) GetPublishLish(ctx context.Context, req *publish.PublishListReq) (resp *publish.PublishListResp, err error) {
	// TODO: Your code here...
	return
}
