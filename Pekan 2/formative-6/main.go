package main

import "fmt"

// soal 1

/*
func main() {
	var luasLigkaran float64
	var kelilingLingkaran float64

	hasilluasLigkaran(&luasLigkaran, 7)
	hasilkelilingLingkaran(&kelilingLingkaran, 14)

	fmt.Println(luasLigkaran)
	fmt.Println(kelilingLingkaran)

}

func hasilluasLigkaran(luasLigkaran *float64, jari_jari float64) {
	// 22/7 * jari * jari

	*luasLigkaran = 22 * jari_jari

}

func hasilkelilingLingkaran(kelilingLingkaran *float64, jari_jari float64) {

	// 22/7 * jari * 2

	*kelilingLingkaran = (22 / 7) * jari_jari * 2
}
*/

// soal2
/*
func main() {

	var sentence string
	introduce(&sentence, "John", "laki-laki", "penulis", "30")

	fmt.Println(sentence) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"
	introduce(&sentence, "Sarah", "perempuan", "model", "28")

	fmt.Println(sentence) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

}

func introduce(sentence2 *string, nama string, jeniskelamin string, pekerjaan string, usia string) {

	*sentence2 = nama + " adalah " + jeniskelamin + " seorang " + pekerjaan + " yang berusia " + usia
}
*/

//soal 3
/*
func main() {
	var buah = [7]string{}

	tampilBuah(&buah)

	//fmt.Println(buah[0])

	for i, s := range buah {
		fmt.Println(i+1, s)
	}

}

func tampilBuah(bu *[7]string) {

	// cara vertikal
	(*bu)[0] = "Jeruk"
	(*bu)[1] = "Semangka"
	(*bu)[2] = "Mangga"
	(*bu)[3] = "Strawberry"
	(*bu)[4] = "Durian"
	(*bu)[5] = "Manggis"
	(*bu)[6] = "Alpukat"

}
*/

// soal4

func main() {
	var dataFilm = map[string]string{}

	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	for key, val := range dataFilm {
		fmt.Println(key, "=>", val)
	}

}

func tambahDataFilm(name string, waktu string, action string, tahun string, dataFilm *map[string]string) {
	sammy := map[string]string{
		"title":    name,
		"duration": waktu,
		"genre":    action,
		"tahun":    tahun}
	fmt.Println(sammy)

	*dataFilm = sammy
}
