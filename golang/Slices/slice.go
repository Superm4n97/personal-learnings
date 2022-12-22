package main

import "fmt"

func main() {
	/*DECLARATION
	var x[]float64
	x := make([]float64, 5), x is the slice, and underlying array len = 5
	x := make(float64, 5, 10), slice 0-5 and array len 10

	var arr[5] float64 -> array
	x = arr[0:5] -> slice length is same as array
	x = arr[0:5] = arr[:5] = arr[0:] = arr[:len(arr)]
	*/

	/*FUNCTIONS
	APPEND:
	newSlice = append(slice1, p) -> p is added after slice1 and new slice is created
	COPY:
	copy(slice2,slice1) -> then slice2 will have the values of slice1. if slice1 is larger than slice2 then
		the values will be overflowed
	*/

	slice1 := []int{6, 7, 8}
	slice2 := make([]int, 5, 6)
	fmt.Println(slice1)
	fmt.Println(slice2)
	copy(slice2, slice1)
	fmt.Println(slice2)
	var x []float64
	fmt.Println(len(x))
	fmt.Println(slice1[2]) //slices can be accessed like arrays

}
