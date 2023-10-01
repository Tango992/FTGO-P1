package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Biodata struct {
	Name string
	Age int
	Occupation string
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	datas := CsvToStruct(file)
	for i := range datas {
		updateStruct(&datas[i])
	}
	fmt.Println(datas)
}

func CsvToStruct(file *os.File) []Biodata {
	var biodatas []Biodata
	
	records, err:= csv.NewReader(file).ReadAll()
	if err != nil {
		panic(err)
	}
	for _, value := range records[1:] {
		age, err := strconv.Atoi(value[1])
		if err != nil {
			panic(err.Error())
		}

		biodatas = append(biodatas, Biodata{
			Name: value[0],
			Age: age,
			Occupation: value[2],
		})
	}
	return biodatas
}

func updateStruct(biodata *Biodata) {
	biodata.Name = strings.ToUpper(biodata.Name)
	biodata.Occupation = "Mr." + biodata.Occupation
}