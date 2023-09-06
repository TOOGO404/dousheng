package comment_service

import (
	"api-gateway/rpc"
	"comment-service/dal/dao"
	comment "comment-service/kitex_gen/comment"
	"context"
	"datasource/database/model"
	"fmt"
	"time"
	usr "user-service/dal/dao"
	"user-service/kitex_gen/user"

	"github.com/cloudwego/hertz/pkg/protocol/consts"
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
		t := time.Unix(cm.CreatedAt, 0)
		tmp.CreateDate = fmt.Sprintf("%2d-%2d", t.Month(), t.Day())
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
	resp.CommentList = comments
	return resp, nil
}

const ADD_COMMENT = 1
const REMOVE_COMMNET = 2

func (s *CommentRpcServiceImpl) CommentAction(ctx context.Context, req *comment.CommentActionRequest) (*comment.CommentActionResponse, error) {

	resp := new(comment.CommentActionResponse)
	switch req.ActionType {
	case ADD_COMMENT:
		{
			_comment := new(model.Comment)
			_comment.Content = *req.CommentText
			_comment.VideoID = req.VideoId
			_comment.Who = req.Uid
			if req.CommentId != nil {
				_comment.SendTo = *req.CommentId
			}
			_comment.CreatedAt = time.Now().Unix()
			cmt_id, err := dao.AddComment(_comment)
			if err != nil {
				return nil, err
			}
			dbusr := usr.GetUserInfo(_comment.Who)
			resp.StatusMsg = &Msg
			resp.StatusCode = consts.StatusOK
			rpcResp, _ := rpc.UserRPCClient.GetUserInfo(ctx, &user.UserInfoReq{
				SendReqUserId: req.Uid,
				ReqUserId:     dbusr.Id,
			})
			resp.Comment = &comment.Comment{
				Id:      cmt_id,
				Content: _comment.Content,
				User: &comment.User{
					Id:   dbusr.Id,
					Name: dbusr.Name,
					//todo here
					FollowCount:     &rpcResp.UserInfo.FollowCount,
					FaviriteCount:   &rpcResp.UserInfo.FaviriteCount,
					FollowerCount:   rpcResp.UserInfo.FollowerCount,
					IsFollow:        rpcResp.UserInfo.IsFollow,
					Avatar:          &rpcResp.UserInfo.Avatar,
					Signature:       &rpcResp.UserInfo.Signature,
					BackgroundImage: &rpcResp.UserInfo.BackgroundImage,
					TotalFavorited:  &rpcResp.UserInfo.TotalFavorited,
					WorkCount:       &rpcResp.UserInfo.WorkCount,
				},
			}
			return resp, nil
		}
	case REMOVE_COMMNET:
		{
			willDelCommentId := req.GetUid()
			dao.DeleteComment(willDelCommentId)
			resp.StatusCode = 0
			resp.StatusMsg = &Msg
			return resp, nil
		}
	default:
		{
			resp.StatusMsg = &Wrongmsg
			resp.StatusCode = consts.StatusBadRequest
			resp.Comment = nil
			return resp, fmt.Errorf("bad action type")
		}
	}
}

func (s *CommentRpcServiceImpl) GetCommentCnt(ctx context.Context, vid int64) (int64, error) {
	return dao.GetCommentCnt(vid), nil
}
