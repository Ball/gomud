package main
import (
"net"
"os"
"fmt"
)

func main(){
  strEcho := "Hello\n"
  service := "localhost:1201"
  tcpAddr, err := net.ResolveTCPAddr("tcp", service)
  if err != nil {
    println("ResolveTCPAddr failed:", err.Error())
    os.Exit(1)
  }

  conn, err := net.DialTCP("tcp", nil, tcpAddr)
  if err != nil {
    println("Dial Failed:", err.Error())
    os.Exit(1)
  }

  _, err = conn.Write([]byte(strEcho))
  if err != nil {
    println("Write to server failed:", err.Error())
    os.Exit(1)
  }

  fmt.Printf("write to server = <<%s>>\n", strEcho)

  reply := make([]byte, 1024)

  length,err := conn.Read(reply)
  if err != nil {
    println("Write to server failed:", err.Error())
    os.Exit(1)
  }

  fmt.Printf("reply from server = <<%s>>\n", string(reply[:length]))
  conn.Close()
}
