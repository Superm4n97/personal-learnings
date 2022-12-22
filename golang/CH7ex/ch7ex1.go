package CH7ex

func c7ex1sum(num []int) int {
	sum := int(0)
	for _, value := range num {
		sum += value
	}
	return sum
}

func c7ex2Half(num int) (int, bool) {
	return num / 2, num%2 == 0
}

func c7ex3(nums ...int) int {
	greatestNumber := 0
	for _, val := range nums {
		if val > greatestNumber {
			greatestNumber = val
		}
	}
	return greatestNumber
}

func fibonacci(n int64) int64 {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
