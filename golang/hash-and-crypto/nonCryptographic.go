package main

import (
	"fmt"
	"hash/crc32"
)

func main() {
	h := crc32.NewIEEE()

	//writing in write interface
	h.Write([]byte("text"))

	//Sum32 returns a uint32 value
	v := h.Sum32()

	fmt.Println(h)
	fmt.Println(v)
}
