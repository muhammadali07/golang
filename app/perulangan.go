package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	// for dengan argumen hanya kondisi
	var x = 0

	for x < 5 {
		fmt.Println("Angka", x)
		x++
	}

	// for tanpa kondisi
	var i = 0

	for {
		fmt.Println("Angka", i)

		i++
		if i == 5 {
			break
		}
	}

	// for with break-countinue
	for y := 1; y <= 10; y++ {
		if y%2 == 1 {
			continue
		}

		if y > 8 {
			break
		}

		fmt.Println("Angka", y)
	}

	//nested for
	for z := 0; z < 5; z++ {
		for j := z; j < 5; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()
	}

	// label dalam perulangan
outerLoop:
	for v := 0; v < 5; v++ {
		for j := 0; j < 5; j++ {
			if v == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", v, "][", j, "]", "\n")
		}
	}
}
