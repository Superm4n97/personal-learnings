package main

import (
	"fmt"
	"reflect"
)

func Add2(num int) int {
	return num + 2
}

func main() {
	//fmt.Println("hello world")
	//package calling
	//hello.SayHello()
	//	fmt.Println(`1
	//2 3
	//4 5 6`)
	//fmt.Println(math.Max(4, 5))

	fmt.Println("Anthing can be printed")
	fmt.Println("After adding 2 the result is", Add2(2))
	//fmt.Println(math.Max(4, 5))

	a := 5

	fmt.Println(reflect.TypeOf(reflect.TypeOf(a)))
	//fmt.Println(a.(reflect.Type))
}
