package main

import "fmt"

// panic function digunakan utk menghentikan program
// panic function biasanya dipanggil ketika terjadi error, saat program berjalan
// saat panic function, program akan terhenti, namun defer akan tetap dieksekusi

func endApp() {
	fmt.Println("End App")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APP ERROR")
	}
	fmt.Println("App running")
}

func main() {
	runApp(true) //ganti status true atau false utk coba panic
}
