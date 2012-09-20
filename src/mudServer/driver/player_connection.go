package driver

import (
	"bufio"
	"io"
)

type PlayerConnection struct {
	In io.Reader
	Out io.Writer
	UserName string
}

func (p *PlayerConnection) Inform(msg string) {
	io.WriteString(p.Out, msg)
}

func (p *PlayerConnection) Play(){
	p.Inform("__SHOULD_BE_A_ROOM_MESSAGE__\n")
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
		}
	}

}
