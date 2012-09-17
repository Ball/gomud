package auth
import (
 "bytes"
 "fmt"
 "strings"
 "testing"
 )

func TestLogsIn(t *testing.T){
  const username, password = "brian","something"
  inReader := strings.NewReader(fmt.Sprintf("%s\n%s\n", username, password))
  outWriter := bytes.NewBufferString("")
  a := Authenticator{inReader, outWriter}
  
  if x,_ := a.Login(); x != username {
    t.Errorf("expected %v got (%v)", username, x)
  }
}
