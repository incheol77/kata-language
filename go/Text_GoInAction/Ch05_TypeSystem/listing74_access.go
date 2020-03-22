package main

import (
	"entities2"
	"fmt"
)

func main() {
	a := entities2.Admin{
		Rights: 10,
	}

	// a.user.Name = "Bill" -> error
	a.Name = "Bill"
	a.Email = "bill@email.com"

	fmt.Printf("User: %v\n", a)
}
