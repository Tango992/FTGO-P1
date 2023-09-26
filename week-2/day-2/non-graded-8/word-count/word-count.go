package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main(){
	var wg sync.WaitGroup
	files := []string{"file1.txt", "file2.txt", "file3.txt"}
	for _, file := range files {
		wg.Add(1)
		go func(file string){
			defer wg.Done()
			fmt.Printf("%v: %d words\n", file, countWords(file))
		}(file)
	}
	wg.Wait()
}

func countWords(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var wordCount int
	for scanner.Scan(){
		line := scanner.Text()
		words := strings.Fields(line)
		wordCount += len(words)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return wordCount
}