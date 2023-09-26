package main

import "fmt"

func main() {
	c := make(chan string)

	go introduce("John", c)
	go introduce("Jane", c)
	go introduce("Foo", c)

	fmt.Println(<- c)
	fmt.Println(<- c)
	fmt.Println(<- c)
	close(c)
}

func introduce(student string, c chan string) {
	c <- fmt.Sprintf("Hi, my name is %v", student)
}