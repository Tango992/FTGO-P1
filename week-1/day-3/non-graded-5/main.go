package main

import (
	"nongraded/helpers"
	"fmt"
)


func main() {
	/*
		Uncommand function di bawah untuk menjalankan program
	*/

	// StructMethod1()
	// StructMethod2()
	// StructHero1()
	StructHero2()
}

func StructMethod1() {
	person := helpers.Person{
		Name: "Bambang",
		Age: 49,
		Job: "Gambler",
	}
	fmt.Println(person)
	person.AddYear()
	fmt.Println(person)
}

func StructMethod2() {
	persons := []helpers.Person{
		{
			Name: "Bambang",
			Age: 20,
			Job: "Gambler",
		},{
			Name: "John",
			Age: 40,
			Job: "Engineer",
		},{
			Name: "Foo",
			Age: 30,
			Job: "Photographer",
		},
	}
	for _, people := range persons {
		people.GetInfo()
	}
}

func StructHero1() {
	hero1 := helpers.Hero{
		Name: "Reyna",
		BaseAttack: 100,
		Defence: 200,
		CriticalDamage: 50,
		HealthPoint: 100,
		Weapons: helpers.Weapon{
			Attack: 30,
		},
	}
	dmg := hero1.CountDamage()
	fmt.Println(dmg)
}

func StructHero2() {
	hero1 := helpers.Hero{
		Name: "Reyna",
		BaseAttack: 100,
		Defence: 200,
		CriticalDamage: 50,
		HealthPoint: 100,
		Weapons: helpers.Weapon{
			Attack: 30,
		},
	}
	hero2 := helpers.Hero{
		Name: "Chamber",
		BaseAttack: 100,
		Defence: 200,
		CriticalDamage: 50,
		HealthPoint: 200,
		Weapons: helpers.Weapon{
			Attack: 150,
		},
	}
	helpers.Battle(hero2, hero1)
}