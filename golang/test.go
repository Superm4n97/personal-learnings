package main

import (
	"fmt"
)

func formating() {
	s := "Rasel"
	println(`
hi! my name is `, s, ` - hello `, s)
}

func sudoFunc(m *map[int]int) {
	*m = make(map[int]int)
}

func reference() {
	var m map[int]int // m is nil
	//m := make(map[int]int) // m is not nil

	sudoFunc(&m)

	if m == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
}

func main() {
	reference()
}
