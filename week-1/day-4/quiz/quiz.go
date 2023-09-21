package main

import (
	"fmt"
	"math"
	"strings"
)

type Book struct {
	Name string
	Author string
	Price float64
}

func main() {
	// Quiz1()
	// Quiz2()
	// Quiz3()
	// Quiz4()
	// Quiz5()
	// Quiz6()
	// Quiz7()
	// Quiz8()
	// Quiz9()
	// Quiz10()
	// fmt.Println(Quiz11())
	// fmt.Println(Quiz12())
	// fmt.Println(Quiz13(10, 11, 12, 13, 14, 15))
	// fmt.Println(Quiz14(10, 11, 12, 15, 14, 1, 4, 9))

	// Quiz 15
	// buku := []Book{
	// 	{
	// 		Name: "Introduction to Go",
	// 		Author: "John Doe",
	// 		Price: 100000,
	// 	},{
	// 		Name: "Art",
	// 		Author: "Jane",
	// 		Price: 50000,
	// 	},{
	// 		Name: "History",
	// 		Author: "Dan",
	// 		Price: 50000,
	// 	},
	// }
	// fmt.Println(totalHarga(buku))

	// Quiz 16
	// fmt.Printf("%2.f\n", hitungBonus("A", 10000000))
	// fmt.Printf("%2.f\n", hitungBonus("B", 10000000))
	// fmt.Printf("%2.f\n", hitungBonus("C", 10000000))
	// fmt.Printf("%2.f\n", hitungBonus("D", 10000000))

	// Quiz17()
	// Quiz18()
	
	// Quiz 19
	// fmt.Println(buildSentence("Obama", "suka","makan", "nasi", "goreng", "dan", "bakso"))

	// Quiz 20
	// nilaiSiswa := []int{85, 60, 78, 92, 45, 73}
	// evaluasiKinerjaSiswa(nilaiSiswa)
}

func Quiz1() {
	nama := "John Doe"
	umur := 22
	fmt.Println(nama, umur)
}

func Quiz2() {
	stock := map[string]float64{
		"Pepaya": 36500,
		"Apel": 15500,
		"Melon": 20500,
	}
	fmt.Println(stock)
}

func Quiz3() {
	var number int
	fmt.Print("Masukkan angka: ")
	fmt.Scanf("%v\n", &number)
	
	if number%2 == 0{
		fmt.Println("Bilangan genap")
	} else {
		fmt.Println("Bilangan ganjil")
	}
}

func Quiz4() {
	var bulan int
	var s string
	fmt.Print("Masukkan angka 1-12: ")
	fmt.Scanf("%v\n", &bulan)

	switch bulan {
	case 1:
		s = "Januari"
	case 2:
		s = "Februari"
	case 3:
		s = "Maret"
	case 4:
		s = "April"
	case 5:
		s = "Mei"
	case 6:
		s = "Juni"
	case 7:
		s = "Juli"
	case 8:
		s = "Agustus"
	case 9:
		s = "September"
	case 10:
		s = "Oktober"
	case 11:
		s = "November"
	case 12:
		s = "Desember"
	default:
		s = "Out of range"
	}
	fmt.Println(s)
}

func Quiz5() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func Quiz6() {
	var number int
	fmt.Print("Masukkan angka: ")
	fmt.Scanf("%v\n", &number)

	for i := 0; i < number; i++ {
		fmt.Printf("%v ",math.Pow(2, float64(i)))
	}
	fmt.Println()
}

func Quiz7() {
	for i := 1; i <= 20; i++ {
		if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func Quiz8() {
	var N int
	fmt.Print("Masukkan angka: ")
	fmt.Scanf("%v\n", &N)

	for i := 1; i <= N; i++ {
		if i%3 == 0 && i%5 == 0{
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func Quiz9() {
	var N int
	fmt.Print("Masukkan angka: ")
	fmt.Scanf("%v\n", &N)

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if i == 0 || i == N-1 || j == 0 || j == N-1 {
				fmt.Print("*")
			} else if j > 0 && j < N-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func Quiz10() {
	rows := 5
	coeff := 1

	for i := 0; i < rows; i++ {
		for space := 1; space < rows - i; space++ {
			fmt.Print("  ")
		}

		for j := 0; j <= i; j++ {
			if j == 0 || i == 0 {
				coeff = 1
			} else {
				coeff = coeff * (i - j + 1) / j
			}
			fmt.Printf("%4d", coeff)
		}
		fmt.Println()
	}
}

func Quiz11() int {
	var n1, n2 int
	fmt.Print("Masukkan angka 1: ")
	fmt.Scanf("%v\n", &n1)
	fmt.Print("Masukkan angka 2: ")
	fmt.Scanf("%v\n", &n2)
	return n1 + n2
}

func Quiz12() int {
	var n1, n2 int
	fmt.Print("Masukkan angka 1: ")
	fmt.Scanf("%v\n", &n1)
	fmt.Print("Masukkan angka 2: ")
	fmt.Scanf("%v\n", &n2)
	return n1 + n2
}

func Quiz13(nums ...int) int {
	var temp int
	for _, num := range nums {
		temp += num
	}
	return temp
}

func Quiz14(nums ...int) int {
	var temp int
	for _, num := range nums {
		if num > temp {
			temp = num
		}
	}
	return temp
}

func totalHarga(buku []Book) float64 {
	var temp float64
	for _, i := range buku{
		temp += i.Price
	}
	return temp
}

func hitungBonus(performa string, gaji float64) float64 {
	var bonus float64
	switch performa{
	case "A":
		bonus = 0.2 * gaji
	case "B":
		bonus = 0.1 * gaji
	case "C":
		bonus = 0.05 * gaji
	case "D":
		bonus = 0
	default:
		fmt.Println("Input tidak valid")
		return 0
	}
	return bonus
}

func Quiz17() {
	N := 100

	for i := 1; i <= N; i++ {
		if i%3 == 0 && i%5 == 0{
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func Quiz18() {
	n := 5
	mid := n / 2
	for i := 0; i < n; i++ {
		space := mid - i
		if space < 0 {
			space *= -1
		}
		star := 2 * (mid - space) + 1
		fmt.Print(strings.Repeat(" ", space))
		fmt.Print(strings.Repeat("*", star))
		fmt.Println()
	}
}

func buildSentence(word ...string) (string, int) {
	sentence := strings.Join(word, " ")
	length := len(word)
	return sentence, length
}

func evaluasiKinerjaSiswa(nilaiSiswa []int) {
	for i, n := range nilaiSiswa {
		urutan := i + 1
		var predikat string

		if n < 50 {
			predikat = "D"
		} else if n >= 50 && n <= 69 {
			predikat = "C"
		} else if n >= 70 && n <= 84 {
			predikat = "B"
		} else {
			predikat = "A"
		}
		fmt.Printf("Siswa %v mendapatkan predikat %v\n", urutan, predikat)
	}
}