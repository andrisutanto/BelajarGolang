package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpAndri NoKTP = "1235454654654897465456"
	fmt.Println(noKtpAndri)

	var statusMarried Married = false
	fmt.Println("Married Status= ", statusMarried)
}
