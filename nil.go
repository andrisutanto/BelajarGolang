package main

import "fmt"

// nil adalah data kosong
// nil bisa digunakan pada beberapa tipe data, interface, function, map, slice, pointer, channel

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {

	var person map[string]string = nil
	fmt.Println(person)

	// begini cara pakainya nil

	data := NewMap("")
	if data == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(data)
	}
}
