package user_service

import (
	"api-gateway/rpc"
	"context"
	"datasource/database/model"
	"encoding/json"
	"fmt"
	"relationship-service/kitex_gen/relationship"
	"sync"
	"time"
	"user-service/dal"
	"user-service/dal/dao"
	user "user-service/kitex_gen/user"
	"utils"
)

// UserRpcServiceImpl implements the last service interface defined in the IDL.
type UserRpcServiceImpl struct {
	utils.Snowflake
}

// UserRegister implements the UserRpcServiceImpl interface.
func (s *UserRpcServiceImpl) UserRegister(_ context.Context, req *user.RegisterReq) (*user.RegisterResp, error) {
	_user := new(model.User)
	_user.Id = s.NextVal()
	_user.Email = req.Email
	pwd_MD5 := utils.GetPwd(req.Pwd)
	_user.Pwd = pwd_MD5
	_user.Name = "ds_" + fmt.Sprint(_user.Id)
	//todo storage
	_user.Avatar = "http://192.168.1.119:5000/avatar/default.png"
	_user.Backgroud = "http://192.168.1.119:5000/background/default.png"
	err := dao.RegisterNewUser(_user)
	if err != nil {
		return nil, err
	}
	resp := new(user.RegisterResp)
	resp.Uid = _user.Id
	return resp, nil
}

// UserLogin implements the UserRpcServiceImpl interface.
func (s *UserRpcServiceImpl) UserLogin(_ context.Context, req *user.LoginReq) (*user.LoginResp, error) {
	uid, err := dao.CheckUserPwd(&model.User{
		Email: req.Email,
		Pwd:   req.Pwd,
	})
	resp := new(user.LoginResp)
	resp.Uid = uid
	return resp, err
}

func getUserKey(uid int64) string {
	return fmt.Sprintf("user-%d", uid)
}

// GetUserInfo implements the UserRpcServiceImpl interface.
func (s *UserRpcServiceImpl) GetUserInfo(ctx context.Context, req *user.UserInfoReq) (*user.UserInfoResp, error) {
	//todo: insert userinfo to redis
	var info model.User
	var key = getUserKey(req.ReqUserId)
	data, err := dal.RdsClient.Get(ctx, key).Bytes()
	if err != nil {
		info = dao.GetUserInfo(req.ReqUserId)
		b, _ := json.Marshal(info)
		dal.RdsClient.Set(ctx, key, b, time.Minute*10)
	} else {
		json.Unmarshal(data, &info)
	}

	wg := sync.WaitGroup{}
	wg.Add(5)
	followerCnt, _ := rpc.RelationRPCClient.GetFollowerCnt(ctx, req.ReqUserId)
	followCnt, _ := rpc.RelationRPCClient.GetFollowCnt(ctx, req.ReqUserId)
	isFl, _ := rpc.RelationRPCClient.CheckSub(ctx, &relationship.CheckReq{
		Who:      req.SendReqUserId,
		ToUserId: req.ReqUserId,
	})
	fc, _ := rpc.FavoriteClient.GetFavorCount(ctx, req.ReqUserId)
	tf, _ := rpc.FavoriteClient.GetTotalFavorited(ctx, req.ReqUserId)
	return &user.UserInfoResp{
		UserInfo: &user.UserInfo{
			Id:              info.Id,
			Name:            info.Name,
			Avatar:          info.Avatar,
			BackgroundImage: info.Backgroud,
			Signature:       info.Signature,
			WorkCount:       info.WorkCount,
			FollowerCount:   followerCnt,
			IsFollow:        isFl,
			FollowCount:     followCnt,
			TotalFavorited:  tf,
			FaviriteCount:   fc,
		},
	}, nil
}
