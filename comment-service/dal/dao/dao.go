package dao

import (
	"comment-service/dal"
	"datasource/database/model"
	"log"
)

func GetCommentListByVideoID(VideoId int64) []*model.Comment {
	var commentList []*model.Comment
	dal.DB.Model(&model.Comment{}).Where("video_id = ?", VideoId).
		Order("create_date desc").
		Find(&commentList)
	log.Printf("commentList len :%d\n", len(commentList))
	return commentList
}
func AddComment(comment *model.Comment) error {
	return dal.DB.Create(comment).Error
}

func DeleteComment(ID1 int64) {
	dal.DB.Delete(&model.Comment{
		ID: ID1,
	})
}
