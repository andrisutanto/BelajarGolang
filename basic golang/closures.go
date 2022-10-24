package main

import "fmt"

// closure adalah kemampuan sebuah function utk berinteraksi dengan data2 disekitarnya, dalam scope yang sama

func main() {
	counter := 0
	name := "Andri"

	increment := func() {
		fmt.Println("Increment")
		counter++
		//dari dalam function ini, dapat mengakses variable diluarnya,
		//hati2 saat menggunakan clousure, karena dapat mengganti variable diluar scope

		//name = "Budi" // kalau diganti menjadi budi, maka saat dicetak, akan menjadi Budi
		name := "Budi" // caranya adalah dgn re-declare, perhatikan = dan :=
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
