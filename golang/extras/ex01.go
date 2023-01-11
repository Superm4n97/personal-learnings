package main

import (
	"fmt"
	"sync"
	"time"
)

//res := 0
var res int
var mutex sync.Mutex

func increment() {
	for i := 0; i < 20; i++ {
		res++
	}
}

func decrement() {
	for i := 0; i < 20; i++ {
		res++
	}
	//fmt.Println(res)
}

func main() {
	go func() {
		for i := 0; i < 100000; i++ {
			mutex.Lock()
			res++
			mutex.Unlock()
		}
	}()
	go func() {
		for i := 0; i < 100000; i++ {
			mutex.Lock()
			res--
			mutex.Unlock()
		}
	}()

	time.Sleep(time.Second * 10)
	fmt.Println(res)

	var s string
	fmt.Scanln(&s)
}
