package driver

import (
	"bufio"
	"io"
)

type PlayerConnection struct {
	In io.Reader
	Out io.Writer
	UserName string
	Room *Room
}

func (p *PlayerConnection) Inform(msg string) {
	io.WriteString(p.Out, msg + "\n")
}

func (p *PlayerConnection) Play(){
	p.Look()
	buffIn := bufio.NewReader(p.In)
	for {
		line, _, err := buffIn.ReadLine()
		command := string(line)
		if err != nil {
			p.Inform("Omg! your connection died!?")
			return
		} else if command == "exit" {
			p.Inform("Thanks for playing!")
			return
		} else if command == "help" {
			p.Inform("try 'exit'")
		} else if command == "look" {
			p.Look()
		}else if p.Room.IsDirection(command) {
			p.ChangeRoom(command)
		} else {
			p.Inform("Unknown command")
		}
	}

}

func (p *PlayerConnection) Look() {
	p.Inform(p.Room.Look())
}

func (p *PlayerConnection) ChangeRoom(direction string) {
	room := p.Room.Direction(direction)
	if room == nil {
		p.Inform("There is no exit that way")
	} else {
		p.Room = room
		p.Look()
	}
}
