package main

import "fmt"

// soal 1
/*
func main() {

	for i := 1; i <= 20; i++ {

		if i%2 == 0 {
			//fmt.Println("Genap", i)
			fmt.Println("Berkualitas", i)
		} else if i%3 == 0 {
			fmt.Println("I Love Coding ", i)
		} else {
			//fmt.Println("Ganjil", i)
			fmt.Println("Santai", i)
		}

	}


}

*/

// soal2
/*
func main() {
	var hasil = ""
	for i := 1; i <= 7; i++ {

		for h := 1; h <= i; h++ {
			//fmt.Println("#", h)

			hasil += "#"

		}
		hasil += "\n"
	}
	fmt.Println(hasil)
}
*/

// soal 3

/*
func main() {
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	fmt.Println(kalimat[2], kalimat[3], kalimat[4], kalimat[5], kalimat[6])
}
*/

// soal 4

/*
func main() {
	var sayuran = []string{"Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun"}
	//fmt.Println(len(sayuran))
	for i := 0; i <= len(sayuran)-1; i++ {

		fmt.Println(i+1, sayuran[i])
	}
}
*/

// soal 5
func main() {
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}

	//fmt.Println(satuan["lebar"])
	for i := 0; i < 1; i++ {

		fmt.Println("Panjang", satuan["panjang"])
		fmt.Println("Lebar", satuan["lebar"])
		fmt.Println("Tinggi", satuan["tinggi"])
		fmt.Println("Volume Balok", satuan["panjang"]*satuan["lebar"]*satuan["tinggi"])
	}
}
