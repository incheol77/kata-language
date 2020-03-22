package main

import (
	"fmt"
)

type duration int

func (d *duration) pretty() string {
	return fmt.Sprintf("Duration : %d", *d)
}

func main() {
	duration(42).pretty()
	/*
	go run ./listing46_duration.go
# command-line-arguments
./listing46_duration.go:14:14: cannot call pointer method on duration(42)
./listing46_duration.go:14:14: cannot take the address of duration(42)
*/
}
