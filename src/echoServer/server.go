package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)
// socket should be a config issue.
func main() {
	service := "localhost:1201"
	tcpAddr, err := net.ResolveTCPAddr("ip4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			// figure out a better error handling situation.
			// i dunno, like logging for one?
			continue
		}
		go handleClient(conn)
	}
}
func handleClient(conn net.Conn) {
	// handler should create a "PlayerConnection"
	// It constructs from a buffered reader and writer and a username/password
	// if I call .SendFrom(sender, message) it should filter out
	// if I send it commands, it processes them and calls other objects
	// To get a username/password, it should use a LoginContext to handle back and forth retries
	defer conn.Close()
	buffered := bufio.NewReader(conn)
	buffOut := bufio.NewWriter(conn)
	for {
		line, _, err := buffered.ReadLine()
		// is err.Error() == EOF, its just a premature disconnect.
		if err != nil {
			fmt.Printf("Oh, crap: %s\n", err.Error())
			return
		} else {
			fmt.Printf("<<%s>>\n", line)
		}
		if strings.HasPrefix(string(line), "exit") {
			return
		}

		_, err2 := conn.Write(line)

		if err2 != nil {
			fmt.Printf("Well, crap: %s\n", err2.Error())
			return
		} else {
			buffOut.Flush()
		}
	}
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
