package main

import (
	"code.google.com/p/go.crypto/bcrypt"
	"flag"
	"encoding/json"
	"os"
)

func userListFrom(fileName string) map[string]string {
	users := make(map[string]string)
	fi,err := os.Stat(fileName)
	if err != nil {
		return users
	}
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	b := make([]byte, fi.Size())
	_, err = f.Read(b)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &users)
	if err != nil {
		panic(err)
	}
	return users
}
func writeUserList(users map[string]string, fileName string){
	b,err := json.Marshal(users)
	if err != nil {
		panic(err)
	}
	f,err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil{
		panic(err)
	}
	defer f.Close()
	f.Write(b)
}

var userName *string
var password *string
var fileName *string
func init(){
	userName = flag.String("user", "", "Username to add")
	password = flag.String("pass", "", "Password for new user")
	fileName = flag.String("pwdfile", "pwdfile", "Password File")
}
func main(){
	flag.Parse()
	if *userName == "" {
		panic("No Username!")
	}

	users := userListFrom(*fileName)

	pass,_ := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	users[*userName] = string(pass)

	writeUserList( users, *fileName)
}
