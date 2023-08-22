package dao

import (
	"datasource/database/model"
	"favorite-service/dal"
	"log"
)

func GetFavoriteVideosbyUser(UserId int64) []*model.Video {
	var VideoList []*model.Video
	var VideoID []int64
	dal.DB.Model(&model.Like{}).Select("video_id").Where("who= ?", UserId).Find(&VideoID)
	log.Printf("VideoList len :%d\n", len(VideoList))
	return GetVideobyID(VideoID)
}

func GetVideobyID(ID []int64) []*model.Video {
	var vd []*model.Video
	dal.DB.Model(&model.Video{}).Where("id IN ?", ID).Find(&vd)
	return vd
}
