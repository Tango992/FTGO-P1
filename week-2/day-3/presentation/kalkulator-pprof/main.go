package main

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"slices"
)

func main() {
	defer catchErr()
	f, err := os.Create("cpu.pprof")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	f1, err := os.Create("mem.pprof")
	if err != nil {
		panic(err)
	}
	defer f1.Close()
	runtime.GC()
	pprof.WriteHeapProfile(f1)
	
	option, err := displayOption()
	if err != nil {
		panic(err.Error())
	}

	i, j, err := getInt()
	if err != nil {
		panic(err.Error())
	}

	switch option{
	case "a":
		penjumlahan(i, j)
	case "b":
		pengurangan(i, j)
	case "c":
		perkalian(i, j)
	case "d":
		if err := pembagian(i, j); err != nil {
			panic(err.Error())
		}
	}
}

func displayOption() (string, error) {
	fmt.Println("Pilih operasi matematika:")
	fmt.Println("Penjumlahan \t(+) a")
	fmt.Println("Pengurangan \t(-) b")
	fmt.Println("Perkalian \t(*) c")
	fmt.Println("Pembagian \t(/) d")

	var option string
	if _, err := fmt.Scanln(&option); err != nil {
		return "", err
	}

	options := []string{"a", "b", "c", "d"}
	if slices.Contains(options, option) {
		return option, nil
	}
	return "", errors.New("invalid input")
}

func getInt() (int, int, error){
	fmt.Print("Masukkan angka pertama: ")
	var num1 int
	if _, err := fmt.Scanln(&num1); err != nil {
		return 0, 0, err
	}
	fmt.Print("Masukkan angka kedua: ")
	var num2 int
	if _, err := fmt.Scanln(&num2); err != nil {
		return 0, 0, err
	}
	return num1, num2, nil
}

func penjumlahan(i int, j int) {
	fmt.Printf("Hasil penjumlahan %v dan %v adalah %v\n", i, j, i+j)
}

func pengurangan(i int, j int) {
	fmt.Printf("Hasil pengurangan %v dan %v adalah %v\n", i, j, i-j)
}

func perkalian(i int, j int) {
	fmt.Printf("Hasil perkalian %v dan %v adalah %v\n", i, j, i*j)
}

func pembagian(i int, j int) error {
	if j == 0 {
		return errors.New("cannot divide with zero")
	}
	fmt.Printf("Hasil pembagian %v dan %v adalah %.2f\n", i, j, float32(i)/ float32(j))
	return nil
}


func catchErr() {
	if r := recover(); r != nil {
		fmt.Println("Error:", r)
	}
}