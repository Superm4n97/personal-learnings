package main

import "fmt"

func main() {
	//way1()
	way2()
}

func way2() {
	i := 3
	for i < 100 {
		fmt.Print(i, " ")
		i += 3
	}
}

func way1() {
	i := 1
	for i <= 100 {
		if i%3 == 0 {
			fmt.Print(i, " ")
		}
		i++
	}
}
