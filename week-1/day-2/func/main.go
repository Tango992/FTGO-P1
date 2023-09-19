package main

import "fmt"

func main() {
	logPesan("INF:", "Server dijalankan", "Database berhasil terkoneksi", "Success")
}

func tambah(a int, b int) int {
	return a + b
}

// fungsi dengan parameter
func operasiMtk(a int, b int, operasi func(int, int)int)int{
	return operasi(a,b)
}

// fungsi variadic
func logPesan(prefix string, pesan ...string) {
	for _, p := range pesan {
		fmt.Println(prefix, p)
	}
}