package main

import "fmt"


type Animal interface {
	Sound() string
	Age() int
}

type Cat struct{
	Name string
	Years int
}

type Dog struct{
	Years int
}

func (c Cat) Age() int {
	return c.Years
}

func (d Dog) Age() int {
	return d.Years
}

func (c Cat) Sound() string {
	return fmt.Sprintf("%v says meowwh", c.Name)
}

func (d Dog) Sound() string {
	return "Woof!"
}

func MakeSound(a Animal){
	fmt.Println(a.Sound())
}

func main() {
	// cat := Cat{Name: "Bleki", Years: 4}
	// dog := Dog{Years: 2}

	// Cara 1
	// fmt.Println(cat.Sound())

	// Cara 2
	// arr := []Animal{cat, dog}
	// for _, v := range arr {
	// 	fmt.Println(v.Sound())
	// }

	// Cara 3
	// MakeSound(cat)
	// MakeSound(dog)

	// Type assertion
	// var a Animal = cat
	// value, ok := a.(Cat)
	// if ok {
	// 	fmt.Printf("This is a cat named %v, and it's %v years old\n", value.Name, value.Age())
	// } else {
	// 	fmt.Println("This is not a cat")
	// }

	// Another examples
	cat1 := Cat{Name: "Whiskers"}
	cat2 := Cat{Name: "Felix"}
	dog1 := Dog{}

	// Interface slice
	animals := []Animal{cat1, cat2, dog1}
	for _, animal := range animals {
		fmt.Println(animal.Sound())
	}

	// Interface Map
	animalMap := map[string]Animal{
		"Whiskers": cat1,
		"Felix": cat2,
		"Doggy": dog1,
	}
	
	for name, animal := range animalMap {
		fmt.Printf("%s: %s\n", name, animal.Sound())
	}
}