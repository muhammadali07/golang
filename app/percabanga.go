package main

import "fmt"

func main() {
	var point = 8

	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", point)
	}

	//variable temporary if-else
	var pointku = 8840.0

	if percent := pointku / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	// switch case
	var point3 = 6

	switch point3 {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	// case untuk banyak kondisi
	var pointcase = 6

	switch pointcase {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	//{} pada case dan default
	var pointcasedefault = 6

	switch pointcasedefault {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be better!")
		}
	}

	// switch case gaya if-else
	var pointifelse = 6

	switch {
	case pointifelse == 8:
		fmt.Println("perfect")
	case (pointifelse < 8) && (pointifelse > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	//fallthrough adl memaksa switch terus berjalan hingga akhir tanpa menghiraukan valuenya salah atau benar
	var pointfallthrough = 6

	switch {
	case pointfallthrough == 8:
		fmt.Println("perfect")
	case (pointfallthrough < 8) && (pointfallthrough > 3):
		fmt.Println("awesome")
		fallthrough
	case pointfallthrough < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	//if-else nested atau bersarang
	var pointnested = 10

	if pointnested > 7 {
		switch pointnested {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if pointnested == 5 {
			fmt.Println("not bad")
		} else if pointnested == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if pointnested == 0 {
				fmt.Println("try harder!")
			}
		}
	}
}
