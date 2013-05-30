package driver

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

func TestAnnouncementOfArrivalToOthers(t *testing.T) {
	room := &Room{Description: "Some room description"}
	expectedOutput := "New Player has entered the room"

	outWriter := bytes.NewBufferString("")
	establishedPlayer := PlayerConnection{In: strings.NewReader(""),
		Out:      outWriter,
		UserName: "Paul",
		Room:     room}
	establishedPlayer.Play()

	newPlayer := PlayerConnection{In: strings.NewReader("exit\n"),
		Out:      bytes.NewBufferString(""),
		UserName: "New Player",
		Room:     room}
	newPlayer.Play()
	time.Sleep(2 * time.Millisecond)
	output := outWriter.String()
	if !strings.Contains(output, expectedOutput) {
		t.Errorf("expected %v got (%v)", expectedOutput, output)
	}
}
func TestDoesNotAnnounceArrivalToSelf(t *testing.T) {
	room := &Room{Description: "Some room description"}
	expectedOutput := "New Player has entered the room"

	outWriter := bytes.NewBufferString("")
	establishedPlayer := PlayerConnection{In: strings.NewReader(""),
		Out:      bytes.NewBufferString(""),
		UserName: "Paul",
		Room:     room}
	establishedPlayer.Play()

	newPlayer := PlayerConnection{In: strings.NewReader("exit\n"),
		Out:      outWriter,
		UserName: "New Player",
		Room:     room}
	newPlayer.Play()
	if strings.Contains(outWriter.String(), expectedOutput) {
		t.Errorf("expected %v got (%v)", expectedOutput, outWriter)
	}
}

// Communications in room
func TestSayToOthers(t *testing.T) {
	room := &Room{Description: "Some room description"}
	expectedOutput := "New Player says Hello"

	outWriter := bytes.NewBufferString("")
	establishedPlayer := PlayerConnection{In: strings.NewReader(""),
		Out:      outWriter,
		UserName: "Paul",
		Room:     room}
	establishedPlayer.Play()

	newPlayer := PlayerConnection{In: strings.NewReader("say Hello\nexit\n"),
		Out:      bytes.NewBufferString(""),
		UserName: "New Player",
		Room:     room}
	newPlayer.Play()
	time.Sleep(2 * time.Millisecond)
	output := outWriter.String()
	if !strings.Contains(output, expectedOutput) {
		t.Errorf("expected %v got (%v)", expectedOutput, output)
	}
}

func TestWhisperToOthers(t *testing.T) {
	room := &Room{Description: "Some room description"}
	expectedOutput := "New Player whispers Hello"

	outWriter1 := bytes.NewBufferString("")
	establishedPlayer1 := PlayerConnection{In: strings.NewReader(""),
		Out:      outWriter1,
		UserName: "Paul",
		Room:     room}
	establishedPlayer1.Play()
	outWriter2 := bytes.NewBufferString("")
	establishedPlayer2 := PlayerConnection{In: strings.NewReader(""),
		Out:      outWriter2,
		UserName: "Sharon",
		Room:     room}
	establishedPlayer2.Play()

	newPlayer := PlayerConnection{In: strings.NewReader("whisper Paul Hello\nexit\n"),
		Out:      bytes.NewBufferString(""),
		UserName: "New Player",
		Room:     room}
	newPlayer.Play()
	time.Sleep(2 * time.Millisecond)
	output := outWriter1.String()
	if !strings.Contains(output, expectedOutput) {
		t.Errorf("expected %v got (%v)", expectedOutput, output)
	}
	output = outWriter2.String()
	if strings.Contains(output, expectedOutput) {
		t.Errorf("expected %v got (%v)", expectedOutput, output)
	}
}
