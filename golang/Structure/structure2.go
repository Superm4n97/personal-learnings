package main

import "fmt"

type Person struct {
	name string
}

type Android struct {
	Person
	version string
}

func (p Person) talk() {
	fmt.Println("Hi, my name is ", p.name)
}

func (a Android) currentVersion() {
	fmt.Println("Current version is ~1.8.1")
}

func main() {
	a := new(Android)
	a.talk() // Embedded types: android has the property of Person
	a.currentVersion()
}
