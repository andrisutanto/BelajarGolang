package main

import "fmt"

func main() {

	// [...] -> ini artinya kita tidak tahu kapasitasnya
	// [8] -> ini artinya kapasitasnya 8
	names := [...]string{
		"Andri",
		"Andre",
		"Andrea",
		"Andreas",
		"Andrias",
		"Andria",
		"Andrian",
	}
	slice := names[4:6]

	//ini untuk ambil slice per data
	fmt.Println(slice[0])
	fmt.Println(slice[1])

	//cetak fungsi slice
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	//coba ubah isi array, maka slicenya juga keubah
	//names[4] = "Ubah"
	//fmt.Println(slice)

	//coba ubah isi slicenya, maka arraynya juga akan keubah
	slice[0] = "Ganti Nama" // -> ini artinya index 0 slice, atau array ke 4 dari names diubah
	fmt.Println(names)

	//contoh append, menggunakan Slice dengan data array yang sudah ada / existing
	days := [...]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Minggu",
	}

	daySlice1 := days[5:]
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(days) // muncul Senin, Selasa, Rabu, Kamis, Jumat, Sabtu Baru, Minggu Baru

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Ups"

	fmt.Println(daySlice2) // [Ups, Minggu Baru, Libur Baru]
	fmt.Println(days)      //Senin, Selasa, Rabu, Kamis, Jumat, Sabtu Baru, Minggu Baru

	//lebih baik membuat slice baru dengan array dibelakangnya
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Andri"
	newSlice[1] = "Andri"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	//copy slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	// perhatikan saat buat array / slice
	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}

// SLICE
// slice mirip dengan array, tapi dia bisa berubah, tidak seperti array
// slice terdapat 3 data, pointer, length, dan capacity
// pointer: adalah penunjuk daat pertama di array
// length: panjang dari slice
// capacity: adalah kapasitas dari slice, lenth tdk boleh lebih dari capacity
// capacity seperti total data yang boleh ditampung

// cara buat slice
// array[low:high] -> slice dari low sampai index sebelum high
// array[low:] -> slice dari array dimulai dari low, sampai index akhir
// array[:high] -> slice dari array 0 sampai index sblm high
// array[:] -> slice dari array 0 sampai index akhir

// setiap data yang diubah dari array asalnya, akan reflek ke slicenya

// FUNCTION SLICE
// len(slice) -> panjang slice
// cap(slice) -> kapasitas slice
// append(slice, data) -> membuat slice baru, dgn append ke slice sblmnya, jika penuh akan buat array baru
// make([]TypeData, length, capacity) -> membuat slice baru
// copy(destination, source) -> menyalin slice dari source ke destination
