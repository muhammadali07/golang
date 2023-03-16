package main

import "fmt"

func main() {
	const firstName string = "john" //const adalah keyword untuk mengunci value dari variable menjadi tidak bisa di ubah
	fmt.Print("halo ", firstName, "!\n")

	// teknik inference dengan menghilangkan tipe data
	const lastName = "wick"
	fmt.Print("nice to meet you ", lastName, "!\n")

	// print() vs println()

	fmt.Println("john wick") //menghasilkan baris baru
	fmt.Println("john", "wick")

	fmt.Print("john wick\n") //tidak menghasilkan baris baru
	fmt.Print("john ", "wick\n")
	fmt.Print("john", " ", "wick\n")

	//deklarasi multi konstanta
	const (
		square         = "kotak"
		isToday  bool  = true
		numeric  uint8 = 1
		floatNum       = 2.2
	)

	// multiple konstanta dalam satu baris
	const satu, dua = 1, 2
	const three, four string = "tiga", "empat"
}
