package main

import (
	"fmt"
)

func main() {
	//printingThe4thElement()
	smallestNumberInArray()
}

func printingThe4thElement() {
	//array
	tempArray := [5]int{5, 4, 3, 2, 1}
	//var tempArray [5]int
	fmt.Println(tempArray[3])

	//var tempSlice = []int{5, 4, 3, 2, 1}
	tempSlice := []int{5, 4, 3, 2, 1}
	fmt.Println(tempSlice[3])
}

func smallestNumberInArray() {
	x := []int{
		48, 96, 86, 68, 57, 82,
		63, 70, 37, 34, 83, 27,
		19, 97, 9, 17,
	}

	var mnValue int = x[0]
	for i := 0; i < len(x); i++ {
		if x[i] < mnValue {
			mnValue = x[i]
		}
	}
	fmt.Println(mnValue)
}
