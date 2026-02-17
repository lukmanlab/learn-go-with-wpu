package main

import "fmt"

func heal(hp *int)  {
	*hp = *hp + 20
	fmt.Println("Pemain disembuhkan +20!")
}

func attack(hp *int, damage int)  {
	*hp -= damage
	fmt.Printf("Pemain diserang! HP berkurang %d\n", damage)
	if *hp <= 0 {
		fmt.Println("Game over! Pemain kalah!")
	}
}

func main()  {
	hp := 50;
	fmt.Println("HP awal:", hp)

	//heal
	heal(&hp)
	fmt.Println("HP sekarang:", hp)

	//attack1
	attack(&hp, 30)
	fmt.Println("HP sekarang", hp)

	//attack2
	attack(&hp, 50)
	fmt.Println("HP sekarang", hp)
}
