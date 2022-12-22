package main

import (
	"fmt"
	"reflect"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

//this line made the Person Slice a type because we have to pass a type in the sort.Sort() function
type ByName []Person

// for sorting you have to declare Len(), Less() and Swap() functions
func (this ByName) Len() int {
	return len(this)
}
func (this ByName) Less(i, j int) bool {
	return this[i].Name < this[j].Name
}
func (this ByName) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

type ByAge []Person

func (this ByAge) Len() int {
	return len(this)
}
func (this ByAge) Less(i, j int) bool {
	return this[i].Age < this[j].Age
}
func (this ByAge) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func main() {
	kids := []Person{
		{"Me", 8},
		{"Hi", 9},
		{"Jack", 10},
	}

	fmt.Println(reflect.TypeOf(Person{}))

	sort.Sort(ByName(kids))
	sort.Sort(ByAge(kids))

	/*can not pass these, because the sort function allow a type to pass through.  */
	//sort.Sort(Person{})
	//sort.Sort(kids)

	fmt.Println(kids)
}
