package main

import (
	"encoding/csv"
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

	records := CsvToSlice(file, &wg, &mutex)

	outputFile, err := os.Create(os.Args[2])
	if err != nil {
		panic(err.Error())
	}

	writer := csv.NewWriter(outputFile)
	writer.WriteAll(records)

	if err := writer.Error(); err != nil {
		panic(err.Error())
	}
}

func CsvToSlice(file *os.File, wg *sync.WaitGroup, mutex *sync.Mutex) [][]string {
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	for i := range records {
		if i == 0 {
			continue
		}
		wg.Add(1)
		go updateValue(i, &records, wg, mutex)
	}
	wg.Wait()
	return records
}

func updateValue(i int, records *[][]string, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	mutex.Lock()
	(*records)[i][0] = strings.ToUpper((*records)[i][0])
	(*records)[i][2] = "Mr." + (*records)[i][2]
	mutex.Unlock()
}
