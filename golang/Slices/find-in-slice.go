package main

import "fmt"

func findInSlice() {
	//sl := []string{"hello","rasel"}
	//sort.SearchStrings(sl,"hi")
	// you can search a element in a sorted slice otherwise you have to iterate the whole slice
}

func findByMap() {
	mp := make(map[string]int)
	sl := []string{"foo", "bar"}
	for _, st := range sl {
		mp[st] = 1
	}
	fmt.Println(mp["foo"])
	fmt.Println(mp["hi"])
}

func main() {
	findInSlice()
	findByMap()
}
