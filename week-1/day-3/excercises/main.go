package main

import "fmt"

type Employee struct {
	name string
	age int
	division string
}

func main() {
	employee1 := Employee{
		name: "Airell",
		age: 23,
		division: "Curicullum developer",
	}
	
	fmt.Println(employee1)
	employee2 := &employee1
	employee2.name = "daniel"
	fmt.Println(employee1)
}