package helpers

import "fmt"

func (p Person) Invokegreet() {
	greet()
	fmt.Printf("Hallo, %v!\n", p.Name)
}