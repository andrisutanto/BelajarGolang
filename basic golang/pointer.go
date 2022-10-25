package main

import "fmt"

// secara default, di golang semua variable di passing by value, bukan by reference
// jadi saat kita kirim variable ke dalam function
// yang dikirim sebenarnya aalah duplikasinya

// dengan pointer, kita bisa membuat kemampuan pass by reference di Golang

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Surabaya", "Jawa Timur", "Indonesia"}
	//address2 := address1 // -> ini akan mengcopy value address 1 ke address 2

	//kalau mau datanya sama, harus menggunakan pointer, caranya:
	address2 := &address1

	address2.City = "Denpasar" // -> saat ada perubahan FIELD, tidak akan memiliki nilai yang sama, jika tidak pointer
	//kalau pointer, data address1 juga akan berubah

	//kalau di pointernya mengarah ke Address (assign), dan merubah VARIABLE, maka berbeda
	//tidak akan reference ke address1
	//perhatikan &Address -> bukan &address1
	//address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}

	//kemudian bagaimana jika ingin address1 juga tetap berubah?
	//harus menggunakan operator *, caranya:
	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"} //tidak pakai & lagi

	//kita coba buat address 1 lagi
	address3 := &address1
	//maka akan berubah mengikuti address2
	//jadi siapapun yang yang refer ke Address, dan menggunakan *, maka akan diubah semua

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	//kita juga bisa membuat pointer baru kosong dengan cara:
	var address4 *Address = new(Address)
	address4.City = "Soerabaya"
	fmt.Println(address4)
}
