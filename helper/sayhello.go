package helper

import "fmt"

var version = 1
var Application = "Belajar Golang"

func SayHello(name string) {
	fmt.Println("Hello", name)
}

/**
perhatikan untuk penamaan functionnya
SayHello menggunakan huruf besar diawal
sedangkan sayGoodbye menggunakan huruf kecil diawal
ini menyebabkan sayGoodbye tidak dapat diakses dariluar
*/

func sayGoodbye(name string) {
	fmt.Println("Goodbye", name)
}
