package main

import "fmt"

// type declaration,
type Filter func(string) string

//func sayHelloWithFilter(name string, filter func(string) string) {
func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Andri", spamFilter)

	sayHelloWithFilter("Anjing", spamFilter)
}
