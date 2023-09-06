package dao

import (
	"context"
	"datasource/database/model"
	"favorite-service/dal"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm/clause"
)

func AddLike(uid, vid int64) {
	dal.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"is_like"}),
	}).
		Create(&model.Like{
			Who:     uid,
			VideoID: vid,
			Id:      genLikeID(uid, vid),
			IsLike:  true,
		})
	ret, err := dal.RdsClient.
		Exists(context.Background(), genFavKey(vid)).Result()
	if err == nil && ret > 0 {
		dal.RdsClient.Incr(context.Background(), genFavKey(vid))
	}
}

func genLikeID(uid, vid int64) string {
	return fmt.Sprintf("%v-%v", uid, vid)
}

func SetDislike(uid, vid int64) {
	dal.DB.Model(&model.Like{}).Where(&model.Like{
		Who:     uid,
		VideoID: vid,
	}).Update("is_like", 0)

	ret, err := dal.RdsClient.
		Exists(context.Background(), genFavKey(vid)).Result()
	if err == nil && ret > 0 {
		dal.RdsClient.Decr(context.Background(), genFavKey(vid))
	}
}

func IsFavorited(uid, vid int64) bool {
	var cnt int64 = 0
	dal.DB.Model(&model.Like{}).
		Where(&model.Like{IsLike: true, Id: genLikeID(uid, vid)}).
		Count(&cnt)
	return cnt != 0
}

func GetFavoriteList(uid int64) []*model.Like {
	var likes []*model.Like
	dal.DB.Where(&model.Like{
		IsLike: true,
		Who:    uid,
	}).Find(&likes)
	return likes
}

func GetTotalFavorited(uid int64) int64 {
	var cnt int64 = 0
	dal.DB.
		Where(&model.Like{
			Who:    uid,
			IsLike: true,
		}).
		Count(&cnt)
	return cnt
}

func genFavKey(uid int64) string {
	return fmt.Sprintf("fav-%v", uid)
}

func GetFavoriteCnt(vid int64) int64 {
	key := genFavKey(vid)
	sc := dal.RdsClient.Get(context.Background(), key)
	if sc.Err() != nil {
		var cnt int64 = 0
		dal.DB.Model(&model.Like{}).
			Where(&model.Like{
				IsLike:  true,
				VideoID: vid}).
			Count(&cnt)
		dal.RdsClient.Set(context.Background(), key, cnt, time.Minute*5)
		log.Println("video like", cnt)
		return cnt
	} else {
		i, _ := sc.Int64()
		return i
	}
}

func GetVideo(vid int64) (*model.Video, *model.User) {
	var video *model.Video = new(model.Video)
	var author *model.User = new(model.User)
	dal.DB.Where("id = ?", vid).First(video)
	dal.DB.Where("id = ?", video.Author).First(author)
	return video, author
}
