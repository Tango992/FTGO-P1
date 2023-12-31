package main

import (
	"fmt"
)


func main() {
	variable1()
	variable2()
	cli()
	array1()
	array2()
	array3()
}


func variable1() {
	var myNum int32 = 50
	fmt.Printf("myNum: %v\n", myNum)
	var myNum2 float32 = 51
	fmt.Printf("myNum2: %v\n", myNum2)
	myNumStr := "50"
	fmt.Printf("myNumStr: %v\n\n", myNumStr)
}


func variable2() {
	var x, y int32 = 5, 10
	z := x + y
	fmt.Printf("z: %v\n\n", z)
}


func cli() {
	var name string
	fmt.Print("Masukkan nama: ")
	fmt.Scanln(&name)
	fmt.Printf("Hello, %v!\n\n", name)
}


func array1() {
	people := []string{"Walt", "Jesse", "Skyler", "Saul"}
	fmt.Println(people)
	fmt.Printf("Panjang slice people: %v\n", len(people))
	people = append(people, "Hank", "Marie")
	fmt.Println(people)
	fmt.Printf("Panjang slice people: %v\n\n", len(people))
}

// With user input
func array2() {
	arr := [3]map[string]string{}
	
	for i := range arr {
		var name, gender string
		fmt.Print("Name\t: ")
		fmt.Scanln(&name)
		fmt.Print("Gender\t: ")
		fmt.Scanln(&gender)
		fmt.Println()

		arr[i] = map[string]string{"name": name, "gender": gender}
	}
	fmt.Printf("Before\t: %v\n", arr)

	for i := range arr {
		if arr[i]["gender"] == "M" {
			arr[i]["name"] = "Mr. " + arr[i]["name"]
		} else if arr[i]["gender"] == "F" {
			arr[i]["name"] = "Mr. " + arr[i]["name"]
		}
	}
	fmt.Printf("After\t: %v\n", arr)
}

// Alternative without user input
func array3() {
	arr := [3]map[string]string{}

	char1 := map[string]string {
		"name": "Hank",
		"gender": "M",
	}
	char2 := map[string]string {
		"name": "Heisenberg",
		"gender": "M",
	}
	char3 := map[string]string {
		"name": "Skyler",
		"gender": "M",
	}
	
	arr[0] = char1
	arr[1] = char2
	arr[2] = char3
	fmt.Printf("Before\t: %v\n", arr)

	for i := range arr {
		if arr[i]["gender"] == "M" {
			arr[i]["name"] = "Mr. " + arr[i]["name"]
		} else if arr[i]["gender"] == "F" {
			arr[i]["name"] = "Mr. " + arr[i]["name"]
		}
	}
	fmt.Printf("After\t: %v\n", arr)
}
