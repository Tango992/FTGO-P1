package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	menu()
}

func looping1() {
	people := []map[string]any{}

	char1 := map[string]any{
		"name": "Hank",
		"age":  50,
		"job":  "Polisi",
	}
	char2 := map[string]any{
		"name": "Heisenberg",
		"age":  52,
		"job":  "Ilmuan",
	}
	char3 := map[string]any{
		"name": "Skyler",
		"age":  48,
		"job":  "Akuntan",
	}

	people = append(people, char1, char2, char3)

	for _, i := range people {
		fmt.Printf("Hi! Perkenalkan, nama saya %v, umur saya %v tahun, dan saya bekerja sebagai %v.\n", i["name"], i["age"], i["job"])
	}
}

func looping2() {
	slices := [][]float64{}
	slice1 := []float64{1, 5, 7, 8, 10, 24, 33}
	slice2 := []float64{1.1, 5.4, 6.7, 9.2, 11.3, 25.2, 33.1}
	slices = append(slices, slice1, slice2)

	for index, slice := range slices {
		var sum float64
		length := len(slice)
		median := slice[length/2]

		for _, j := range slice {
			sum += j
		}

		avg := sum / float64(length)
		fmt.Printf("Slice %v\n", index+1)
		fmt.Printf("jumlah: %v, rata-rata: %.2f, median: %v \n", sum, avg, median)
	}
}

func logic1() {
	var reversed string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan kata: ")
	scanner.Scan()
	word := scanner.Text()

	for _, char := range word {
		reversed = string(char) + reversed
	}
	// Explore library strings

	if word == reversed {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}

func logic2() {
	var xCounter, oCounter int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan kata: ")
	scanner.Scan()
	word := scanner.Text()

	for _, char := range word {
		switch string(char) {
		case "x", "X":
			xCounter++
		case "o", "O":
			oCounter++
		}
	}
	// Explore library strings (strings.Count)

	if xCounter == oCounter {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}

func logic3() {
	slice := []int{1, 2, 2, 3, 4, 4, 5}

	for i := len(slice); i > 0; i-- {
		for j := 1; j < i; j++ {
			if slice[j-1] < slice[j] {
				buffer := slice[j-1]
				slice[j-1] = slice[j]
				slice[j] = buffer
			}
		}
	}
	fmt.Println(slice)
}

func logic4() {
	var rows int
	fmt.Print("Jumlah rows: ")

	_, err := fmt.Scanf("%d\n", &rows)
	if err != nil {
		fmt.Println("Input bukan berupa angka")
		return
	}

	for i := 0; i < rows; i++ {
		fmt.Println("*")
	}
}

func logic5() {
	var rows int
	fmt.Print("Jumlah rows: ")

	_, err := fmt.Scanf("%d\n", &rows)
	if err != nil {
		fmt.Println("Input bukan berupa angka")
		return
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func logic6() {
	var rows int
	fmt.Print("Jumlah rows: ")

	_, err := fmt.Scanf("%d\n", &rows)
	if err != nil {
		fmt.Println("Input bukan berupa angka")
		return
	}

	for i := 0; i < rows; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func logic7() {
	var rows int
	fmt.Print("Jumlah rows: ")

	_, err := fmt.Scanf("%d\n", &rows)
	if err != nil {
		fmt.Println("Input bukan berupa angka")
		return
	}

	for i := rows; i > 0; i-- {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func logic8() {
	rows := 5

	for i := 0; i < rows; i++ {
		for j := 0; j < rows-i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func menu() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("List:")
	fmt.Println("1. Looping 1")
	fmt.Println("2. Looping 2")
	fmt.Println("3. Palindrome")
	fmt.Println("4. XOXO")
	fmt.Println("5. Bubble Sort")
	fmt.Println("6. Asterisk 1")
	fmt.Println("7. Asterisk 2")
	fmt.Println("8. Asterisk 3")
	fmt.Println("9. Astersik 4")
	fmt.Println("10. Astersik 4 (Alternatif)")
	fmt.Print("Input nomor: ")
	scanner.Scan()
	fmt.Println()
	i := scanner.Text()
	switch i {
	case "1":
		looping1()
	case "2":
		looping2()
	case "3":
		logic1()
	case "4":
		logic2()
	case "5":
		logic3()
	case "6":
		logic4()
	case "7":
		logic5()
	case "8":
		logic6()
	case "9":
		logic7()
	case "10":
		logic8()
	default:
		fmt.Println("Input nomor di luar range")
	}
}