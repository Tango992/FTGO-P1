package main

import "fmt"


type Cat struct{
	Name string
	Years int
}

type Dog struct{
	Years int
}


func Describe(i interface{}) {
	switch v := i.(type) {
	case Cat:
		fmt.Println("This is a cat named:", v.Name, "and it is", v.Years, "years old.")
	case Dog:
		fmt.Println("This is a dog and it is", v.Years, "years old.")
	default:
		fmt.Println("Unknown type!")
	}
}

func main() {
	cat := Cat{Name: "Whiskers", Years: 3}
	dog := Dog{Years: 5}

	Describe(cat)
	Describe(dog)
}