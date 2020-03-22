package main

import (
	"fmt"
)

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("send email to user : %s<%s>\n",
		u.name,
		u.email,
	)
}

func main() {
	u := user{"Bill", "bill@email.com"}
	sendNotification(&u)
}

func sendNotification(n *notifier) {
	n.notify()
}

/*
go build ./listing36_nofify.go
# command-line-arguments
./listing36_nofify.go:25:19: cannot use &u (type *user) as type *notifier in argument to sendNotification:
	*notifier is pointer to interface, not interface
./listing36_nofify.go:29:3: n.notify undefined (type *notifier is pointer to interface, not interface)
*/
