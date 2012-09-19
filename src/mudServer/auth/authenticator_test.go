package auth

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

type MappedUserStore struct {
	Store map[string]string
}

func (u MappedUserStore) IsUser(username string, password string) bool {
	return u.Store[username] == password
}

func TestLogsIn(t *testing.T) {
	const username, password = "brian", "something"
	const expectedOutput = "Welcome to gomud!\nUsername: Password: Welcome!\n"

	inReader := strings.NewReader(fmt.Sprintf("%s\n%s\n", username, password))
	outWriter := bytes.NewBufferString("")
	mapUserStore := MappedUserStore{map[string]string{username: password}}
	a := Authenticator{inReader, outWriter, mapUserStore}

	if x, _ := a.Login(); x != username {
		t.Errorf("expected username %v got (%v)", username, x)
	}
	if expectedOutput != outWriter.String() {
		t.Errorf("expected username %v got (%v)", expectedOutput, outWriter)
	}
}
func TestLogInWithBadPassword(t *testing.T) {
	const username, validPassword, password = "brian", "password", "something_wrong"
	const expectedOutput = "Welcome to gomud!\nUsername: Password: I don't know that username/password.\nUsername: "

	inReader := strings.NewReader(fmt.Sprintf("%s\n%s\n\u0000", username, password))
	outWriter := bytes.NewBufferString("")
	mapUserStore := MappedUserStore{map[string]string{username: validPassword}}
	a := Authenticator{inReader, outWriter, mapUserStore}

	a.Login()

	if expectedOutput != outWriter.String() {
		t.Errorf("expected username %v got (%v)", expectedOutput, outWriter)
	}
}
