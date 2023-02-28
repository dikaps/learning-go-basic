package main

import "fmt"

func judul(title string) {
	const line string = "-----"
	fmt.Println(line, title, line)
}

// Operator

func main() {
	judul("Perulangan cara standar")

	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	judul("Perulangan gaya while")
	i := 0
	for i < 5 {
		fmt.Println("Angka", i)
		i++
	}

	judul("for tanpa argumen")
	i = 0
	for {
		fmt.Println("Angka", i)
		i++

		if i == 3 {
			break
		}
	}

	judul("Penggunaan break dan continue")
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}

	judul("Perluangan bersarang")
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}

	judul("Label dalam Perulangan")
outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}
}
