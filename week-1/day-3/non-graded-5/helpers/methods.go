package helpers

import (
	"fmt"
	"math/rand"
)

func (p *Person) AddYear() {
	p.Age++
	if p.Age >= 50 {
		p.Job = "Retired"
	}
}

func (p Person) GetInfo() {
	fmt.Printf("Halo %v, usia %v tahun, pekerjaan %v\n", p.Name, p.Age, p.Job)
}

func (p Hero) CountDamage() int {
	damages := p.BaseAttack + p.Weapons.Attack
	if rand.Intn(2) == 1 {
		damages += p.CriticalDamage
	}
	return damages
}

func (receiver *Hero) IsAttackedBy(attacker Hero) {
	totalDmg := attacker.CountDamage()
	if totalDmg >= receiver.Defence {
		receiver.HealthPoint -= (totalDmg - receiver.Defence)
	}
}

