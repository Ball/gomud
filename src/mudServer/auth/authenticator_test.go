package auth

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)


type PasswordMap map[string]string
func (s PasswordMap) IsUser(username string, password string) bool {
	return s[username] == password
}

func TestLogsIn(t *testing.T) {
	const username, password = "brian", "something"
	const expectedOutput = "Welcome to gomud!\r\nUsername: Password: Welcome!\n"

	inReader := strings.NewReader(fmt.Sprintf("%s\n%s\n", username, password))
	outWriter := bytes.NewBufferString("")
	a := Authenticator{inReader, outWriter, PasswordMap{username: password}}

	if x, _ := a.Login(); x != username {
		t.Errorf("expected username %v got (%v)", username, x)
	}
	if expectedOutput != outWriter.String() {
		t.Errorf("expected %v got (%v)", expectedOutput, outWriter)
	}
}
func TestLogInWithBadPassword(t *testing.T) {
	const username, validPassword, password = "brian", "password", "something_wrong"
	const expectedOutput = "Welcome to gomud!\r\nUsername: Password: I don't know that username/password.\r\nUsername: "

	inReader := strings.NewReader(fmt.Sprintf("%s\n%s\n\u0000", username, password))
	outWriter := bytes.NewBufferString("")
	a := Authenticator{inReader, outWriter, PasswordMap{username: validPassword}}

	a.Login()

	if expectedOutput != outWriter.String() {
		t.Errorf("expected username %v got (%v)", expectedOutput, outWriter)
	}
}
