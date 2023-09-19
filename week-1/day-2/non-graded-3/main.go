package main

import (

)

func main() {
	looping1()
}

func looping1() {
	arr := [3]map[string]any{}

	char1 := map[string]any{
		"name": "Hank",
		"age": 50,
		"job": "Polisi",
	}
	char2 := map[string]any{
		"name": "Heisenberg",
		"age": 52,
		"job": "Ilmuan",
	}
	char3 := map[string]any{
		"name": "Skyler",
		"age": 48,
		"Job": "Akuntan",
	}
	
	arr[0] = char1
	arr[1] = char2
	arr[2] = char3
}