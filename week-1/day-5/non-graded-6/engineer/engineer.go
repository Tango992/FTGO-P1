package main

import (
	"fmt"
)

type Employee interface {
    Language() string
    Age() int
}

type Engineer struct {
    Name string
    Aged int
}

func (e Engineer) Age() int {
	return e.Aged
}

func (e Engineer) Language() string {
    return e.Name + " programs in Go"
}

type Developer struct {
    Name string
    Aged int
}

func (e Developer) Age() int {
	return e.Aged
}

func (e Developer) Language() string {
    return e.Name + " programs in JS"
}


func main() {
    // This will throw an error
    var programmers []Employee
    elliot := Engineer{Name: "Elliot", Aged: 40}
    john := Developer{Name: "John", Aged: 30}
    // Engineer does not implement the Employee interface
    // you'll need to implement Age() and Random()
    programmers = append(programmers, elliot, john)

    for _, programmer := range programmers {
        fmt.Printf("%v, aged %v years old.\n", programmer.Language(), programmer.Age())
    }
}
