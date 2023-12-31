// Code generated by hertz generator.

package api

import (
	api "api-gateway/biz/model/api"
	"api-gateway/rpc"
	"comment-service/kitex_gen/comment"
	"context"
	"log"
	"net/http"
	"user-service/kitex_gen/user"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

var Okmsg string = "ok"
var Wrongmsg string = "wrong"

// CommentAction .
// @router /douyin/comment/action/ [POST]
func CommentAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.CommentActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.CommentActionResponse)
	Uid := c.GetInt64("uid")
	if Uid == 0 {
		resp.StatusMsg = &Wrongmsg
		resp.StatusCode = http.StatusBadRequest
		c.JSON(http.StatusBadRequest, resp)
		c.Abort()
		return
	}
	log.Println(Uid, req)
	act := req.GetActionType()
	switch act {
	case 1:
		_, err := rpc.CommentRPCClient.CommentAction(ctx, &comment.CommentActionRequest{
			Uid:         Uid,
			VideoId:     req.VideoID,
			ActionType:  req.ActionType,
			CommentText: req.CommentText,
		})
		if err != nil {
			log.Println("Here")
		}
		resp.StatusCode = 0
		resp.StatusMsg = nil
		//resp.Comment = &api.Comment{
		//	ID:         rpcResp.Comment.Id,
		//	Content:    rpcResp.Comment.Content,
		//	CreateDate: rpcResp.Comment.CreateDate,
		//	User: &api.User{
		//		ID:              rpcResp.Comment.User.Id,
		//		Name:            rpcResp.Comment.User.Name,
		//		FollowerCount:   rpcResp.Comment.User.FollowerCount,
		//		IsFollow:        true,
		//		Avatar:          rpcResp.Comment.User.Avatar,
		//		BackgroundImage: rpcResp.Comment.User.BackgroundImage,
		//		Signature:       rpcResp.Comment.User.Signature,
		//	},
		//}

		resp.Comment = nil
	case 2:
		_, err := rpc.CommentRPCClient.CommentAction(ctx, &comment.CommentActionRequest{
			Uid:        Uid,
			VideoId:    req.VideoID,
			ActionType: req.ActionType,
			CommentId:  req.CommentID,
		})
		if err != nil {
			c.JSON(consts.StatusInternalServerError, utils.H{})
		}

		resp.StatusCode = 0
		resp.StatusMsg = nil
		resp.Comment = nil
		//resp.Comment = &api.Comment{
		//	ID:         rpcResp.Comment.Id,
		//	Content:    rpcResp.Comment.Content,
		//	CreateDate: rpcResp.Comment.CreateDate,
		//	User: &api.User{
		//		ID:              rpcResp.Comment.User.Id,
		//		Name:            rpcResp.Comment.User.Name,
		//		FollowerCount:   rpcResp.Comment.User.FollowerCount,
		//		IsFollow:        true,
		//		Avatar:          rpcResp.Comment.User.Avatar,
		//		BackgroundImage: rpcResp.Comment.User.BackgroundImage,
		//		Signature:       rpcResp.Comment.User.Signature,
		//	},
		//}

	}

	c.JSON(consts.StatusOK, resp)
}

// GetComment .
// @router /douyin/comment/list/ [GET]
func GetComment(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.CommentListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		c.Abort()
		return
	}

	uid := c.GetInt64("uid")
	rpcResp, _ := rpc.CommentRPCClient.CommentGet(ctx, &comment.CommentListRequest{
		Uid:     uid,
		VideoId: req.VideoID,
	})
	resp := new(api.CommentListResponse)
	resp.StatusCode = 0
	resp.StatusMsg = &Okmsg
	Comments := make([]*api.Comment, len(rpcResp.CommentList))
	for index, _comment := range rpcResp.CommentList {
		com := new(api.Comment)

		com.CreateDate = _comment.CreateDate
		com.Content = _comment.Content
		com.ID = _comment.Id
		usr := _comment.User
		id := usr.Id
		usrrpcResp, _ := rpc.UserRPCClient.GetUserInfo(ctx, &user.UserInfoReq{
			SendReqUserId: uid,
			ReqUserId:     id,
		})
		com.User = &api.User{
			ID:              usrrpcResp.UserInfo.Id,
			Name:            usrrpcResp.UserInfo.Name,
			FollowerCount:   usrrpcResp.UserInfo.FollowerCount,
			IsFollow:        true,
			Avatar:          &usrrpcResp.UserInfo.Avatar,
			BackgroundImage: &usrrpcResp.UserInfo.BackgroundImage,
			Signature:       &usrrpcResp.UserInfo.Signature}
		Comments[index] = com
	}
	resp.CommentList = Comments
	c.JSON(consts.StatusOK, resp)
}
