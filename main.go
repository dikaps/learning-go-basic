package main

import "fmt"

func batas() {
	fmt.Println("-----------------------------------")
}

func main() {

	// tipe data numerik non ddesimal
	var positiveNumber uint8 = 82
	var negativeNumber = -1243423644

	fmt.Printf("Bilangan positif %d\n", positiveNumber)
	fmt.Printf("Bilangan negatif %d\n", negativeNumber)

	batas()

	// tipe data numerik desimal
	var decimalNumber = 3.85
	fmt.Printf("bilangan desimal %f \n", decimalNumber)
	fmt.Printf("bilangan desimal %.3f \n", decimalNumber)

	batas()

	// tipe data boolean
	var isExist bool = true
	fmt.Printf("Is Exist? %t \n", isExist)

	batas()

	// tipe data string

	var fullName string = "Andika Permana Sidiq"

	fmt.Printf("Nama saya adalah %s \n", fullName)

	var message = `Halo saya sedang mendeklarasikan variabel menggunakan backticks "hahaahha"
kalo di enter dia akan di print enter juga "hehe"
	`
	fmt.Println(message)
}
