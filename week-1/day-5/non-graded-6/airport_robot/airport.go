package main

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(s string) string
}

func SayHello(name string, greet Greeter) {
	fmt.Printf("I can speak %v: %v\n", greet.LanguageName(), greet.Greet(name))
}

type Germany struct{}

func (s Germany) LanguageName() string {
	return "German"
}

func (s Germany) Greet(name string) string {
	return fmt.Sprintf("Hallo %v!", name)
}

type Italian struct{}

func (s Italian) LanguageName() string {
	return "Italian"
}

func (s Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %v!", name)
}

type Portuguese struct{}

func (s Portuguese) LanguageName() string {
	return "Portugese"
}

func (s Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %v!", name)
}

func main() {
	p1 := Germany{}
	SayHello("Dietrich", p1)

	p2 := Italian{}
	SayHello("Fabrizio", p2)

	p3 := Portuguese{}
	SayHello("Ronaldo", p3)
}