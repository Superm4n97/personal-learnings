package main

import (
	"fmt"
)

type B struct {
	n int
}

type in interface {
	Add()
}

type A struct {
	a int
	B
	in
}

type inT struct {
}

func (i *inT) Add() {

}

func NewAWithInT() *A {
	return &A{
		a: 4,
		B: B{
			n: 5,
		},
		in: &inT{},
	}
}

func New() A {
	return A{
		a: 4,
		B: B{
			n: 5,
		},
	}
}

func NewP() *A {
	return &A{
		a: 4,
		B: B{
			n: 5,
		},
	}
}

func main() {
	fmt.Println("Hello World")
}
