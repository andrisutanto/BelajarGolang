package main

import "fmt"

func getFullName2() (firstName string, middleName string, lastName string) {
	// ini berguna, agar saat return, kita tidak usah definisikan satu-persatu untuk return valuenya
	// karena sudah didefiniskan sblmnya
	firstName = "Andri"
	middleName = "Sutanto"
	lastName = "Tan"

	return
}

func main() {
	a, b, c := getFullName2()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
