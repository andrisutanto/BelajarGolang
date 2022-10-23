package main

import "fmt"

func getFullName() (string, string) { //karena ada 2 return valuenya, maka tuliskan string sebanyak 2x
	return "Andri", "Sutanto"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	// kalau ingin menghiraukan value yang tdk dibutuhkan, bisa menggunakan _
	// contoh hanya ingin ambil firstname saja:

	firstName2, _ := getFullName() // -> jadi last namenya di ignore
	fmt.Println(firstName2)
}
