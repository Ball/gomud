package auth

import (
  "bufio"
  "io"
)

type Authenticator struct {
  In io.Reader
  Out io.Writer
}
func (a *Authenticator) Login() (string,error){
  in := bufio.NewReader(a.In)
  user,_,_ :=  in.ReadLine()
  //_,_,_ := in.ReadLine()
  return string(user), nil
}
