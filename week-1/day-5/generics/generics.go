package main

import "fmt"

type UserId int

func Hello[T int | float32 | string](param T){
	fmt.Println("hello", param)
}

func Add[T ~int | float64](a T, b T) T {
	return a + b
}

type ScoreType interface{
	int | float64
}

type Person[T ScoreType] struct {
	Name string
	Scores []T
}


func (p *Person[T]) AddScores(vals ...T){
	for _, val := range vals {
		p.Scores = append(p.Scores, val)
	}
}

func main() {
	// Hello("world")

	// a := UserId(0)
	// b := UserId(1)
	// fmt.Println(Add(a, b))

	john := Person[float64]{
		Name: "John",
		Scores: []float64{
			90, 81.2, 93.2,
		},
	}
	fmt.Println(john.Scores)

	john.AddScores(4, 5, 7.9)
	fmt.Println(john.Scores)
}