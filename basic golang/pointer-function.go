package main

import "fmt"

type Address1 struct {
	City, Province, Country string
}

func ChangeAddressToIndonesia(address *Address1) {
	address.Country = "Indonesia"
}

func main() {
	address := Address1{"Subang", "Jawa Barat", ""}
	ChangeAddressToIndonesia(&address)

	fmt.Println(address) // tidak berubah datanya
	//datanya tidak berubah karena yang dikirim adalah copy valuenya aja
}
