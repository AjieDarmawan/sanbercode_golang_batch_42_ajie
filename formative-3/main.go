package main

import "fmt"

// soal 1;

/*
func main() {
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"

	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	pp, _ := strconv.Atoi(panjangPersegiPanjang)
	lp, _ := strconv.Atoi(lebarPersegiPanjang)
	as, _ := strconv.Atoi(alasSegitiga)
	ts, _ := strconv.Atoi(tinggiSegitiga)

	var luas_persegi_panjang = (pp * lp)
	var keliling_persegi_panjang = 2 * (pp + lp)
	var luas_segitiga = (as * ts) / 2

	//luas persegi panjang
	//keliling persegi panjang
	//luas segitiga

	fmt.Println("luas persegi panjang  \n", luas_persegi_panjang)
	fmt.Println("keliling persegi panjang  \n", keliling_persegi_panjang)
	fmt.Println("luas segitiga  \n", luas_segitiga)

}*/

//soal 2

/*
func main() {

	var nilaiJohn = 75
	var nilaiDoe = 50

	if nilaiJohn >= 80 {
		fmt.Println("Nilai John A")
	} else if nilaiJohn >= 70 && nilaiJohn < 80 {
		fmt.Println("Nilai John B")
	} else if nilaiJohn >= 60 && nilaiJohn < 70 {
		fmt.Println("Nilai John C")
	} else if nilaiJohn >= 50 && nilaiJohn < 60 {
		fmt.Println("Nilai John D")
	} else {
		fmt.Println("Nilai John E")
	}

	if nilaiDoe >= 80 {
		fmt.Println("Nilai Doe A")
	} else if nilaiDoe >= 70 && nilaiDoe < 80 {
		fmt.Println("Nilai Doe B")
	} else if nilaiDoe >= 60 && nilaiDoe < 70 {
		fmt.Println("Nilai Doe C")
	} else if nilaiDoe >= 50 && nilaiDoe < 60 {
		fmt.Println("Nilai Doe D")
	} else {
		fmt.Println("Nilai Doe E")
	}
*/

//}

//soal 3

/*
func main() {
	var tanggal = 17
	var bulan = 4
	var tahun = 1996

	var bulanindo string

	switch bulan {
	case 1:
		bulanindo = "Januari"
	case 2:
		bulanindo = "Februari"
	case 3:
		bulanindo = "Maret"
	case 4:
		bulanindo = "April"
	case 5:
		bulanindo = "Mei"
	case 6:
		bulanindo = "Juni"
	case 7:
		bulanindo = "Juli"
	case 8:
		bulanindo = "Agustus"
	case 9:
		bulanindo = "September"
	case 10:
		bulanindo = "Oktober"
	case 11:
		bulanindo = "November"
	case 12:
		bulanindo = "Desember"
	}

	fmt.Println(tanggal, bulanindo, tahun)

}
*/

func main() {

	var tahun = 1996

	// Baby boomer, kelahiran 1944 s.d 1964
	// Generasi X, kelahiran 1965 s.d 1979
	// Generasi Y (Millenials), kelahiran 1980 s.d 1994
	// Generasi Z, kelahiran 1995 s.d 2015

	if tahun >= 1944 && tahun <= 1964 {
		fmt.Println("Generasi Baby Boomer")
	} else if tahun >= 1965 && tahun <= 1979 {
		fmt.Println("Generasi X")
	} else if tahun >= 1980 && tahun <= 1994 {
		fmt.Println("Generasi Y Millenials")
	} else if tahun >= 1995 && tahun <= 2015 {
		fmt.Println("Generasi Z")
	} else {
		fmt.Println("Belum ada istilah nya berarti")
	}
}
