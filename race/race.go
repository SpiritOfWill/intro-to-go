package main

import (
	"fmt"
	"time"
)

var x int // shared(package) variable

func a() {
	x = 1
}

func b() {
	x = 2
}

func race() {
	go a()
	go b()
}

func main() {
	race()
	time.Sleep(time.Nanosecond)
	fmt.Println(x)
}
