package main

import "fmt"

func hitungSkor() (int, int) {
	soal := 0
	skor := 0

	for i := 0; i < 8; i++ {
		var waktu int
		fmt.Scan(&waktu)

		if waktu < 301 {
			soal++
			skor += waktu
		}
	}
	return soal, skor
}

func main() {
	var nama string
	var pemenang string
	var maxSoal, minSkor int
	first := true

	for {
		fmt.Scan(&nama)
		if nama == "Selesai" {
			break
		}
		soal, skor := hitungSkor()

		if first || soal > maxSoal || (soal == maxSoal && skor < minSkor) {
			pemenang = nama
			maxSoal = soal
			minSkor = skor
			first = false
		}
	}

	fmt.Println(pemenang, maxSoal, minSkor)
}
