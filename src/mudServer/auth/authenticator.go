package auth

import (
	"bufio"
	"errors"
	"io"
)

type UserStore interface {
	IsUser(username string, password string) bool
}

type Authenticator struct {
	In    io.Reader
	Out   io.Writer
	Users UserStore
}

func (a *Authenticator) Login() (string, error) {
	io.WriteString(a.Out, "Welcome to gomud!\r\n")
	in := bufio.NewReader(a.In)
	for {
		io.WriteString(a.Out, "Username: ")
		user, _, err := in.ReadLine()
		if nil != err {
			return "", err
		}
		userName := string(user)
		if string(user) == "\u0000" {
			break
		}
		io.WriteString(a.Out, "Password: ")
		pass, _, _ := in.ReadLine()
		password := string(pass)
		if a.Users.IsUser(userName, password) {
			io.WriteString(a.Out, "Welcome!\n")
			return string(user), nil
		} else {
			io.WriteString(a.Out, "I don't know that username/password.\r\n")
		}
	}
	return "", errors.New("Failed to log in")
}
