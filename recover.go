package main

import "fmt"

// recover ini untuk menangkap data error / panic

// sebelumnya, jika menggunakan panic, saat program dijalankan dan menemukan error
// maka akan dihentikan
// jika kita menggunakan recover, saat menemukan error
// maka program akan tetap dijalankan

// recover

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Terjadi error dengan message: ", message)
	}

	fmt.Println("APLIKASI SELESAI")
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("APLIKASI BERJALAN")
}

func main() {
	runApp(true)
	fmt.Println("Success")
}
