package main

import (
	"fmt"
)

type Circle struct {
	x, y, r float64
}

func main() {
	/*
		c := new(Circle)
		fmt.Println(c.x, c.y, c.r) // 0 0 0
		c2 := Circle{1, 1, 5}
		fmt.Println(c2.r)

	*/
	c := Circle{1, 1, 3}
	fmt.Println(circleArea(&c))

	fmt.Println(c.center())
}

func circleArea(c *Circle) float64 {
	//c.r = 5 // works
	/*structure pass korle * diye dore normally kaaj kora jaay,
	kintu integer pass korle * ditye dhore abar * diyei update korte hobe
	ba kaaj korte hobe.
	*/
	return c.r * c.r
}

/*it's called method: here the term between func and function name is called
receiver. by declaring this way we can call a function like
c.center().
*/
func (c *Circle) center() (float64, float64) {
	return c.x, c.y
}
