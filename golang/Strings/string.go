package main

import "fmt"

func main() {
	fmt.Println("Simple string")
	fmt.Println("endline -> Simple \n \t string")
	fmt.Println(`non endline -> Simple \n \t string`)
	fmt.Println("ABHello"[1]) // Zero indexed
}
