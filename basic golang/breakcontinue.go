package main

import "fmt"

func main() {
	// contoh break,
	// kalau break, kode program dibawahnya tidak akan di execute

	for i := 0; i < 10; i++ {

		if i == 5 {
			break
			// begitu di break, dia akan berhenti di perulangan ke 5
		}

		fmt.Println("Perulangan ke", i)
	}

	//kalau continue, dia akan skip statement dibawahnya, dan melanjutkan loopingnya

	for j := 0; j < 10; j++ {
		if j%2 == 0 {
			continue
		}

		fmt.Println("Perulangan ke", j)
	}
}
