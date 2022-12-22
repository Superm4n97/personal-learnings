package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	var a int
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(a)

	// x := make([]int,0) it creates the variable, and it is not nil
	x := make([]int, 0)

	// it assigns the variable, initial value is nil,
	var s []int

	//s[5] = 10

	fmt.Println(s == nil)
	fmt.Println(x == nil)
	fmt.Println(unsafe.Sizeof(x))
	fmt.Println(unsafe.Sizeof(s))
}
