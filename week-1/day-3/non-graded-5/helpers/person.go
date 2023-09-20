package helpers

import (
	"fmt"
)

type Person struct {
	Name string
	Age int
	Job string
}


func (p *Person) AddYear() {
	p.Age++
	if p.Age >= 50 {
		p.Job = "Retired"
	}
}

func (p Person) GetInfo() {
	fmt.Printf("Halo %v, usia %v tahun, pekerjaan %v\n", p.Name, p.Age, p.Job)
}

