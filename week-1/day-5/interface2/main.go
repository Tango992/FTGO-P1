package main

import "fmt"


type Cat struct {
	Name string
	Years int
}


func main() {
	data := []interface{}{42, "Hello", true, Cat{Name: "Whiskers", Years: 2}}
	fmt.Println("Empty interface slice:")
	for _, item := range data {
		fmt.Println(item)
	}
	fmt.Println()

	// Empty Interface Map
	mapData := map[interface{}]interface{}{
		"number": 42,
		42: "is the answer",
		true: Cat{Name: "Whiskers"},
	}
	fmt.Println("Empty interface map:")
	for key, value := range mapData {
		fmt.Printf("%v\t: %v\n", key, value)
	}
}