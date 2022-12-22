package main

import "fmt"

type Dog struct {
	names string
	age   int
}

func main() {
	/*
		dogs := []Dog{{"tommy", 2}, {"jimmy", 4}}
		allDogs := []Dog{{"bella", 5}}
		for _, dog := range dogs {
			allDogs = append(allDogs, dog)
		}

		a := make([]Dog, len(dogs))
		copy(a, dogs)

		fmt.Println(a)

	*/

	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := a[:]
	a = append(a[:3], a[6:]...)
	fmt.Println(a)
	fmt.Println(b)
}
