package base

import "strings"

type Command struct {
	Text   string
	Action string
	Args   []string
}

func NewCommand(text string) *Command {
	args := strings.Split(text, " ")

	return &Command{
		Text:   text,
		Action: args[0],
		Args:   args,
	}
}
