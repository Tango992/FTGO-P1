package main

import (
	"fmt"
	"strconv"
)

func main() {
	var myNum int32 = 50
	fmt.Printf("myNum: %v\n", myNum)
	var myNum2 float32 = 51
	fmt.Printf("myNum2: %v\n", myNum2)
	myNumStr := "50"
	fmt.Printf("myNumStr: %v\n\n", myNumStr)

	var x, y int32 = 5, 10
	z := x + y
	fmt.Printf("z: %v\n\n", z)

	var name string
	fmt.Print("Masukkan nama: ")
	fmt.Scanln(&name)
	fmt.Printf("Hello, %v!\n\n", name)

	people := []string{"Walt", "Jesse", "Skyler", "Saul"}
	fmt.Println(people)
	fmt.Printf("Panjang slice people: %v\n", len(people))
	people = append(people, "Hank", "Marie")
	fmt.Println(people)
	fmt.Printf("Panjang slice people: %v\n\n", len(people))

	people2 := make([][]string, 3)

	for i := range people2 {
		var name, gender string
		fmt.Print("Name\t: ")
		fmt.Scanln(&name)
		fmt.Print("Gender\t: ")
		fmt.Scanln(&gender)
		fmt.Println()
		people2[i] = []string{strconv.Itoa(i), name, gender}
	}
	fmt.Printf("Before\t: %v\n", people2)

	for i := range people2 {
		if people2[i][2] == "M" {
			people2[i][1] = "Mr. " + people2[i][1]
		} else if people2[i][2] == "F" {
			people2[i][1] = "Mrs. " + people2[i][1]
		}
	}
	fmt.Printf("After\t: %v\n\n", people2)
}