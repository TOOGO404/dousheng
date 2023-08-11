// Code generated by hertz generator.

package api

import (
	api "api-gateway/biz/model/api"
	"api-gateway/mw"
	"api-gateway/rpc"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"log"
	"user-service/kitex_gen/user"
)

var signedTokenErr api.RegisterResponse
var existsTheSameEmailErr api.RegisterResponse

func init() {
	var sErrMsg = "获取token失败"
	signedTokenErr = api.RegisterResponse{
		StatusCode: -2,
		StatusMsg:  &sErrMsg,
		UserID:     0,
		Token:      "",
	}
	var eErrMsg = "存在相同的Email"
	existsTheSameEmailErr = api.RegisterResponse{
		StatusCode: -1,
		StatusMsg:  &eErrMsg,
		UserID:     0,
		Token:      "",
	}
}

// Register .
// @router /douyin/user/register/ [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.RegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		c.Abort()
		return
	}

	rpcResp, err := rpc.UserRPCClient.UserRegister(ctx, &user.RegisterReq{
		Email: req.Email,
		Pwd:   req.Password,
	})
	if err != nil {
		c.JSON(consts.StatusOK, &existsTheSameEmailErr)
	}
	tokenStr, err := mw.SignedToken(rpcResp.Uid)

	if err != nil {
		c.JSON(consts.StatusOK, &signedTokenErr)
	} else {
		resp := new(api.RegisterResponse)
		resp.Token = tokenStr
		resp.StatusCode = rpcResp.StatusCode
		resp.StatusMsg = &rpcResp.StatusMsg
		resp.UserID = rpcResp.Uid
		c.JSON(consts.StatusOK, resp)
	}

}

// Login .
// @router /douyin/user/login/ [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.LoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		c.Abort()
		return
	}
	rpcResp, err := rpc.UserRPCClient.UserLogin(ctx, &user.LoginReq{
		Email: req.Username,
		Pwd:   req.Password,
	})
	if err != nil {

	}
	resp := new(api.LoginResponse)
	tokenStr, err := mw.SignedToken(rpcResp.Uid)
	resp.Token = tokenStr
	resp.StatusCode = 0
	resp.UserID = rpcResp.Uid
	resp.StatusMsg = nil
	c.JSON(consts.StatusOK, resp)
}

// GetUserInfo .
// @router /douyin/user/ [GET]
func GetUserInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UserInfoRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		c.Abort()
		return
	} else {
		uid := c.GetInt64("uid")
		rpcResp, _ := rpc.UserRPCClient.GetUserInfo(ctx, &user.UserInfoReq{
			SendReqUserId: uid,
			ReqUserId:     req.UserID,
		})
		log.Println(uid)
		resp := new(api.UserInfoResponse)
		resp.StatusCode = 0

		resp.StatusMsg = nil
		resp.User = &api.User{
			ID:              req.UserID,
			Name:            rpcResp.UserInfo.Name,
			FollowerCount:   rpcResp.UserInfo.FollowerCount,
			IsFollow:        true,
			Avatar:          rpcResp.UserInfo.Avatar,
			BackgroundImage: rpcResp.UserInfo.BackgroundImage,
			Signature:       rpcResp.UserInfo.Signature}
		c.JSON(consts.StatusOK, resp)
	}

}