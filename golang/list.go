package main

import (
	"container/list"
	"fmt"
	"reflect"
)

func main() {
	/*
		a list can contain any type of data. it is a doubly-link list, so link list er moto kore iterate korte hobe.
	*/
	var x list.List
	x.PushBack(1)
	x.PushBack("hi")
	x.PushBack(3.6)

	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value, reflect.TypeOf(e.Value))
	}
}
