package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() { //kalau "Man" nya masih "Man" saja, maka yang dicetak hanya EKO, tidak ada perubahan data hasil dari proses function
	//supaya merubah dai main function, kita ubah "Man" ---> "*Man"
	man.Name = "Mr. " + man.Name
	fmt.Println("Di Method", man.Name)
}

func main() {
	eko := Man{"Eko"}
	eko.Married()

	fmt.Println(eko.Name)
}
