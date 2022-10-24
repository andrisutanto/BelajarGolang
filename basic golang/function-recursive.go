// ada beberapa case, akan lebih mudah diselesaikan dgn recursive

package main

import "fmt"

// ini contoh menggunakan loop
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursive(value int) int {
	if value == 1 { //di recursive, harus ada kondisi untuk berhenti
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	loop := factorialLoop(5)
	fmt.Println(loop)
	fmt.Println(5 * 4 * 3 * 2 * 1)
	fmt.Println("-----")
	fmt.Println(factorialRecursive(5))
}
