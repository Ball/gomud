package driver

import (
	"fmt"
	"strings"
)
type Exit struct {
	Direction string
	Destination *Room
}
type Room struct {
	Description string
	exits []Exit
	players []*PlayerConnection
}

func (r *Room) IsSay(command string) bool {
	if r.exits == nil {
		r.exits = make([]Exit, 0)
	}
	if strings.HasPrefix(command, "say") {
		return true
	}
	return false
}
func (r *Room) Say(command string, player *PlayerConnection) {
	for _, p := range r.players {
		if p == player {
			continue
		}
		message := strings.TrimSpace(strings.Replace(command, "say", "", 1))
		p.Inform(fmt.Sprintf("%s says %s", player.UserName, message))
	}
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
func (r *Room) Enter(player *PlayerConnection){
	if r.players == nil {
		r.players = make([]*PlayerConnection, 0)
	}
	for _, p := range r.players {
		p.Inform(fmt.Sprintf("%s has entered the room\n", player.UserName))
	}
	r.players = append(r.players, player)
	player.Enter(r)
}
func (r *Room) Look() string {
	return r.Description
}
