package main

import "fmt"

func main() {
	var name string

	name = "Andri Sutanto"
	fmt.Println(name)

	name = "Andri Tan"
	fmt.Println(name)

	var name2 = "Andri Sutanto Tan"
	fmt.Println(name2)

	var age = 33
	fmt.Println(age)

	name3 := "Andri Sutanto Chen"
	fmt.Println(name3)

	var (
		firstName = "Andri First"
		lastName  = "Andri Last"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

}

//variable adalah tempat untuk menyimpan data
//untuk membuat variable, dimulai dengan kata var, diikuti nama variablenya dan tipe datanya
//contoh var name string

//kalau langsung diinisialisasi datanya, maka tidak perlu menuliskan tipe datanya

//bisa juga tanpa harus menuliskan var saat inisialisasi variablenya, namun
//harus menggunakan :=
//contoh name:="Andri"
//ini hanya berlaku untuk deklarasi awal saja,
//untuk mengubah data selanjutnya, cukup = saja

// MULTIPLE VARIABLE
// ini bisa deklarasi variable secara banyak sekaligus
