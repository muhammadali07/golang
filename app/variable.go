package main

import "fmt"

func main() {
	// variable beserta tipe data
	var firstName string = "john"

	// menggunakan keyword var

	var lastName string
	lastName = "wick"

	// cara mencetak string
	fmt.Printf("halo john wick!\n")
	fmt.Printf("halo %s %s!\n", firstName, lastName)
	fmt.Println("halo", firstName, lastName+"!")

	// varible tanpa deklarasi tipe data
	lastName_tanp_tip_dt := "wick"
	fmt.Printf("halo %s %s!\n", firstName, lastName_tanp_tip_dt)

	// variable multi fungsi
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"
	fmt.Printf("%s, %s, %s", first, second, third)

	// deklarasi variable menggunakan new
	name := new(string)

	fmt.Println(name)
	fmt.Println(*name)

}
