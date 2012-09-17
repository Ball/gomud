package auth

import (
  "bufio"
  "errors"
  "io"
)
type UserStore interface {
  IsUser(username string, password string) bool;
}

type Authenticator struct {
  In io.Reader
  Out io.Writer
  Users UserStore
}
func (a *Authenticator) Login() (string,error){
  io.WriteString(a.Out,"Welcome to gomud!\n")
  in := bufio.NewReader(a.In)
  for {
    io.WriteString(a.Out, "Username: ")
    user,_,err :=  in.ReadLine()
    if nil != err {
      return "", err
    }
    if string(user) == "\u0000" {
	    break
    }
    io.WriteString(a.Out, "Password: ")
    pass,_,_ := in.ReadLine()
    if a.Users.IsUser( string(user), string(pass))  {
      io.WriteString(a.Out, "Welcome!\n")
      return string(user), nil
    } else {
      io.WriteString(a.Out, "I don't know that username/password.\n")
      //return "", errors.New("Failed to log in")
    }
  }
  //_,_,_ := in.ReadLine()
  return "", errors.New("Failed to log in")
}
