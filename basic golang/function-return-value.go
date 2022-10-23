package main

import "fmt"

func getHello(name string) string { //string disebelah parameter adalah tipe data returnnya
	return "Hello " + name
}

func main() {
	result := getHello("Andri")
	fmt.Println(result)
}
