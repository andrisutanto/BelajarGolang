package main

import (
	"belajar-golang/database"
	"fmt"
	/**
	semisal kita ingin import "database"
	namun tidak kita gunakan method atau variablenya
	maka akan muncul error "unused import"
	cara untuk ttp mengimport tanpa muncul error,adalah dengan menambahkan _
	BLANK IDENTIFIER, contohnya seperti dibawah
	import _ "belajar-golang/database"
	*/)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
