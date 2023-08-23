package favorite_service

import (
	"context"
	"datasource/database/model"
	"favorite-service/dal/dao"
	favorite "favorite-service/kitex_gen/favorite"
	"log"
	"net/http"
	"strconv"
	usr "user-service/dal/dao"
)

var Msg string = "ok"
var Wrongmsg string = "oh, something is wrong"
var Tmp int64 = 0
var OK string = "ok"

type FavoriteRpcServiceImpl struct{}

func (f FavoriteRpcServiceImpl) FavoriteAction(ctx context.Context, req *favorite.FavoriteActionRequest) (r *favorite.FavoriteActionResponse, err error) {
	act := req.ActionType
	resp := new(favorite.FavoriteActionResponse)
	log.Println(req)
	switch act {
	case 1: //like
		lk := new(model.Like)
		lk.Who = req.Uid
		lk.VideoID = req.VideoId
		_usr := usr.GetUserInfo(req.Uid)
		lk.User = _usr
		lk.Video = dao.GetAVideo(req.VideoId)
		err := dao.AddLike(lk)
		if err != nil {
			log.Println("Here")
			return nil, err
		}

		resp.StatusMsg = &Msg
		resp.StatusCode = http.StatusOK

		return resp, nil
	case 2:
		err := dao.DeleteFavorite(req.Uid, req.VideoId)
		if err != nil {
			log.Println("Here")
			return nil, err
		}
		resp.StatusMsg = &Msg
		resp.StatusCode = http.StatusOK
	}
	return resp, nil
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
		tmp.IsFavorite = IsFavorited(ID, vd.Id)
		tmp.FavoriteCount = GetFavorCount(vd.Id)
		tmp.CommentCount = 0
		Total := strconv.FormatInt(GetTotalFavorited(vd.User.Id), 10)
		tmp.Author = &favorite.User{
			Id:              vd.User.Id,
			Name:            vd.User.Name,
			FollowCount:     &Tmp,
			FaviriteCount:   &tmp.FavoriteCount,
			FollowerCount:   0,
			IsFollow:        true,
			Avatar:          &Msg,
			Signature:       &Msg,
			BackgroundImage: &Msg,
			TotalFavorited:  &Total,
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

func GetTotalFavorited(Userid int64) int64 {
	return dao.GetTotalFavoritedByusr(Userid)
}
func IsFavorited(Userid, Videoid int64) bool {
	return dao.IsFavored(Userid, Videoid)
}
func GetFavorCount(Userid int64) int64 {
	return dao.GetCount(Userid)
}
