package main

import "fmt"

// parameter yang berada di posisi terakhir, memiliki kemampuan utk dijadikan varargs
// variable argument artinya bisa menerima lebih dari 1 inputan (anggap sprti array)
// bedanya dgn array: kalau array, wajib membuat array sblm dikirim ke function
// kalau parameter, kita bisa langsung kirim datanya, dipisahkan tanda koma

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	total := sumAll(10, 10, 10, 10, 10, 10, 10)
	fmt.Println(total)

	//di Golang, selain memasukkan parameternya, bs juga memasukkan parameter dalam bentuk slice
	//contoh

	numbers := []int{10, 20, 30, 40}
	total = sumAll(numbers...) //kasih ... untuk memberitahu kalau itu adalah slice
	fmt.Println(total)
}
