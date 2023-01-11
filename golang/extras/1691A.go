package main

import (
	"bufio"
	. "fmt"
	"os"
)

/*
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tc int
	Fscan(in, &tc)

	for _t := 0; _t < tc; _t++ {
		var n int
		Fscan(in, &n)

		var numOfEven, numOfOdd, ans int

		for i := 0; i < n; i++ {
			var temp int
			Fscan(in, &temp)

			if temp%2 == 0 {
				numOfEven++
			} else {
				numOfOdd++
			}
		}
		ans = numOfEven
		if numOfOdd > numOfEven {
			ans = numOfOdd
		}
		Fprintln(out, n-ans)
	}
}

*/

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tc int
	Fscan(in, &tc)

	for _t := 0; _t < tc; _t++ {
		var n int
		Fscan(in, &n)

		num := make([]int, n, n)
		parityCounter := make(map[int]int)
		for i := 0; i < n; i++ {
			Fscan(in, &num[i])
			parityCounter[num[i]%2]++
		}
		var ans = parityCounter[0]
		if parityCounter[1] > ans {
			ans = parityCounter[1]
		}
		Fprintln(out, n-ans)
	}
}
