// Code generated by hertz generator.

package api

import (
	"api-gateway/rpc"
	"context"
	"favorite-service/kitex_gen/favorite"
	"net/http"

	api "api-gateway/biz/model/api"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// FavoriteAction .
// @router /douyin/favorite/dao/ [POST]
func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.FavoriteActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.FavoriteActionResponse)

	Uid := c.GetInt64("uid")
	if Uid == 0 {
		resp.StatusMsg = &Wrongmsg
		resp.StatusCode = http.StatusUnauthorized
		c.JSON(http.StatusUnauthorized, resp)
		c.Abort()
		return
	}
	_, err = rpc.FavoriteRPCClient.FavoriteAction(ctx, &favorite.FavoriteActionRequest{
		Uid:        Uid,
		VideoId:    req.VideoID,
		ActionType: req.ActionType,
	})
	if err != nil {
		resp.StatusCode = http.StatusBadRequest
		resp.StatusMsg = &Wrongmsg
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	resp.StatusMsg = &Okmsg
	resp.StatusCode = 200
	c.JSON(consts.StatusOK, resp)
}

// GetFavoriteList .
// @router /douyin/favorite/list/ [GET]
func GetFavoriteList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.FavoriteListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.FavoriteListResponse)
	Uid := c.GetInt64("uid")
	if Uid == 0 {
		resp.StatusMsg = &Wrongmsg
		resp.StatusCode = http.StatusUnauthorized
		c.JSON(http.StatusUnauthorized, resp)
		c.Abort()
		return
	}
	rpcresp, err := rpc.FavoriteRPCClient.GetFavoriteList(ctx, &favorite.FavoriteListRequest{
		Uid:    Uid,
		UserId: req.UserID,
	})
	if err != nil {
		resp.StatusCode = http.StatusBadRequest
		resp.StatusMsg = &Wrongmsg
		resp.VideoList = nil
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	resp.StatusMsg = &Okmsg
	resp.StatusCode = 200
	videos := make([]*api.Video, len(rpcresp.VideoList))
	for index, video := range rpcresp.VideoList {
		av := new(api.Video)

		av.ID = video.Id
		av.Title = video.Title
		av.Author = &api.User{
			ID: video.Author.Id,
		}
		av.PlayURL = video.PlayUrl
		av.CoverURL = video.CoverUrl
		av.IsFavorite = true
		av.CommentCount = 0
		av.FavoriteCount = 0
		videos[index] = av
	}
	resp.VideoList = videos
	c.JSON(consts.StatusOK, resp)
}
