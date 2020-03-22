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
	fmt.Printf("Sending email to user : %s <%s>\n",
		u.name,
		u.email)
}

type admin struct {
	user
	level string
}

func (a *admin) notify() {
	fmt.Printf("Sending email to admin : %s <%s>\n",
		a.name,
		a.email)
}

func main() {
	ad := admin{
		user: user{
			name:  "john smith",
			email: "john@email.com",
		},
		level: "super",
	}

	sendNotification(&ad)
	sendNotification(&ad.user)
}

func sendNotification(n notifier) {
	n.notify()
}
