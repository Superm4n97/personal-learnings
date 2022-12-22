package main

import "fmt"

func main() {
	//defer execute after all the task is completed in this function
	defer defFn1()
	defer defFn2() // multiple defer works as stack, last defer will execute first
	defFn3()
}

func defFn1() {
	fmt.Println("First Function")
}

func defFn2() {
	fmt.Println("Second Funcion")
}

func defFn3() {
	fmt.Println("Third Funcion")
}
