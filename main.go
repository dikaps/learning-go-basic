package main

import "fmt"

func batas() {
	fmt.Println("-----------------------------------")
}

// Konstanta

func main() {
	const firstName = "dIkaps"

	fmt.Print("halo ", firstName, "\n")

	batas()

	// desklarasi multi konstanta
	const (
		square         = "kotak"
		isToday  bool  = false
		numeric  uint8 = 1
		floatNum       = 2.2
		pertama        = "ini string"
		kedua          // ketika pendeklarasian konstanta tanpa nilai, maka nilai nya akan berisi nilai dari variabel yang diatasnya
	)

	fmt.Println(square, isToday, numeric, floatNum, pertama, kedua)
}
