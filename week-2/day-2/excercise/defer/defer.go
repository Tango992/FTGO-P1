package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lineCount, wordCount int
	for scanner.Scan(){
		lineCount++
		line := scanner.Text()
		words := strings.Fields(line)
		wordCount += len(words)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Number of lines: %v\n", lineCount)
	fmt.Printf("Number of words: %v\n", wordCount)
}