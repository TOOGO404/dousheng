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

func DeleteFavorite(usrID, vdID int64) error {
	return dal.DB.Delete(&model.Like{
		Who:     usrID,
		VideoID: vdID,
	}).Error
}
func GetVideobyID(ID []int64) []*model.Video {
	var vd []*model.Video
	dal.DB.Model(&model.Video{}).Where("id IN ?", ID).Find(&vd)
	return vd
}

func AddLike(like *model.Like) error {
	return dal.DB.Create(like).Error
}

func GetAVideo(ID int64) model.Video {
	var v model.Video
	dal.DB.Model(&model.Video{}).Where("id = ?", ID).Find(&v)
	return v
}

func IsFavored(Userid, Videoid int64) bool {
	ok := dal.DB.Where("who = ? AND video_id = ?", Userid, Videoid).RowsAffected
	if ok == 1 {
		return true
	} else {
		return false
	}
}
func GetUserVideoIDs(Userid int64) []int64 {
	var ans []int64
	dal.DB.Model(&model.Like{}).Select("video_id").Where("who = ?", Userid).Find(&ans)
	return ans
}
func GetTotalFavoritedByusr(Userid int64) int64 {
	var ans int64
	arr := GetUserVideoIDs(Userid)
	dal.DB.Model(&model.Like{}).Where("video_id IN ?", arr).Count(&ans)
	return ans
}
func GetCount(Userid int64) int64 {
	var ans int64
	dal.DB.Model(&model.Like{}).Where("who = ?", Userid).Count(&ans)
	return ans
}
