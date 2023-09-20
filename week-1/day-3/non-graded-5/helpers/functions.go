package helpers

import "fmt"

func Battle(attacker Hero, receiver Hero) {
	fmt.Printf("%v is being attacked by %v!\n", receiver.Name, attacker.Name)
	fmt.Printf("%v Health Before \t: %v\n", receiver.Name, receiver.HealthPoint)
	receiver.IsAttackedBy(attacker)
	fmt.Printf("%v Health After \t: %v\n", receiver.Name, receiver.HealthPoint)
}