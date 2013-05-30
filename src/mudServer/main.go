package main

import (
	"flag"
	"fmt"
	"mudServer/auth"
	"mudServer/driver"
	"net"
	"os"
	"strconv"
)

var defaultRoom *driver.Room
var port = flag.Int("port", 1201, "The connection port")

func main() {
	flag.Parse()
	service := "localhost:" + strconv.Itoa(*port)
	tcpAddr, err := net.ResolveTCPAddr("", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	loadDefaultRoom()

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func loadDefaultRoom() {
	defaultRoom = new(driver.Room)
	defaultRoom.Description = "A simple lobby"
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	pwdStore := auth.JsonUserStore{"mudlib/pwdfile"}
	a := auth.Authenticator{conn, conn, pwdStore}
	username, err := a.Login()
	if err == nil {
		fmt.Printf("User %s Logged In\n", username)
	} else {
		fmt.Printf("Failed to log in %s\n", err.Error())
		return
	}

	(&driver.PlayerConnection{conn, conn, username, defaultRoom}).Play()
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
