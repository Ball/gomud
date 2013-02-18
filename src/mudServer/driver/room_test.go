package driver

import (
	"bytes"
	"strings"
	"testing"
	"time"
)
func TestAnnouncementOfArrivalToOthers(t *testing.T){
	room := &Room{Description:"Some room description"}
	expectedOutput := "New Player has entered the room"

	outWriter := bytes.NewBufferString("")
	establishedPlayer := PlayerConnection { In: strings.NewReader(""),
	                                        Out: outWriter,
						UserName: "Paul",
						Room: room }
	establishedPlayer.Play()

	newPlayer := PlayerConnection { In: strings.NewReader("exit\n"),
	                                Out: bytes.NewBufferString(""),
					UserName: "New Player",
					Room: room }
	newPlayer.Play()
	time.Sleep(2 * time.Millisecond)
	output := outWriter.String()
	if ! strings.Contains(output, expectedOutput) {
		t.Errorf("expected %v got (%v)", expectedOutput, output)
	}
}
func TestDoesNotAnnounceArrivalToSelf(t *testing.T){
	room := &Room{Description:"Some room description"}
	expectedOutput := "New Player has entered the room"

	outWriter := bytes.NewBufferString("")
	establishedPlayer := PlayerConnection { In: strings.NewReader(""),
	                                        Out: bytes.NewBufferString(""),
						UserName: "Paul",
						Room: room }
	establishedPlayer.Play()

	newPlayer := PlayerConnection { In: strings.NewReader("exit\n"),
	                                Out: outWriter,
					UserName: "New Player",
					Room: room }
	newPlayer.Play()
	if strings.Contains(outWriter.String(), expectedOutput) {
		t.Errorf("expected %v got (%v)", expectedOutput, outWriter)
	}
}
