package driver

import (
	"strings"
)
type Exit struct {
	Direction string
	Destination *Room
}
type Room struct {
	Description string
	exits []Exit
}

func (r *Room) IsDirection(command string) bool {
	if r.exits == nil {
		r.exits = make([]Exit, 0)
	}
	for _,e := range r.exits {
		if strings.HasPrefix(command, e.Direction) {
			return true
		}
	}
	for _,e := range [4]string{"north", "south", "east", "west"}{
		if command == e {
			return true
		}
	}
	return false
}
func (r *Room) AddExit(direction string, destination *Room) {
	if r.exits == nil {
		r.exits = make([]Exit, 0)
	}
	exit := Exit { Direction: direction, Destination: destination }
	r.exits = append( r.exits, exit)
}
func (r *Room) Direction(direction string) *Room{
	for _, exit := range r.exits {
		if exit.Direction == direction {
			return exit.Destination
		}
	}
	return nil
}
func (r *Room) Look() string {
	return r.Description
}
