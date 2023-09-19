package main

import (
	"fmt"
	"strings"
)

func main() {
	s := AlayGen("hello", "world", "zzz")
	fmt.Println(s)

	i := fibbonaci(20)
	fmt.Println(i)
}

func AlayGen(str ...string) string {
	var combined string
	for _, slice := range str {
		combined += strReplace(slice) + " "
	}
	return combined
}

func strReplace(s string) string {
	var str string
	for _, char := range strings.Split(s, "") {
		switch char {
		case "a":
			str += "4"
		case "e":
			str += "3"
		case "i":
			str += "!"
		case "l":
			str += "1"
		case "n":
			str += "N"
		case "s":
			str += "5"
		case "x":
			str += "*"
		default:
			str += char
		}
	}
	return str
}

func fibbonaci(i int) int {
	if i <= 1 {
		return i
	}
	return fibbonaci(i - 1) + fibbonaci(i - 2)
}