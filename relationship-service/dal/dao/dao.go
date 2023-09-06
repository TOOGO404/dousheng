package dao

import (
	"datasource/database/model"
	"fmt"
	"relationship-service/dal"
	"relationship-service/kitex_gen/relationship"

	"gorm.io/gorm/clause"
)

func AddSubRecord(who int64, to_user_id int64) (bool, error) {
	res := dal.DB.
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"is_sub"}),
		}).
		Create(&model.Sub{
			Id:       fmt.Sprintf("%v-%v", who, to_user_id),
			Who:      who,
			ToUserID: to_user_id,
			IsSub:    true,
		}).Error
	return res == nil, res
}

func RemoveSubRecord(who int64, to_user_id int64) (bool, error) {
	res := dal.DB.Model(&model.Sub{}).Where("id = ?",
		fmt.Sprintf("%d-%d", who, to_user_id)).
		Update("is_sub", false).Error
	return res == nil, res
}

func IsSub(who int64, to_user_id int64) bool {
	var sub model.Sub
	var err = dal.DB.Where(&model.Sub{
		Id: fmt.Sprintf("%v-%v", who, to_user_id),
	}).
		First(&sub).Error
	if err != nil {
		return false
	} else {
		return sub.IsSub
	}

}

func GetSubList(who int64) []*relationship.UserInfo {

	var users []*model.Sub
	dal.DB.Where("who = ?", who).
		Find(&users)
	var userinfos []*relationship.UserInfo = make([]*relationship.UserInfo, len(users))
	for idx, user := range users {
		var temp_user model.User
		dal.DB.Where("id = ?", user.ToUserID).
			First(&temp_user)
		userinfos[idx] = &relationship.UserInfo{
			Id:       temp_user.Id,
			Name:     temp_user.Name,
			IsFollow: true,
			Avatar:   temp_user.Avatar,
		}
	}
	return userinfos
}

func GetFollowerList(toUid int64) []*relationship.UserInfo {
	var subs []*model.Sub
	dal.DB.Where(&model.Sub{
		ToUserID: toUid,
		IsSub:    true,
	}).
		Find(&subs)
	var rs = make([]*relationship.UserInfo, len(subs))
	for idx, sub := range subs {
		var temp_user model.User
		dal.DB.Where("id = ?", sub.Who).
			First(&temp_user)
		rs[idx] = &relationship.UserInfo{
			Id:       temp_user.Id,
			Name:     temp_user.Name,
			IsFollow: true,
			Avatar:   temp_user.Avatar,
		}
	}
	return rs
}

func GetFollowCnt(uid int64) int64 {
	var cnt int64
	dal.DB.Model(&model.Sub{}).
		Where(&model.Sub{
			Who:   uid,
			IsSub: true,
		}).
		Count(&cnt)
	return cnt
}

func GetFollowerCnt(uid int64) int64 {
	var cnt int64
	dal.DB.Model(&model.Sub{}).
		Where(&model.Sub{
			ToUserID: uid,
			IsSub:    true}).
		Count(&cnt)
	return cnt
}

func GetFriendList(who int64) []*relationship.UserInfo {
	return GetFollowerList(who)
}
