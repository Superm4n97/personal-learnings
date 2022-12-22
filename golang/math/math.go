package math

const Pi = 3.1416

func Max(args ...int) (ret int) {
	ret = args[0]
	for _, val := range args {
		if val > ret {
			ret = val
		}
	}
	return
}

func Min(args ...int) (ret int) {
	ret = args[0]
	for _, val := range args {
		if val < ret {
			ret = val
		}
	}
	return
}

func Average(args []float64) (ret float64) {
	ret = 0
	for _, val := range args {
		ret += val
	}
	ret /= float64(len(args)) + 1
	return
}

func AddTwo(num int) int {
	return num + 2
}
