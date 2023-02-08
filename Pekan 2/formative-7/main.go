package main

type Buah struct {
	nama        string
	warna       string
	ada_bijinya string
	harga       int
}

//soal 1
/*
func main() {
	var buah Buah
	buah.nama = "Nanas"
	buah.warna = "Kuning"
	buah.ada_bijinya = "Tidak"
	buah.harga = 90000

	var buah1 Buah
	buah1.nama = "Jeruk"
	buah1.warna = "Oranye"
	buah1.ada_bijinya = "Ada"
	buah1.harga = 8000

	var buah2 Buah
	buah2.nama = "Semangka"
	buah2.warna = "Hijau & Merah"
	buah2.ada_bijinya = "Ada"
	buah2.harga = 100000

	var buah3 Buah
	buah3.nama = "Pisang"
	buah3.warna = "Kuning"
	buah3.ada_bijinya = "Tidak"
	buah3.harga = 5000

	fmt.Println("nama :", buah.nama)
	fmt.Println("warna :", buah.warna)
	fmt.Println("ada_bijinya :", buah.ada_bijinya)
	fmt.Println("harga :", buah.harga)

	fmt.Println("=====================")

	fmt.Println("nama :", buah1.nama)
	fmt.Println("warna :", buah1.warna)
	fmt.Println("ada_bijinya :", buah1.ada_bijinya)
	fmt.Println("harga :", buah1.harga)

	fmt.Println("=====================")

	fmt.Println("nama :", buah2.nama)
	fmt.Println("warna :", buah2.warna)
	fmt.Println("ada_bijinya :", buah2.ada_bijinya)
	fmt.Println("harga :", buah2.harga)

	fmt.Println("=====================")

	fmt.Println("nama :", buah3.nama)
	fmt.Println("warna :", buah3.warna)
	fmt.Println("ada_bijinya :", buah3.ada_bijinya)
	fmt.Println("harga :", buah3.harga)
}
*/

//soal2

/*
type segitiga struct {
	alas, tinggi int
}

type persegi struct {
	sisi int
}

type persegiPanjang struct {
	panjang, lebar int
}

func main() {
	var segitiga = segitiga{12, 4}
	segitiga.saySegitiga()

	var persegi = persegi{6}
	persegi.sayPersegi()

	var persegiPanjang = persegiPanjang{6, 10}
	persegiPanjang.sayPersegiPanjang()
}

func (s segitiga) saySegitiga() {
	hasil_segitiga := (s.alas * s.tinggi) / 2
	fmt.Println("luas segitiga", hasil_segitiga)
}

func (s persegi) sayPersegi() {
	hasil_persegi := s.sisi * s.sisi
	fmt.Println("luas persegi", hasil_persegi)
}

func (s persegiPanjang) sayPersegiPanjang() {
	hasil_segitiga := (s.panjang * s.lebar) / 2
	fmt.Println("luas persegiPanjang", hasil_segitiga)
}
*/

// BELUM SEMPET NGISI

// soal3

// type phone struct {
// 	name, brand string
// 	year        int
// 	colors      []string
// }

// func main() {

// }

// // soal 4

// func main(){
// 	var dataFilm = []movie{}

// tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
// tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
// tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
// tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

// // isi dengan jawaban anda untuk menampilkan data
// }
