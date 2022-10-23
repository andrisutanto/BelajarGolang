package main

import "fmt"

func main() {
	name := "Andri"

	switch name {
	case "Andri":
		fmt.Println("Hello Andri")
	case "Andre":
		fmt.Println("Hello Andre")
	default:
		fmt.Println("Hi, Who Are You?")
	}

	// SHORT STATEMENT
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	// SWITCH WITHOUT CONDITION
	length2 := len(name)

	switch {
	case length2 > 10:
		fmt.Println("Nama kepanjangan")
	case length2 > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}

}
