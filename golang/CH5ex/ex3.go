package main

import "fmt"

func main() {
	fizzBuzz()
}

func fizzBuzz() {
	i := 1

	for i <= 100 {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print("FizzBuzz ")
		} else if i%3 == 0 {
			fmt.Print("Fizz ")
		} else if i%5 == 0 {
			fmt.Print("Buzz ")
		} else {
			fmt.Print(i, " ")
		}
		i++
	}
}
