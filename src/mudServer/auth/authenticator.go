package auth

import (
	"bufio"
	"code.google.com/p/go.crypto/bcrypt"
	"encoding/json"
	"errors"
	"io"
	"os"
)

type UserStore interface {
	IsUser(username string, password string) bool
}
type JsonUserStore struct {
	FileName string
}

func (a JsonUserStore) IsUser(username string, password string) bool {
	users := make(map[string]string)

	fi, err := os.Stat(a.FileName)
	f, err := os.Open(a.FileName)
	if err != nil {
		return false
	}
	b := make([]byte, fi.Size())
	_, err = f.Read(b)
	if err != nil {
		return false
	}
	err = json.Unmarshal(b, &users)
	if err != nil {
		return false
	}
	err = bcrypt.CompareHashAndPassword([]byte(users[username]), []byte(password))
	return err == nil
}

type Authenticator struct {
	In    io.Reader
	Out   io.Writer
	Users UserStore
}

func (a *Authenticator) Login() (string, error) {
	io.WriteString(a.Out, "Welcome to gomud!\n")
	in := bufio.NewReader(a.In)
	for {
		io.WriteString(a.Out, "Username: ")
		user, _, err := in.ReadLine()
		if nil != err {
			return "", err
		}
		if string(user) == "\u0000" {
			break
		}
		io.WriteString(a.Out, "Password: ")
		pass, _, _ := in.ReadLine()
		if a.Users.IsUser(string(user), string(pass)) {
			io.WriteString(a.Out, "Welcome!\n")
			return string(user), nil
		} else {
			io.WriteString(a.Out, "I don't know that username/password.\n")
		}
	}
	return "", errors.New("Failed to log in")
}
