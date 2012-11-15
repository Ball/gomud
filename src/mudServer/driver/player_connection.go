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
	io.WriteString(p.Out, msg)
}

func (p *PlayerConnection) Play(){
	p.Inform(p.Room.Description + "\n")
	buffIn := bufio.NewReader(p.In)
	for {
		line, _, err := buffIn.ReadLine()
		command := string(line)
		if err != nil {
			p.Inform("Omg! your connection died!?\n")
			return
		} else if command == "exit" {
			p.Inform("Thanks for playing!\n")
			return
		} else if command == "help" {
			p.Inform("try 'exit'\n")
		} else if p.Room.KnowsCommand(command) {
			p.Inform("There is no exit that way\n")
		} else {
			p.Inform("Unknown command\n")
		}
	}

}
