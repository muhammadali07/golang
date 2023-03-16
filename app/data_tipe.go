package main

import "fmt"

func main() {
	// nomurik non desimal
	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	// numerik desimal
	var decimalNumber = 2.62

	fmt.Printf("bilangan desimal: %f\n", decimalNumber)   //f menunjukkan 6 digit angkat di belakang koma
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber) //3f menunjukkan 3 digit angkat di belakang koma

	// tipe data boolean
	var exist bool = true
	fmt.Printf("exist? %t \n", exist)

	// tipe data string
	var message1 string = "Halo"
	fmt.Printf("message: %s \n", message1)
	// string menggunakan tanda grave accent/backticks(')
	var message = `Nama saya "John Wick".
		Salam kenal.
		Mari belajar "Golang".`

	fmt.Println(message)
}
