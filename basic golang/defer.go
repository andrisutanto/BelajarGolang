package main

import "fmt"

// defer function adalah function yg bisa dijadwalkan utk dijalankan setelah sebuah function dieksekusi
// defer function akan selalu dijalankan walaupun terjadi error

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	defer logging() //jadi syntax defer ini memberitahu golang utk execute function ini setelah running functionnya
	// posisi defer harus diatas
	fmt.Println("Run Application")
}

func main() {
	runApplication()
}
