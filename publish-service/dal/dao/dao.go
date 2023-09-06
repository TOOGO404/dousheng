package dao

import (
	"datasource/database/model"
	"publish-service/dal"
)

func GetVideoList(id int64) []*model.Video {
	var videos []*model.Video
	dal.DB.Where(&model.Video{
		Author: id,
	}).Find(&videos)
	return videos
}
