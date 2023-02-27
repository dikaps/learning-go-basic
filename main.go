package main

import "fmt"

func main() {
	var firstName string = "Andika Permana"

	var lastName, _ string
	lastName = "Sidiq"

	fmt.Printf("Halo %s %s!\n", firstName, lastName)

	/*
		deklrasi variabel tanpa tipe data ( type inference )
		type inference merupakan sebuah teknik pendeklrasian variable
		 tanpa mendefinisikan keyword var berikut dengan tipe datanya
	*/
	tanpaTipeData := "ini tuh string"
	fmt.Printf("deklrasi variabel tanpa keyword var %s\n", tanpaTipeData)

	// deklarasi multiple variabel
	one, two := 1, "abcd"

	fmt.Println("deklrasi multiple variabel", one, two)

	// Variabel Underscore digunakan untuk menampung nilai yang tidak akan dipakai

	_ = "WOw"
	_ = "eh"

	// deklarasi variabel dengan keyword new
	name := new(string)
	fmt.Println(name)
	fmt.Println(*name)

}
