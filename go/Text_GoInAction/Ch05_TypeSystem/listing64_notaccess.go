package main

import (
	"counters"
	"fmt"
)

func main() {
	counter := counters.alertCounter(10)
	fmt.Printf("Counter : %d\n", counter)

	/*
			go run ./listing64_notaccess.go
		# command-line-arguments
		./listing64_notaccess.go:9:13: cannot refer to unexported name counters.alertCounter
		./listing64_notaccess.go:9:13: undefined: counters.alertCounter
	*/
}
