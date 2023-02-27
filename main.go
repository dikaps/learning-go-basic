package main

import "fmt"

func judul(title string) {
	const line string = "-----"
	fmt.Println(line, title, line)
}

// Operator

func main() {
	judul("Operator Aritmatika")
	var value = (((2+6)%3)*4 - 2) / 3
	fmt.Println(value)

	judul("Operator Perbandingan")
	var isEqual bool = value == 2
	fmt.Printf("nilai %d == %t \n", value, isEqual)

	judul("Operator Logika")
	const left = false
	const right = true

	const leftAndRight = left && right
	fmt.Printf("left && right \t = %t \n", leftAndRight)

	const leftOrRight = left || right
	fmt.Printf("left || right \t = %t \n", leftOrRight)

	const leftReverse = !left
	fmt.Printf("!left \t = %t \n", leftReverse)
}
