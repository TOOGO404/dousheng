package favoriteservice

import (
	"context"
	"favorite-service/dal/dao"
	favorite "favorite-service/kitex_gen/favorite"
)

// FavoriteRpcServiceImpl implements the last service interface defined in the IDL.
type FavoriteRpcServiceImpl struct{}

const LIKE_ACT = 1
const DISLIKE_ACT = 2

// FavoriteAction implements the FavoriteRpcServiceImpl interface.
func (s *FavoriteRpcServiceImpl) FavoriteAction(ctx context.Context, req *favorite.FavoriteActionReq) (*favorite.FavoriteActionResp, error) {
	resp := new(favorite.FavoriteActionResp)
	resp.StatusCode = 0
	switch req.ActionType {
	case LIKE_ACT:
		{
			dao.AddLike(req.Uid, req.VideoId)
			return resp, nil
		}
	case DISLIKE_ACT:
		{
			dao.SetDislike(req.Uid, req.VideoId)
			return resp, nil
		}
	default:
		{
			return nil, nil
		}
	}
}

// GetFavoriteList implements the FavoriteRpcServiceImpl interface.
func (s *FavoriteRpcServiceImpl) GetFavoriteList(ctx context.Context, req *favorite.FavoriteListReq) (*favorite.FavoriteListResp, error) {
	resp := new(favorite.FavoriteListResp)
	likes := dao.GetFavoriteList(req.UserId)
	resp.VideoList = make([]*favorite.Video, len(likes))
	for idx, like := range likes {
		cnt, _ := s.GetFavorCount(ctx, like.VideoID)
		v, u := dao.GetVideo(like.VideoID)
		resp.VideoList[idx] = &favorite.Video{
			Id: like.VideoID,
			Author: &favorite.User{
				Id:   u.Id,
				Name: u.Name,
			},
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: cnt,
		}
	}
	return resp, nil
}

// GetTotalFavorited implements the FavoriteRpcServiceImpl interface.
func (s *FavoriteRpcServiceImpl) GetTotalFavorited(ctx context.Context, uid int64) (resp int64, err error) {
	return dao.GetTotalFavorited(uid), nil
}

// IsFavorited implements the FavoriteRpcServiceImpl interface.
func (s *FavoriteRpcServiceImpl) IsFavorited(ctx context.Context, req *favorite.CheckFavoritedReq) (resp bool, err error) {
	return dao.IsFavorited(req.Uid, req.Vid), nil
}

// GetFavorCount implements the FavoriteRpcServiceImpl interface.
func (s *FavoriteRpcServiceImpl) GetFavorCount(ctx context.Context, vid int64) (resp int64, err error) {
	// TODO: Your code here...
	return dao.GetFavoriteCnt(vid), nil
}
