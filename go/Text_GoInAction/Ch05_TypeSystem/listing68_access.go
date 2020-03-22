package main

import (
	"counters2"
	"fmt"
)

func main() {
	counter := counters2.New(10)
	fmt.Printf("Counter : %d\n", counter)
}
