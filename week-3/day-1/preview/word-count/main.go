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

	if len(os.Args) < 2 {
		fmt.Println("Missing command line argument")
		os.Exit(1)
	}

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
		wg.Add(1)
		go updateMap(sliceOfWords, &result, wg, mutex)
	}
	return result
}

func updateMap(sliceOfWords []string, result *map[string]int, wg *sync.WaitGroup, mutex *sync.Mutex){
	defer wg.Done()
	for _, word := range sliceOfWords {
		mutex.Lock()
		(*result)[word]++
		mutex.Unlock()
	}
}

func displayResult(result map[string]int) {
	fmt.Println("Word count:")
	for key, value := range result {
		fmt.Printf("%v\t: %v\n", key, value)
	}
}
