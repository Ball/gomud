package driver

import (
	"bytes"
	"strings"
	"testing"
)

func TestExiting(t *testing.T) {
	room := &Room{Description:"Some room description"}
	const input = "exit\n"
	expectedOutput := room.Description + "\nThanks for playing!\n"
	inReader := strings.NewReader(input)
	outWriter := bytes.NewBufferString("")
	p := PlayerConnection { inReader, outWriter, "tony", room }
	p.Play()

	if expectedOutput != outWriter.String() {
		t.Errorf("expected %v go (%v)", expectedOutput, outWriter)
	}
}

func TestListingExits(t *testing.T){
	room := &Room{Description: "Some room description"}
	const input = "north\nexit\n"
	expectedOutput := room.Description + "\nThere is no exit that way\nThanks for playing!\n"
	inReader := strings.NewReader(input)
	outWriter := bytes.NewBufferString("")
	p := PlayerConnection { inReader, outWriter, "tony", room }
	p.Play()

	if expectedOutput != outWriter.String() {
		t.Errorf("expected %v got (%v)", expectedOutput, outWriter)
	}
}
