package driver

import (
	"strings"
)

type Room struct {
	Description string
}

func (r *Room) KnowsCommand(command string) bool {
	return strings.HasPrefix(command, "north")
}
