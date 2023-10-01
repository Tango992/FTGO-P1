package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	result := scanLine(file, &wg, &mutex)
	wg.Wait()
	displayResult(result)
}

func scanLine(file *os.File, wg *sync.WaitGroup, mutex *sync.Mutex) map[string]int {
	scanner := bufio.NewScanner(file)
	result := map[string]int{}

	for scanner.Scan() {
		line := scanner.Text()
		sliceOfWords := strings.Fields(line)
		for _, word := range sliceOfWords {
			wg.Add(1)
			go func(s string){
				defer wg.Done()
				mutex.Lock()
				result[s]++
				mutex.Unlock()
			}(word)
		}
	}
	return result
}

func displayResult(result map[string]int) {
	fmt.Println("Word count:")
	for key, value := range result {
		fmt.Printf("%v\t: %v\n", key, value)
	}
}
