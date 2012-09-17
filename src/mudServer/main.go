package main

import (
  "bufio"
  "fmt"
  "net"
  "os"
  "strings"
  "mudServer/auth"
)

func main(){
  // TODO : config for port
  service := "localhost:1201"
  tcpAddr, err := net.ResolveTCPAddr("ip4", service)
  checkError(err)

  listener, err := net.ListenTCP("tcp", tcpAddr)
  checkError(err)

  for {
    conn, err := listener.Accept()
    if err != nil {
      continue
    }
    go handleClient(conn)
  }
}

func handleClient(conn net.Conn){
  defer conn.Close()
  // TODO : Hand off to login system
  a := auth.Authenticator{conn, conn}
  username, _ := a.Login()
  fmt.Printf("User %s Logged In\n", username)

  // TODO : Hand off to player connection
  buffIn := bufio.NewReader(conn)
  buffOut := bufio.NewWriter(conn)
  for {
    line, _, err := buffIn.ReadLine()
    if err != nil {
      fmt.Printf("Error: %s\n", err.Error())
      return
    } else {
      fmt.Printf("<<%s>>\n", line)
    }
    if strings.HasPrefix(string(line), "exit"){
      return
    }
    _,err = conn.Write(line)

    if err != nil {
      fmt.Printf("Error: %s\n", err.Error())
      return
    } else {
      buffOut.Flush()
    }
  }
}
func checkError(err error){
  if err != nil {
    fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
    os.Exit(1)
  }
}
