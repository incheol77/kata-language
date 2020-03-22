package main

import (
	"entities"
	"fmt"
)

func main() {
	u := entities.User{
		Name:  "Bill",
		email: "bill@email.com",
	}

	fmt.Printf("User: %v\n", u)

	/*
			 go run ./listing71_noaccess.go
		# command-line-arguments
		./listing71_noaccess.go:11:3: unknown field 'email' in struct literal of type entities.User
	*/
}
