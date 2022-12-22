package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "half_second"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "two_second"
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		//msg1 := <-c1:
		//fmt.Println(msg1)
		//msg2 := <-c2:
		//fmt.Println(msg2)
		//in the above code msg2 will wait for msg1 to complete. hence although c1 generates for every half second, it will wait for
		//msg2 to finished. Hence, we use select it will check whether any of the message received something. If they do it will print them
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}

}
