package publish_service

import (
	"context"
	model "datasource/database/model"
	"fmt"
	"os"
	"publish-service/dal"
	"publish-service/dal/dao"
	publish "publish-service/kitex_gen/publish"
	"utils"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

// PublishRpcServiceImpl implements the last service interface defined in the IDL.
type PublishRpcServiceImpl struct {
	utils.Snowflake
}

const CoverPath = "/home/navy/Desktop/temp/cover"
const VideoPath = "/home/navy/Desktop/temp/video"
const TempPath = "/home/navy/Desktop/temp/temp"
const StoagePath = "http://192.168.1.119:5000"

// PublishAction implements the PublishRpcServiceImpl interface.
func (s *PublishRpcServiceImpl) PublishAction(ctx context.Context, data *publish.VideoData) (*publish.PublishActionResp, error) {
	videoId := s.NextVal()
	//todo:will move to transcoding-service
	fileName := fmt.Sprintf("%s/%d.%s", VideoPath, videoId, data.FileType)
	coverName := fmt.Sprintf("%s/%d", CoverPath, videoId)
	fmt.Printf("%s/%d", CoverPath, videoId)
	err := os.WriteFile(fileName, data.Data, os.ModePerm)
	if err != nil {
		return &publish.PublishActionResp{
			VideoId: -1,
		}, err
	}

	err = ffmpeg.Input(fileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output(coverName, ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		Run()
	if err != nil {
		return &publish.PublishActionResp{
			VideoId: -1,
		}, err
	}
	outputName := fmt.Sprintf("%s/ds_%d.%s", VideoPath, videoId, "m3u8")
	err = ffmpeg.Input(fileName).
		Output(outputName, ffmpeg.KwArgs{
			"c:v":           "libx264",
			"c:a":           "aac",
			"f":             "hls",
			"hls_list_size": 0,
			"hls_time":      5,
		}).
		Run()

	if err != nil {
		return &publish.PublishActionResp{
			VideoId: -1,
		}, err
	} else {
		video := model.Video{
			Id:       videoId,
			Author:   data.Uid,
			Title:    data.Title,
			Status:   model.VIDEO_PUBLISHED,
			PlayUrl:  fmt.Sprintf("%s/video/ds_%d.%s", StoagePath, videoId, "m3u8"),
			CoverUrl: fmt.Sprintf("%s/cover/%d", StoagePath, videoId),
		}
		dal.DB.Create(&video)
		var user model.User
		dal.DB.Where("uid = ?", data.Uid).First(&user)
		user.WorkCount += 1
		dal.DB.Updates(&user)
		return &publish.PublishActionResp{
			VideoId: videoId,
		}, nil
	}
}

// GetPublishLish implements the PublishRpcServiceImpl interface.
func (s *PublishRpcServiceImpl) GetPublishLish(ctx context.Context, req *publish.PublishListReq) (*publish.PublishListResp, error) {
	vs := dao.GetVideoList(req.UserId)
	fmt.Println(len(vs))
	v := make([]*publish.Video, len(vs))
	for idx, video := range vs {
		v[idx] = &publish.Video{
			Id:       video.Id,
			AuthorID: video.Author,
			PlayUrl:  video.PlayUrl,
			CoverUrl: video.CoverUrl,
			Title:    video.Title,
		}
	}
	resp := publish.PublishListResp{
		VideoList: v,
	}
	return &resp, nil
}
