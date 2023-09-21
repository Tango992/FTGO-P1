package main

type Greeter interface{
	LanguageName() string
	Greet(s string) string
}

func SayHello(name string, greet Greeter) {

}

type Germany struct{}

type Italian struct{}

type Portuguese struct{}

func main() {

}