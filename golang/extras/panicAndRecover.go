package main

import (
	"fmt"
)

func main() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()

	panicFunction()

	fmt.Println("is this line will not be printed!!")
}

func panicFunction() {
	panic("There is a panic")
}
