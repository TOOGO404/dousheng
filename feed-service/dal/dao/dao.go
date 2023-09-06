package dao

import (
	"datasource/database/model"
	"feed-service/dal"
)

func GetVIdeoListByLatestTime(latest_time int64) []*model.Video {
	var video_list []*model.Video

	dal.DB.Model(&model.Video{}).Where("created_at < ? and status = ?",
		latest_time, model.VIDEO_PUBLISHED).
		Order("id desc").
		Limit(15).
		Find(&video_list)
	return video_list
}
