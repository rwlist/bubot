package base

import (
	"github.com/petuhovskiy/telegram"
	"strconv"
)

type Response struct {
	ChatID    int
	UserID    int
	MessageID int
	Bot       *telegram.Bot
}

func (r *Response) Reply(text string) (*telegram.Message, error) {
	return r.Bot.SendMessage(&telegram.SendMessageRequest{
		ChatID:           strconv.Itoa(r.ChatID),
		Text:             text,
		ReplyToMessageID: r.MessageID,
	})
}
