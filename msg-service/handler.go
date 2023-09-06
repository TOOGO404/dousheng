package msg_service

import (
	"context"
	"datasource/database/model"
	"errors"
	"msg-service/dal/dao"
	message "msg-service/kitex_gen/message"
	"utils"
)

// MeassgeRpcServiceImpl implements the last service interface defined in the IDL.
type MeassgeRpcServiceImpl struct {
	utils.Snowflake
}

const SEND_MSG = 1

// MessageAction implements the MeassgeRpcServiceImpl interface.
func (s *MeassgeRpcServiceImpl) MessageAction(ctx context.Context, req *message.MessageActionReq) (resp bool, err error) {
	switch req.ActionType {
	case SEND_MSG:
		{
			msg := model.Message{
				Id:       s.NextVal(),
				FromUser: req.UserId,
				ToUser:   req.ToUserId,
				Msg:      *req.Msg,
			}
			return dao.AddMsg(&msg), nil
		}
	default:
		{
			return false, errors.New("no such action")
		}
	}
}

// GetMessageChat implements the MeassgeRpcServiceImpl interface.
func (s *MeassgeRpcServiceImpl) GetMessageChat(ctx context.Context, req *message.MessageChatReq) (resp *message.MessageChatResp, err error) {
	resp = new(message.MessageChatResp)
	msgs := dao.GetMessageChat(req.UserId, req.ToUserId, req.PreMsgTime)
	resp.MessageList = make([]*message.Message, len(msgs))
	for idx, msg := range msgs {
		resp.MessageList[idx] = &message.Message{
			Id:         msg.Id,
			FromUserId: msg.FromUser,
			ToUserId:   msg.ToUser,
			Msg:        msg.Msg,
			CreateTime: msg.CreatedAt,
		}
	}
	return resp, nil
}
