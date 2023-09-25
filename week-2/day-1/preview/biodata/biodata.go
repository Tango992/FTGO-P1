package main

import (
	"fmt"
	"os"
	"strconv"
)

type Biodata struct{
	Nama string
	Alamat string
	Pekerjaan string
	Alasan string
}

var Kelas = []Biodata{
	{
		Nama: "John Doe",
		Alamat: "Jalan Pegangsaan",
		Pekerjaan: "Back End Developer",
		Alasan: "Ingin mengasah ilmu Golang",
	},{
		Nama: "Jane Doe",
		Alamat: "Jalan Sudirman",
		Pekerjaan: "Bankir",
		Alasan: "Ingin switch career",
	},{
		Nama: "Samantha",
		Alamat: "Kebon Jeruk",
		Pekerjaan: "Fresh Graduate IT",
		Alasan: "Ingin menambah Ilmu.",
	},{
		Nama: "Foo",
		Alamat: "Jalan jalan",
		Pekerjaan: "Fresh Graduate Teknik Industri",
		Alasan: "Ingin switch career",
	},
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Missing command line argument!")
		os.Exit(1)
	}

	i, err := strconv.Atoi(os.Args[1])

	if err != nil {
		os.Exit(1)
	}
	showBiodata(i)
	
}

func showBiodata(j int) {
	i := j - 1
	if i < 0 || i > len(Kelas)-1 {
		fmt.Println("Index out of range!")
		os.Exit(1)
	}
	fmt.Printf("Data teman dengan absen %v :\n", j)
	fmt.Printf("Nama\t\t: %v\n", Kelas[i].Nama)
	fmt.Printf("Alamat\t\t: %v\n", Kelas[i].Alamat)
	fmt.Printf("Pekerjaan\t: %v\n", Kelas[i].Pekerjaan)
	fmt.Printf("Alasan\t\t: %v\n", Kelas[i].Alasan)
}