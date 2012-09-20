package driver

import (
	"bytes"
	"strings"
	"testing"
)

func TestExits(t *testing.T) {
	const input = "exit\n"
	const expectedOutput = "__SHOULD_BE_A_ROOM_MESSAGE__\nThanks for playing!\n"
	inReader := strings.NewReader(input)
	outWriter := bytes.NewBufferString("")
	p := PlayerConnection { inReader, outWriter, "tony" }
	p.Play()

	if expectedOutput != outWriter.String() {
		t.Errorf("expected %v go (%v)", expectedOutput, outWriter)
	}
}

