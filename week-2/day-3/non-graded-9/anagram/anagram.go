package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
)

func main() {
	defer catchErr()
	
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter first word: ")
	scanner.Scan()
	word1 := scanner.Text()
	fmt.Print("Enter second word: ")
	scanner.Scan()
	word2 := scanner.Text()

	if err := validateInput(word1, word2); err != nil {
		panic(err)
	}

	splice1 := strings.Split(word1, "")
	splice2 := strings.Split(word2, "")
	var counter int
	for _, char := range splice1 {
		if slices.Contains(splice2, char) {
			counter++
		}
	}
	
	if counter == len(word1) {
		fmt.Printf("%v & %v merupakan anagram\n", word1, word2)
	} else {
		fmt.Printf("%v & %v bukan merupakan anagram\n", word1, word2)
	}
}

func validateInput(word1 string, word2 string) error {
	r, _ := regexp.Compile(`\w`)
	match1 := r.MatchString(word1)
	match2 := r.MatchString(word1)
	if len(word1) == len(word2) && len(word1) < 10 && len(word2) < 10 {
		if match1 && match2 {
			return nil
		}
	}
	return errors.New("input tidak valid")
}

func catchErr() {
	if r := recover(); r != nil {
		fmt.Println("Error:", r)
	}
}