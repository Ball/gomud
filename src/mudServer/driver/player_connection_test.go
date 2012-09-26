package driver

import (
	"bytes"
	"strings"
	"testing"
)

func TestExits(t *testing.T) {
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

