package dao

import (
	"datasource/database/model"
	"feed-service/dal"
	"log"
)

func GetVIdeoListByLatestTime(latest_time int64) []*model.Video {
	var video_list []*model.Video
	dal.DB.Model(&model.Video{}).Where("created < ?", latest_time).
		Order("id desc").
		Limit(10).
		Find(&video_list)
	log.Printf("video len :%d\n", len(video_list))
	return video_list
}
