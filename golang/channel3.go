package main

import (
	"fmt"
)

func hi(id int, c chan int) {
	for x := range c {
		fmt.Println(x)
	}
}

func main() {
	//c := make(chan int)
	//var c chan int

	var mp map[int]int

	if mp == nil {
		fmt.Println("is nill")
	}

	mp = make(map[int]int)
	if mp != nil {
		fmt.Println("is not nil")
	}

	var inpt string
	fmt.Scanln(&inpt)
}
