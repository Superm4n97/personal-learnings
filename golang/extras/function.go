package main

import "fmt"

func main() {
	//fmt.Println(minMax(10, 15))
	//fmt.Println(findAverage([]float64{1.1, 1.3, 1.4, 1.7}))
	fmt.Println(variadicFunction(1, 2, 3, 4, 5))
	//fmt.Println(variadicFunction([]int{1,2,3,4,5})) // won't work
}

func minMax(a int, b int) (int, int) {
	if a > b {
		return b, a
	} else {
		return a, b
	}
}

func findAverage(arr []float64) float64 {
	var totalSum float64 = 0.0
	for i := 0; i < len(arr); i++ {
		totalSum += arr[i]
	}

	return totalSum / float64(len(arr))
}

func variadicFunction(parameter ...int) int {
	sum := 0
	for _, value := range parameter {
		sum += value
	}
	return sum
}
