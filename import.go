package main

// begini untuk cara importnya

import (
	"belajar-golang/helper"
	"fmt"
)

func main() {
	helper.SayHello("Andri")
	//helper.sayGoodbye("Hello") //akan muncul error, karena nama functionnya kecil, tdk dapat diimport

	fmt.Println(helper.Application) //->> bisa diakses
	//fmt.Println(helper.version) //->> tidak bisa diakses
}
