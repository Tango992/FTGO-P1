package main

import "belajar-go/helpers"

func main() {
	helpers.Greet()
	person := helpers.Person{
		Name: "Daniel",
		Age: 21,
		Address: "Bekasi",
	}
	person.Invokegreet()
}