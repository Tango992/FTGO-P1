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

func calculator(operation string, v1 string, v2 string) {
	i, err1 := strconv.Atoi(v1)
	j, err2 := strconv.Atoi(v2)
	if err1 != nil || err2 != nil {
		fmt.Println("Input is not an integer!")
		os.Exit(1)
	}

	switch operation {
	case "add":
		add(i, j)
	case "sub":
		sub(i, j)
	case "mul":
		mul(i, j)
	case "div":
		div(i, j)
	default:
		fmt.Println("Invalid operation!")
		os.Exit(1)
	}
}

func add(i int, j int) {

}

func sub(i int, j int) {

}

func mul(i int, j int) {

}

func div(i int, j int) {

}