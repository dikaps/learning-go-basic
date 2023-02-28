package main

import "fmt"

func judul(title string) {
	const line string = "-----"
	fmt.Println(line, title, line)
}

// Operator

func main() {
	judul("Seleksi Kondisi")

	var point = 8

	if point == 10 {
		fmt.Println("Lulus dengan nilai sempurna")
	} else if point >= 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d \n", point)
	}

	// temporary variabel pada if - else
	// variabel yang hanya bisa digunakan dalam scope seleksi kondisi saja
	judul("temporary variabel")
	var points = 8840.0

	if percent := points / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect! \n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good \n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad \n", percent, "%")
	}

	// switch case
	// seleksi kondisi yang sifatnya pokus pada satu variabel
	judul("switch case")
	switch point {
	case 8:
		{
			fmt.Println("just say ehlo")
			fmt.Println("perfect")
		}
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	// switch dengan gaya if else
	judul("switch dengan gaya if else")
	point = 6
	switch {
	case point == 8:
		fmt.Println("perfect")
	case point < 8 && point > 3:
		fmt.Println("awesome")
		fallthrough // keyword untuk memaksa melanjutkan ke perintah case yang selanjutnya tanpa memperdulikan case nya benar atau salah
	case point < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	judul("seleksi kondisi bersarang")
	point = 10
	if point > 7 {
		switch point {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if point == 5 {
			fmt.Println("Not bad")
		} else if point == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if point == 0 {
				fmt.Println("try harder!")
			}
		}
	}
}
