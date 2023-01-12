package main

import (
	"fmt"
)

func fn2() {
	for i := 1; i <= 10; i++ {
		//time.Sleep(time.Millisecond * 500)
		fmt.Println("fn2", i)
	}
}

func fn3() {
	for i := 1; i <= 10; i++ {
		//time.Sleep(time.Millisecond * 300)
		fmt.Println("fn3", i)
	}
}

func main() {
	go fn2()
	go fn3()
}
