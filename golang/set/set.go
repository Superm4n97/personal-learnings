package main

import (
	"fmt"
	"gomodules.xyz/sets"
)

func initialization() {
	set := sets.NewString()
	set.Insert("hi")
	fmt.Println(set.Has("hi"))
}

func main() {
	initialization()
}
