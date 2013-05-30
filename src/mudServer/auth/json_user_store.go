package auth

import (
	"code.google.com/p/go.crypto/bcrypt"
	"encoding/json"
	"os"
)

type JsonUserStore struct {
	FileName string
}

func bytesFromFile(filename string) ([]byte, error) {
	fi, err := os.Stat(filename)
	if err != nil {
		return nil, err
	}
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	b := make([]byte, fi.Size())
	_, err = f.Read(b)

	return b, err
}
func userMapFromFile(filename string) (map[string]string, error) {
	users := make(map[string]string)
	b, err := bytesFromFile(filename)

	if err != nil {
		return users, err
	}

	err = json.Unmarshal(b, &users)
	return users, err
}
func (a JsonUserStore) IsUser(username string, password string) bool {
	users, err := userMapFromFile(a.FileName)
	if err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(users[username]), []byte(password))
	return err == nil
}
