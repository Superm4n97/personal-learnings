package main

import "fmt"

func main() {
	loopIteration()
}

func loopIteration() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}
}
