package main

import "fmt"

func main() {
	// operator masih sama dengan di python bentuknya
	var value = (((2+6)%3)*4 - 2) / 3
	var isEqual = (value == 2)

	fmt.Printf("nilai %d (%t) \n", value, isEqual)

	//operator logika
	var left = false
	var right = true

	var leftAndRight = left && right // and
	fmt.Printf("left && right \t(%t) \n", leftAndRight)

	var leftOrRight = left || right // or
	fmt.Printf("left || right \t(%t) \n", leftOrRight)

	var leftReverse = !left // nor atau not or
	fmt.Printf("!left \t\t(%t) \n", leftReverse)
}
