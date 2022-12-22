package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func server(ready chan bool) {
	// listen on a port
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println("failed to listen to 9999, reason: ", err)
		return
	}

	ready <- true

	for {
		// accept a connection
		c, err := ln.Accept()
		if err != nil {
			fmt.Println("failed to accept connection, reason: ", err)
			continue
		}

		// handle the connection
		go handleServerConnection(c)
	}
}

func handleServerConnection(c net.Conn) {
	//receive the message
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println("failed to decode message, reason: ", err)
	} else {
		fmt.Println("Received: ", msg)
	}

	c.Close()
}

func client(serverReady chan bool) {
	<-serverReady //
	c, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println("failed to connect to server, reason: ", err)
		return
	}

	// send message
	msg := "hello World"
	fmt.Println("sending message...")
	//time.Sleep(time.Second)

	err = gob.NewEncoder(c).Encode(msg)

	if err != nil {
		fmt.Println("failed to encode msg, reason: ", err)
	}

	c.Close()
}

func main() {
	ch := make(chan bool)
	go server(ch)
	go client(ch)

	var tempString string
	fmt.Scanln(&tempString)
}
