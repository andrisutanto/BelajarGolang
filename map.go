package main

import "fmt"

func main() {
	person := map[string]string{ // --> tipe [key],tipe value
		"name":    "Andri",
		"address": "Surabaya",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	// untuk add data di map, bisa menggunakan:
	person["title"] = "Programmer"
	fmt.Println(person)

	delete(person, "address")
	fmt.Println(person)
}

// MAP
// map adalah tipe data yang berisikan sekumpulan data, namun bisa menentukan jenis tipe datanya
// sederhananya, ini sekumpuluna key - value data
