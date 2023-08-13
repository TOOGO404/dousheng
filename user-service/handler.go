package user_service

import (
	"context"
	"datasource/database/model"
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
	// TODO: Your code here...

	_user := new(model.User)
	_user.Id = s.NextVal()
	_user.Email = req.Email
	_user.Pwd = req.Pwd
	_user.Name = "user" + string(_user.Id)

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
	// TODO: Your code here...
	uid := dao.CheckUserPwd(&model.User{
		Email: req.Email,
		Pwd:   req.Pwd,
	})
	resp := new(user.LoginResp)
	resp.Uid = uid
	return resp, nil

}

// GetUserInfo implements the UserRpcServiceImpl interface.
func (s *UserRpcServiceImpl) GetUserInfo(ctx context.Context, req *user.UserInfoReq) (*user.UserInfoResp, error) {
	// TODO: Your code here...
	info := dao.GetUserInfo(req.ReqUserId)
	return &user.UserInfoResp{
		UserInfo: &user.UserInfo{
			Id:              info.Id,
			Name:            info.Name,
			Avatar:          &info.Avatar,
			FollowerCount:   0,
			IsFollow:        true,
			BackgroundImage: &info.Backgroud,
			Signature:       &info.Signature,
		},
	}, nil
}
