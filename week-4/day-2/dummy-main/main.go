package main

import (
	"fmt"
)

func SumOddNumbers(numbers []int) int {
	var total int
	for _, number := range numbers {
		if number % 2 != 0 {
			total += number
		}	
	}
	return total
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(SumOddNumbers(numbers)) // Harusnya mengembalikan 9 karena 1+3+5=9
}
