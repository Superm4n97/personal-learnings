package main

import "fmt"

func main() {
	//fn1()

	// in closure a function can also return a funcion
	nextEven := returnableFunction()
	fmt.Println(nextEven()) //0
	fmt.Println(nextEven()) //2
	fmt.Println(nextEven()) //4
}

func fn1() {
	// a function can have another function inside it
	// it's a local function, it can not be accessed by outside f1()
	add := func(a int, b int) int {
		return a + b
	}

	fmt.Println(add(10, 20))
}

//complex one
func returnableFunction() func() int {
	i := int(0)
	return func() (ret int) {
		ret = i
		i += 2
		return
	}
}
