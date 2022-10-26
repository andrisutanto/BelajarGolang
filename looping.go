package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke: ", counter)
		counter++
	}

	// atau bisa juga seperti ini:
	// counter2 := 1 -> ini adalah init statement
	// counter2 <= 10 -> kondisi perulangan
	// counter2++ -> post statement
	for counter2 := 1; counter2 <= 10; counter2++ {
		fmt.Println("Perulangan ke: ", counter2)
	}

	// coba digabungkan dengan slice
	slice := []string{"Andri", "Sutanto", "Tan"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// bisa juga menggunakan range
	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}

	// kalau mau cetak valuenya aja begini:
	for _, value := range slice { //variable i, tidak digunakan maka akan muncul error, maka harus diganti _
		//fmt.Println("Index", i, "=", value)
		fmt.Println(value)
	}

	// kalau untuk data MAP begini
	person := make(map[string]string)
	person["name"] = "Andri"
	person["title"] = "Fullstack Developer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
