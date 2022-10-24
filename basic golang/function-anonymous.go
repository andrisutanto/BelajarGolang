package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcom", name)
	}
}

func main() {

	// ini anonymous functionnya:
	blacklist := func(name string) bool {
		return name == "admin"
		//jadi artinya username admin ini di blocked
	}

	registerUser("admin", blacklist) // ini seolah2 mau register user bernama admin
	registerUser("andri", blacklist)

	// cara susahnya:
	registerUser("root", func(name string) bool {
		return name == "root"
	})
}
