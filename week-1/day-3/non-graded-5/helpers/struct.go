package helpers

type Person struct {
	Name string
	Age int
	Job string
}

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