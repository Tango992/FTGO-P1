package helpers

import "math/rand"

type Weapon struct {
	Attack int
}

type Hero struct {
	Name string
	BaseAttack int
	Defence int
	CriticalDamage int
	HealthPoint int
	Weapons Weapon
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

func Battle(attacker *Hero, receiver *Hero) {
	receiver.IsAttackedBy(*attacker)
}