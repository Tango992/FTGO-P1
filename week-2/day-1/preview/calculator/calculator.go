package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Missing command line argument!")
		os.Exit(1)
	}
	calculator(os.Args[1], os.Args[2], os.Args[3])
}

func calculator(operation string, value1 string, value2 string) {
	input1, err1 := strconv.Atoi(value1)
	input2, err2 := strconv.Atoi(value2)
	if err1 != nil || err2 != nil {
		fmt.Println("Input is not an integer!")
		os.Exit(1)
	}
	i := float64(input1)
	j := float64(input2)

	switch operation {
	case "add":
		fmt.Print("Result: ")
		add(i, j)
	case "sub":
		fmt.Print("Result: ")
		sub(i, j)
	case "mul":
		fmt.Print("Result: ")
		mul(i, j)
	case "div":
		fmt.Print("Result: ")
		div(i, j)
	default:
		fmt.Println("Invalid operation!")
		os.Exit(1)
	}
}

func add(i float64, j float64) {
	fmt.Printf("%.2f\n", i + j)
}

func sub(i float64, j float64) {
	fmt.Printf("%.2f\n", i - j)
}

func mul(i float64, j float64) {
	fmt.Printf("%.2f\n", i * j)
}

func div(i float64, j float64) {
	fmt.Printf("%.2f\n", i / j)
}