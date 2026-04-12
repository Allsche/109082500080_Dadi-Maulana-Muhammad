package main

import "fmt"

func main() {
	var uang, kurs, admin int
	fmt.Scan(&uang, &kurs, &admin)

	fmt.Println("=== MESIN PENUKARAN ===")
	fmt.Println("Uang masuk :", uang)
	fmt.Println("Biaya admin :", admin)

	// cek uang cukup untuk admin
	if uang <= admin {
		fmt.Println("Status : DITOLAK")
		fmt.Println("Alasan : Uang tidak cukup untuk biaya admin.")
		return
	}

	sisaUang := uang - admin
	fmt.Println("Sisa uang :", sisaUang)

	poin := 0

	// konversi menggunakan loop
	for sisaUang >= kurs {
		sisaUang -= kurs
		poin++
	}

	// cek apakah dapat poin
	if poin == 0 {
		fmt.Println("Status : DITOLAK")
		fmt.Println("Alasan : Uang tidak cukup untuk ditukar.")
		return
	}

	// bonus challenge
	if poin >= 10 {
		poin++
		fmt.Println("Bonus : +1 poin")
	}

	fmt.Println("Total poin :", poin)
	fmt.Println("Sisa rupiah :", sisaUang)

	if sisaUang >= 1000 {
		fmt.Println("Info : Sisa bisa ditabung")
	}

	fmt.Println("Status : BERHASIL")
}
