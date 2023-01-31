package main

import (
	"fmt"
	"strings"
)

// soal 1
/*
func main() {
	//Bootcamp Digital Skill Sanbercode Golang
	var pertama string = "Bootcamp"

	var kedua string
	kedua = "Digital"

	var ketiga string
	ketiga = "Skill"

	var keempat string
	keempat = "Sanbercode"

	var kelima string
	kelima = "Golang"

	fmt.Printf("halo %s %s %s %s %s!\n", pertama, kedua, ketiga, keempat, kelima)
}

*/

//soal 2
/*
func main() {
	//var firstName string = "john"
	name := "Hallo Golang"

	fmt.Printf("%s!\n", name)

}
*/

//soal3

/*
func main() {
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"

	fmt.Printf("halo %s %s %s %s!\n", kataPertama, kataKedua, kataKetiga, kataKeempat)
}
*/

// soal 4
/*
func main() {
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"
	// string to int
	a, err := strconv.Atoi(angkaPertama)
	b, err := strconv.Atoi(angkaKedua)
	c, err := strconv.Atoi(angkaKetiga)
	d, err := strconv.Atoi(angkaKeempat)

	if err != nil {
		// ... handle error
		panic(err)
	}

	var hasil = (a + b + c + d)

	fmt.Println(hasil)
} */

// soal 5
func main() {
	kalimat := "halo halo bandung"
	angka := 2021

	var replace string = strings.Replace(kalimat, "halo", "hei", 3)

	fmt.Println(replace, angka)
}
