package dao

import (
	"comment-service/dal"
	"datasource/database/model"
	"log"
)

func GetCommentListByVideoID(VideoId int64) []*model.Comment {
	var commentList []*model.Comment
	dal.DB.Model(&model.Comment{}).Where("video_id = ?", VideoId).
		Order("created_at desc").
		Find(&commentList)
	return commentList
}
func AddComment(comment *model.Comment) (int64, error) {
	err := dal.DB.Create(comment).Error
	var cmt model.Comment
	dal.DB.Where(comment).First(&cmt)
	return cmt.ID, err
}

func DeleteComment(ID int64) {
	dal.DB.Where("id = ?", ID).Delete(&model.Comment{})
}

func GetCommentCnt(vid int64) int64 {
	var cnt int64
	dal.DB.Model(&model.Comment{}).
		Where("video_id = ?", vid).Count(&cnt)
	log.Println(vid, cnt)
	return cnt
}
