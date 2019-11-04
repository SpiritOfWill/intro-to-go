package main

import (
	"fmt"
	"time"
)

var x int // shared(package) variable

func a() {
	x = 1
	fmt.Println(x)
}

func b() {
	x = 2
	fmt.Println(x)
}

func main() {
	go a()
	go b()

	time.Sleep(2 * time.Second)
}
