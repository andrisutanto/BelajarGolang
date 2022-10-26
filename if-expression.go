package main

import "fmt"

func main() {
	name := "Andri"

	if name == "Andri" {
		fmt.Println("Hello Andri")
	} else if name == "Andre" {
		fmt.Println("Hello Andre")
	} else {
		fmt.Println("Hello selain Andri")
	}

	// SHORT STATEMENT
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
