package database

var connection string

// ini adalah contoh inisialisasi function
// jadi akan menjalankan func ini sebelum menjalankan kode yang lain
func init() {
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
