package main

import "fmt"

func main() {
	//var x[5] float64
	x := [5]float64{
		1.2,
		1.4,
		1.5,
		1.8,
		1.9,
	}

	/*i for index and val for value of x
	the following line is similar to
	i =: 0
	for i<len(x) { value = x[i]
	*/

	/*
		for _,_ := range x{}
		it is a pair, first element is the index and the second element is the value in that index
	*/

	var total float64 = 0.0
	for _, value := range x {
		total += value
	}

	fmt.Println(total / float64(len(x)))
}
