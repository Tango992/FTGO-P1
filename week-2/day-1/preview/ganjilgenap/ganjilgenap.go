package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	i := getInt()
	isEven(i)
	evenNumbers(i)
	isPrime(i)
}

func getInt() int {
	for {
		var i int
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Please enter a number: ")
		scanner.Scan()
		_, err := fmt.Sscanf(scanner.Text(), "%d", &i)
		if err != nil {
			fmt.Println("Input is not an integer!")
			continue
		}
		return i
	}
}

func isEven(i int) {
	if i % 2 == 0 {
		fmt.Println("The number is even.")
		return
	}
	fmt.Println("The number is odd.")
}

func evenNumbers(num int) {
	if num < 2 {
		fmt.Println("There is no even numbers up to your input.")
		return
	}
	fmt.Println("Even numbers up to your input are:")
	for i := 2; i <= num; i++ {
		if i % 2 == 0 {
			fmt.Println(i)
		}
	}
}

func isPrime(num int) {
	checker := false
	for i := 2; i < num; i++ {
		if num % i == 0 || num == 1 {
			checker = true
			break
		}
	}

	if checker {
		fmt.Println("The number is not a prime number.")
	} else {
		fmt.Println("The number is a prime number.")
	}
}