package consumer

import (
	"log"

	"github.com/nsqio/go-nsq"
)

type TranscodingHandler struct{}

func (h *TranscodingHandler) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		return nil
	}

	log.Println(m.Body)

	return nil
}
