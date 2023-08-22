package favorite_service

import (
	"context"
	"favorite-service/dal/dao"
	favorite "favorite-service/kitex_gen/favorite"
)

var Msg string = "ok"
var Wrongmsg string = "oh, something is wrong"
var Tmp int64 = 0
var OK string = "ok"

type FavoriteRpcServiceImpl struct{}

func (f FavoriteRpcServiceImpl) FavoriteAction(ctx context.Context, req *favorite.FavoriteActionRequest) (r *favorite.FavoriteActionResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (f FavoriteRpcServiceImpl) GetFavoriteList(ctx context.Context, req *favorite.FavoriteListRequest) (r *favorite.FavoriteListResponse, err error) {
	ID := req.Uid
	_Videos := dao.GetFavoriteVideosbyUser(ID)
	resp := new(favorite.FavoriteListResponse)
	resp.StatusMsg = &OK
	resp.StatusCode = 0

	Videos := make([]*favorite.Video, len(_Videos))
	for index, vd := range _Videos {
		tmp := new(favorite.Video)
		tmp.Id = vd.Id
		tmp.IsFavorite = true
		tmp.FavoriteCount = 0
		tmp.CommentCount = 0
		tmp.Author = &favorite.User{
			Id:              vd.User.Id,
			Name:            vd.User.Name,
			FollowCount:     &Tmp,
			FaviriteCount:   &Tmp,
			FollowerCount:   0,
			IsFollow:        true,
			Avatar:          &Msg,
			Signature:       &Msg,
			BackgroundImage: &Msg,
			TotalFavorited:  &Msg,
			WorkCount:       &Tmp,
		}
		tmp.CoverUrl = vd.CoverUrl
		tmp.PlayUrl = vd.PlayUrl
		tmp.Title = vd.Title
		Videos[index] = tmp
	}
	resp.VideoList = Videos
	return resp, nil
}
