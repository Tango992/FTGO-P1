package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan soal 1 / 2: ")
	scanner.Scan()
	i := scanner.Text()
	
	if i == "1" {
		conditional1()
	} else if i == "2" {
		conditional2()
	}
}

func conditional1() {
	var name string
	fmt.Print("Masukkan nama: ")
	fmt.Scanf("%v\n", &name)
	

	i := math.Round(rand.Float64() * 100)
	switch {
	case i > 80:
		fmt.Printf("Selamat %v, anda sangat beruntung\n", name)
	case i <= 80 && i > 60:
		fmt.Printf("Selamat %v, anda beruntung\n", name)
	case i <= 60 && i > 40:
		fmt.Printf("Mohon maaf %v, anda kurang beruntung\n", name)
	default:
		fmt.Printf("Mohon maaf %v, anda sial\n", name)
	}
}

func conditional2() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukkan nama: ")
	scanner.Scan()
	name := strings.TrimSpace(scanner.Text())
	fmt.Print("Masukkan umur: ")
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())

	if age, err := strconv.Atoi(input); err == nil {
		if age < 0 || age > 100 {
			fmt.Printf("Umur tidak valid, %v\n", name)
		} else if age > 18 {
			fmt.Printf("Silakan masuk, %v\n", name)
		} else {
			fmt.Printf("Dilarang masuk, minimal umur 19, %v\n", name)
		}
	} else {
		fmt.Println("Umur bukan tipe data integer")
	}
}