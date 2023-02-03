package main

import "fmt"

// soal 1

/*
func main() {
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println("luas persegi panjang", luas)
	fmt.Println("keliling Persegi Panjang", keliling)
	fmt.Println("Volume Balok", volume)
}

func luasPersegiPanjang(panjang, lebar int) int {

	var hasil = panjang * lebar

	return hasil
}

func kelilingPersegiPanjang(panjang, lebar int) int {

	var hasil = 2 * (panjang + lebar)

	return hasil
}

func volumeBalok(panjang, lebar, tinggi int) int {

	var hasil = panjang * lebar * tinggi

	return hasil
}
*/

// soal 2

/*
func main() {

	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"
}

func introduce(nama string, jenis_kelamin string, pekerjaan string, umur string) string {
	hasil := nama + " adalah seorang " + pekerjaan + " yang berusia " + umur + " tahun"

	return hasil
}
*/

// soal 3

/*
func main() {
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}

	var buahFavoritJohn = buahFavorit("John", buah...)

	fmt.Println(buahFavoritJohn)
	// halo nama saya john dan buah favorit saya adalah "semangka", "jeruk", "melon", "pepaya"
}

func buahFavorit(name string, databuah ...string) string {
	hasil := "halo nama saya " + name + " dan buah favorit saya adalah " + databuah[0] + " " + databuah[1] + " " + databuah[2] + " " + databuah[3]

	return hasil
}
*/

// soal 4

func main() {

	var dataFilm = []map[string]string{}
	// buatlah closure function disini

	// var chicken = map[string]int{
	// 	"januari":  50,
	// 	"februari": 40,
	// 	"maret":    34,
	// 	"april":    67,
	// }

	tambahDataFilm("LOTR", "2 jam", "action1", "1999")
	tambahDataFilm("avenger", "2 jam", "action2", "2019")
	tambahDataFilm("spiderman", "2 jam", "action3", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}
}

func tambahDataFilm(data ...string) string {
	var hasil2 = ""
	hasil := data

	for i := 0; i <= len(hasil)-1; i++ {

		hasil2 = "genre " + hasil[2] + " jam " + hasil[1] + " tahun " + hasil[3] + " title " + hasil[0]
	}
	hasil2 += "\n"
	fmt.Println(hasil2)

	return hasil2
}
