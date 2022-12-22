package main

import "fmt"

/*
ekta variable er address ber korar jonno & use kora hoy,
ar *address diye oi address er variable ke bujhay
*/

func main() {
	x := 5
	pointerPassing(&x)
	fmt.Println(x)

	ptr := new(int)
	pointerPassing(ptr)
	fmt.Println(*ptr)
}

func pointerPassing(pointerOfX *int) {
	*pointerOfX = 3
}
