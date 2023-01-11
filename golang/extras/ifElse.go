package main

import "fmt"

func main() {
	//ifElseEvenOdd()
	semicolonInIfElse()
}

func semicolonInIfElse() {
	a := true
	b := false

	b = a && b

	/*
		the simple operation will took place first then the comparison will be
		checked.
	*/
	if fmt.Println("simple operation"); b {
		fmt.Println(" true ")
	} else {
		fmt.Println("false")
	}
}

func ifElseEvenOdd() {
	i := 0
	for i <= 10 {
		if i == 0 {
			fmt.Println(i, " zero")
		} else if i%2 == 1 {
			fmt.Println(i, " odd")
		} else {
			fmt.Println(i, "even")
		}
		i++
	}
}
