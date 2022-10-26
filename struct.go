package main

import "fmt"

// struct ini adalah template gabungan dari beberapa tipe data yang berbeda2 tipenya
// sebenarnya bisa array atau map, tapi di array dan map hanya bisa 1 tipe data yang sama
// sederhananya, struct adalah kumpulan dari field

// struct tidak bisa langsung digunakan
// kita harus membuat data/object dari struct yang telah kita buat

type Customer struct { //biasanya penulisannya upper to lower case
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func (a Customer) sayHuuuu() {
	fmt.Println("Huuuuuu from", a.Name)
}

func main() {

	// berikut contoh cara buat struct
	var andri Customer
	andri.Name = "Andri Sutanto"
	andri.Address = "Indonesia"
	andri.Age = 30

	fmt.Println(andri)

	fmt.Println(andri.Name)
	fmt.Println(andri.Address)
	fmt.Println(andri.Age)

	//selain diatas, kita juga bisa membuat struct dengan cara:
	budi := Customer{
		Name:    "Budi",
		Address: "Singapore",
		Age:     28,
	}

	fmt.Println(budi)

	//atau bisa juga seperti ini:
	joko := Customer{"Joko", "Hindia Belanda", 32}
	fmt.Println(joko)

	andri.sayHi("Mike")
	andri.sayHuuuu()

}
