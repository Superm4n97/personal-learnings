package main

import "fmt"

var fName = "Rasel" //variable (global)
const age = 25      //constant. once constant is declared, it can not be changed

func main() {
	var sName = " Hossain" // variable local
	firstName()
	fmt.Println(sName)
	fmt.Println(age)
	multipleVariables()
}

func firstName() {
	fmt.Print(fName)
}

func multipleVariables() {
	var (
		a = 5
		b = 10
		c = 15
	)
	fmt.Println(a + b + c)

	/*
		so , var newVariable = b, is similar to
		newVariable := b
		you dont have to declear var anymore
	*/
	newVariable := 5
	fmt.Println(newVariable)
}
