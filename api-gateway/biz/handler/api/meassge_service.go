// Code generated by hertz generator.

package api

import (
	"context"
	"log"
	"msg-service/kitex_gen/message"

	api "api-gateway/biz/model/api"
	"api-gateway/rpc"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// MessageAction .
// @router /douyin/message/action/ [POST]
func MessageAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.MessageActionRequest
	err = c.BindAndValidate(&req)
	uid := c.GetInt64("uid")
	if err != nil || uid == 0 {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	rpc.MsgRPCClient.MessageAction(ctx, &message.MessageActionReq{
		UserId:     uid,
		ToUserId:   req.ToUserID,
		Msg:        &req.Content,
		ActionType: req.ActionType,
	})
	resp := new(api.MessageActionResponse)
	resp.StatusCode = 0
	log.Println(uid, " send msg")
	c.JSON(consts.StatusOK, resp)
}

// GetMessageChat .
// @router /douyin/message/chat/ [GET]
func GetMessageChat(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.MessageChatRequest
	err = c.BindAndValidate(&req)
	uid := c.GetInt64("uid")
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	rpcResp, err := rpc.MsgRPCClient.GetMessageChat(ctx, &message.MessageChatReq{
		UserId:     uid,
		ToUserId:   req.ToUserID,
		PreMsgTime: req.PreMsgTime,
	})
	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}
	resp := new(api.MessageChatResponse)
	resp.StatusCode = 0
	resp.MessageList = make([]*api.Message, len(rpcResp.MessageList))
	for idx, msg := range rpcResp.MessageList {
		resp.MessageList[idx] = &api.Message{
			ID:         msg.Id,
			FromUserID: msg.FromUserId,
			ToUserID:   msg.ToUserId,
			CreateTime: msg.CreateTime,
			Content:    msg.Msg,
		}
	}
	log.Println("get msg")
	c.JSON(consts.StatusOK, resp)
}
