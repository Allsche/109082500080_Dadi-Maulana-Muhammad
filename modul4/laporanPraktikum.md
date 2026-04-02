# <h1 align="center">Laporan Praktikum Modul 4</h1>
<p align="center">Dadi Maulana Muhammad - 109082500080</p>

## Unguided 

### Soal Latihan Modul 4.1
Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika
diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng
untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian
membantu Jonas? (tidak tentunya ya :p)
Masukan terdiri dari empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi,
dengan syarat a ≥ c dan b ≥ d.
Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi a
terhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d.

```go
package main

import "fmt"

func factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func permutasi(n, r int) int {
	return factorial(n) / factorial(n-r)
}

func kombinasi(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	p1 := permutasi(a, c)
	c1 := kombinasi(a, c)
	p2 := permutasi(b, d)
	c2 := kombinasi(b, d)

	fmt.Println(p1, c1)
	fmt.Println(p2, c2)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Allsche/109082500080_Dadi-Maulana-Muhammad/blob/main/modul4/output/output-soal1.png)

Program Golang ini digunakan untuk menghitung permutasi dan kombinasi dari dua pasang bilangan yang diinput oleh pengguna. Pertama, dibuat fungsi factorial untuk menghitung nilai faktorial dari suatu bilangan dengan perulangan. Kemudian, fungsi permutasi(n, r) menghitung banyaknya susunan (P(n, r)) dengan rumus n!/(n−r)!, sedangkan fungsi kombinasi(n, r) menghitung banyaknya cara memilih (C(n, r)) dengan rumus n!/(r!(n−r)!). Di dalam fungsi main, program membaca empat bilangan a, b, c, d, lalu menghitung permutasi dan kombinasi untuk pasangan (a, c) dan (b, d). Hasilnya disimpan dalam variabel p1, c1, p2, c2, kemudian dicetak ke layar dalam dua baris, masing-masing berisi hasil permutasi dan kombinasi dari setiap pasangan bilangan.

### Soal Latihan Modul 4.2
Kompetisi pemrograman tingkat nasional berlangsung ketat. Setiap peserta diberikan 8 soal
yang harus dapat diselesaikan dalam waktu 5 jam saja. Peserta yang berhasil menyelesaikan
soal paling banyak dalam waktu paling singkat adalah pemenangnya.
Buat program gema yang mencari pemenang dari daftar peserta yang diberikan. Program
harus dibuat modular, yaitu dengan membuat prosedur hitungSkor yang mengembalikan total
soal dan total skor yang dikerjakan oleh seorang peserta, melalui parameter formal.
Pembacaan nama peserta dilakukan di program utama, sedangkan waktu pengerjaan dibaca
di dalam prosedur.
prosedure hitungSkor(in/out soal, skor : integer)
Setiap baris masukan dimulai dengan satu string nama peserta tersebut diikuti dengan adalah
8 integer yang menyatakan berapa lama (dalam menit) peserta tersebut menyelesaikan soal.
Jika tidak berhasil atau tidak mengirimkan jawaban maka otomatis dianggap menyelesaikan
dalam waktu 5 jam 1 menit (301 menit).
Satu baris keluaran berisi nama pemenang, jumlah soal yang diselesaikan, dan nilai yang
diperoleh. Nilai adalah total waktu yang dibutuhkan untuk menyelesaikan soal yang berhasil
diselesaikan.
#### soal2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Allsche/109082500080_Dadi-Maulana-Muhammad/blob/main/modul4/output/output-soal2.png)

Program Golang ini digunakan untuk menentukan pemenang kompetisi berdasarkan jumlah soal yang berhasil diselesaikan dan total waktu pengerjaan. Fungsi hitungSkor membaca 8 buah waktu pengerjaan soal, lalu menghitung berapa soal yang berhasil diselesaikan (waktu kurang dari 301 menit) serta menjumlahkan total waktunya, kemudian mengembalikan kedua nilai tersebut. Pada fungsi main, program membaca nama peserta satu per satu hingga ditemukan kata "Selesai", dan untuk setiap peserta akan dipanggil hitungSkor. Selanjutnya, program membandingkan hasil setiap peserta untuk menentukan pemenang, yaitu peserta dengan jumlah soal terbanyak, atau jika sama maka yang memiliki total waktu paling kecil. Variabel first digunakan untuk memastikan peserta pertama langsung dijadikan acuan awal. Di akhir, program mencetak nama pemenang beserta jumlah soal yang diselesaikan dan total waktunya.

### Soal Latihan Modul 4.3
Skiena dan Revilla dalam Programming Challenges mendefinisikan sebuah deret bilangan.
Deret dimulai dengan sebuah bilangan bulat n. Jika bilangan n saat itu genap, maka suku
berikutnya adalah 1⁄2n, tetapi jika ganjil maka suku berikutnya bernilai 3n+1. Rumus yang sama
digunakan terus menerus untuk mencari suku berikutnya. Deret berakhir ketika suku terakhir bernilai 1. Sebagai contoh jika dimulai dengan n=22, maka deret bilangan yang diperoleh
adalah:

22 11 34 17 52 26 13 40 20 10 5 16 8 4 2 1

Untuk suku awal sampai dengan 1000000, diketahui deret selalu mencapai suku dengan nilai
1.
Buat program skiena yang akan mencetak setiap suku dari deret yang dijelaskan di atas untuk
nilai suku awal yang diberikan. Pencetakan deret harus dibuat dalam prosedur cetakDeret
yang mempunyai 1 parameter formal, yaitu nilai dari suku awal.
prosedure cetakDeret(in n : integer )
Masukan berupa satu bilangan integer positif yang lebih kecil dari 1000000.
Keluaran terdiri dari satu baris saja. Setiap suku dari deret tersebut dicetak dalam baris yang
dan dipisahkan oleh sebuah spasi.
#### soal3.go

```go
package main

import "fmt"

func cetakDeret(n int) {
	for {
		fmt.Print(n)

		if n == 1 {
			break
		}
		fmt.Print(" ")

		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	cetakDeret(n)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Allsche/109082500080_Dadi-Maulana-Muhammad/blob/main/modul4/output/output-soal3.png)

Program Golang ini digunakan untuk mencetak deret bilangan berdasarkan aturan Collatz (3n+1) dari sebuah bilangan awal yang diinput pengguna. Fungsi cetakDeret(n int) bertugas mencetak nilai n secara berurutan dalam satu baris, kemudian selama nilainya belum mencapai 1, program akan mengecek apakah n genap atau ganjil: jika genap maka dibagi 2, dan jika ganjil maka dikalikan 3 lalu ditambah 1. Setiap hasil dicetak dan dipisahkan dengan spasi hingga akhirnya mencapai angka 1 dan perulangan berhenti. Pada fungsi main, program hanya membaca satu bilangan integer sebagai input, lalu memanggil prosedur cetakDeret untuk menampilkan seluruh deret tersebut.