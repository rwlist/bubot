package base

import (
	"github.com/petuhovskiy/telegram"
	"strings"
)

type Request struct {
	Update  *telegram.Update
	Message *telegram.Message
	Command *Command

	Response *Response
}

func NewRequest(bot *telegram.Bot, upd telegram.Update) *Request {
	resp := &Response{
		Bot: bot,
	}

	req := &Request{
		Update:   &upd,
		Response: resp,
	}

	if msg := upd.Message; msg != nil {
		req.Message = msg
		resp.MessageID = msg.MessageID

		if msg.Chat != nil {
			resp.ChatID = msg.Chat.ID
		}

		if msg.From != nil {
			resp.UserID = msg.From.ID
		}

		if strings.HasPrefix(msg.Text, "/") {
			req.Command = NewCommand(msg.Text[1:])
		}
	}

	return req
}
