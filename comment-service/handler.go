package comment_service

import (
	"comment-service/dal/dao"
	comment "comment-service/kitex_gen/comment"
	"context"
	"datasource/database/model"
	"log"
	"net/http"
	"time"
	usr "user-service/dal/dao"
)

type CommentRpcServiceImpl struct{}

var Msg string = "ok"
var Wrongmsg string = "oh, something is wrong"
var Tmp int64 = 0

func (s *CommentRpcServiceImpl) CommentGet(ctx context.Context, req *comment.CommentListRequest) (r *comment.CommentListResponse, err error) {
	commentList := dao.GetCommentListByVideoID(req.VideoId)
	var resp = new(comment.CommentListResponse)
	comments := make([]*comment.Comment, len(commentList))
	for index, cm := range commentList {
		tmp := new(comment.Comment)

		tmp.Id = cm.ID
		tmp.Content = cm.Content
		tmp.CreateDate = cm.CreateDate.String()
		dbusr := usr.GetUserInfo(cm.Who)
		_usr := new(comment.User)
		_usr.Id = dbusr.Id
		_usr.Name = dbusr.Name
		_usr.Avatar = &dbusr.Avatar
		_usr.BackgroundImage = &dbusr.Backgroud

		_usr.Signature = &dbusr.Signature

		tmp.User = _usr
		comments[index] = tmp
	}
	log.Println("len:", len(comments))
	resp.CommentList = comments
	return resp, nil
}

func (s *CommentRpcServiceImpl) CommentAction(ctx context.Context, req *comment.CommentActionRequest) (r *comment.CommentActionResponse, err error) {
	act := req.ActionType
	resp := new(comment.CommentActionResponse)
	log.Println(req)
	switch act {
	case 1:
		com := new(model.Comment)
		com.Content = *req.CommentText
		com.VideoID = req.VideoId
		com.Who = req.Uid
		com.User = usr.GetUserInfo(com.Who)
		com.CreateDate = time.Now()
		err := dao.AddComment(com)
		if err != nil {
			log.Println("Here")
			return nil, err
		}

		resp.StatusMsg = &Msg
		resp.StatusCode = http.StatusOK

		resp.Comment = &comment.Comment{
			Id:         com.ID,
			Content:    com.Content,
			CreateDate: com.CreateDate.String(),
			User: &comment.User{
				Id:              com.User.Id,
				Name:            com.User.Name,
				FollowCount:     &Tmp,
				FaviriteCount:   &Tmp,
				FollowerCount:   0,
				IsFollow:        true,
				Avatar:          &Msg,
				Signature:       &Msg,
				BackgroundImage: &Msg,
				TotalFavorited:  &Msg,
				WorkCount:       &Tmp,
			},
		}
		return r, nil

	case 2:
		com := new(model.Comment)
		com.Content = req.GetCommentText()
		com.VideoID = req.GetVideoId()
		userID := req.GetUid()
		com.Who = userID

		com.User = usr.GetUserInfo(userID)
		IDtodelete := *req.CommentId
		dao.DeleteComment(IDtodelete)
		resp.StatusMsg = &Msg
		resp.Comment = &comment.Comment{
			Id:         com.ID,
			Content:    com.Content,
			CreateDate: com.CreateDate.String(),
			User: &comment.User{
				Id:              com.User.Id,
				Name:            com.User.Name,
				FollowCount:     &Tmp,
				FaviriteCount:   &Tmp,
				FollowerCount:   0,
				IsFollow:        true,
				Avatar:          &Msg,
				Signature:       &Msg,
				BackgroundImage: &Msg,
				TotalFavorited:  &Msg,
				WorkCount:       &Tmp,
			},
		}
		resp.StatusCode = http.StatusOK
		return resp, nil
	}
	resp.StatusMsg = &Wrongmsg
	resp.StatusCode = http.StatusBadRequest
	resp.Comment = nil
	return resp, nil
}
