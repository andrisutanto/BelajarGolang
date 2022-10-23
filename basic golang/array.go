package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Andri"
	names[1] = "Sutanto"
	names[2] = "Tan"

	var namaAmbil = names[2]

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	fmt.Println(namaAmbil)

	//bisa juga deklarasi seperti ini:

	var values = [3]int{
		90,
		80,
		50,
	}

	fmt.Println(values)

	//untuk cari panjang array

	fmt.Println(len(names))
	fmt.Println(len(values))
}

// ARRAY
// untuk array, kita harus menentukan jumlah datanya pada saat awal
