package dao

import (
	"datasource/database/model"
	"favorite-service/dal"
	"fmt"
	"log"
	"time"
)

func AddMsg(msg *model.Message) bool {
	msg.ComID = getCmdID(msg.FromUser, msg.ToUser)
	msg.CreatedAt = time.Now().UnixMilli() - 2000
	res := dal.DB.Create(msg)
	return res.Error == nil
}

func getCmdID(from, to int64) string {
	if from > to {
		return fmt.Sprintf("%v-%v", to, from)
	} else {
		return fmt.Sprintf("%v-%v", from, to)
	}

}

func GetMessageChat(from, to, ts int64) []*model.Message {
	var msgs []*model.Message
	dal.DB.Where(&model.Message{
		ComID: getCmdID(from, to),
	}).
		Where("created_at > ?", ts).
		Order("created_at ASC").
		Find(&msgs)
	log.Println(from, to, ts, len(msgs))
	return msgs
}
