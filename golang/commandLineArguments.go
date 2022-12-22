package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {
	//define flags
	maxp := flag.Int("max", 6, "the max value")

	//msg := flag.String("tempflag", "git status", "shows the git status")

	// parse
	flag.Parse()

	//Genearate a number between 0 and max
	fmt.Println(rand.Intn(*maxp))
	//fmt.Println(*maxp)
	//fmt.Println(*msg)
}
