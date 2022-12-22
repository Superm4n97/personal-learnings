package main

import (
	"fmt"
	"time"
)

func ponger(c chan string) {
	for i := 1; ; i++ {
		c <- "pong"
	}
}

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Millisecond * 1000)
	}
}

func pinger(c chan string) {
	for i := 1; ; i++ {
		c <- "ping"
	}
}

func main() {
	/*
		jokhon function gulo concurrently cholte thake tokhon tader main function theke control korar jonno channel
		use kora hoy.
	*/
	var c chan string = make(chan string)
	go pinger(c)
	go ponger(c)
	go printer(c)
	fmt.Println("Finished")

	var input string
	fmt.Scanln(&input)
}
