# <h1 align="center">Laporan Praktikum Modul 5</h1>
<p align="center">Dadi Maulana Muhammad - 109082500080</p>

## Unguided 

### 1. Soal Latihan Modul 5A
Deret fibonacci adalah sebuah deret dengan nilai suku ke-0 dan ke-1 adalah 0 dan 1, dan nilai
suku ke-n selanjutnya adalah hasil penjumlahan dua suku sebelumnya. Secara umum dapat
diformulasikan Sn = Sn−1 + Sn−2 . Berikut ini adalah contoh nilai deret fibonacci hingga suku
ke-10. Buatlah program yang mengimplementasikan fungsi rekursif pada deret fibonacci
tersebut.
#### soal1.go

```go
package main

import "fmt"

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i <= n; i++ {
		fmt.Print(fibonacci(i), " ")
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Allsche/109082500080_Dadi-Maulana-Muhammad/blob/main/modul5/output/output-soal1.png)

Program ini menghitung deret Fibonacci menggunakan fungsi rekursif fibonacci(n), di mana kondisi dasar (base case) adalah saat n == 0 menghasilkan 0 dan n == 1 menghasilkan 1. Untuk nilai selain itu, fungsi akan memanggil dirinya sendiri dengan rumus fibonacci(n-1) + fibonacci(n-2), sehingga setiap suku merupakan jumlah dua suku sebelumnya. Di main, program membaca input n lalu mencetak deret dari indeks 0 hingga n dengan memanggil fungsi tersebut berulang kali. Pendekatan ini sederhana tetapi kurang efisien untuk nilai besar karena banyak perhitungan yang berulang.

### 2. Soal Latihan Modul 5B
Buatlah sebuah program yang digunakan untuk menampilkan pola bintang berikut ini dengan
menggunakan fungsi rekursif. N adalah masukan dari user.
#### soal2.go

```go
package main

import "fmt"

func printBintang(n int, nilai int) {
	if nilai > n {
		return
	}
	for i := 0; i < nilai; i++ {
		fmt.Print("*")
	}
	fmt.Println()
	printBintang(n, nilai+1)
}

func main() {
	var n int
	fmt.Scan(&n)
	printBintang(n, 1)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Allsche/109082500080_Dadi-Maulana-Muhammad/blob/main/modul5/output/output-soal2.png)

Program ini mencetak pola segitiga bintang menggunakan fungsi rekursif printStars(n, current), di mana current menunjukkan baris yang sedang dicetak. Jika current lebih besar dari n, rekursi berhenti. Pada setiap pemanggilan, program mencetak sejumlah bintang sesuai nilai current, lalu memanggil dirinya sendiri dengan current+1 untuk mencetak baris berikutnya. Dengan cara ini, pola terbentuk dari 1 bintang hingga n bintang secara bertahap.

### 3. Soal Latihan Modul 5C
Buatlah program yang mengimplementasikan rekursif untuk menampilkan faktor bilangan dari
suatu N, atau bilangan yang apa saja yang habis membagi N.
Masukan terdiri dari sebuah bilangan bulat positif N.
Keluaran terdiri dari barisan bilangan yang menjadi faktor dari N (terurut dari 1 hingga N ya).
#### soal3.go

```go
package main

import "fmt"

func faktor(n int, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Print(i, " ")
	}
	faktor(n, i+1)
}

func main() {
	var n int
	fmt.Scan(&n)
	faktor(n, 1)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Allsche/109082500080_Dadi-Maulana-Muhammad/blob/main/modul5/output/output-soal3.png)

Program ini menampilkan semua faktor dari sebuah bilangan n menggunakan fungsi rekursif faktor(n, i), dengan i sebagai pembagi yang diuji mulai dari 1 hingga n. Jika i lebih besar dari n, proses berhenti. Pada setiap langkah, program memeriksa apakah n % i == 0, yang berarti i adalah faktor dari n, lalu mencetaknya. Fungsi kemudian memanggil dirinya sendiri dengan i+1 untuk melanjutkan pengecekan hingga semua faktor ditemukan.

### 3. Soal Latihan Modul 5D
Buatlah program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan
tertentu.
Masukan terdiri dari sebuah bilangan bulat positif N.
Keluaran terdiri dari barisan bilangan dari N hingga 1 dan kembali ke N.
#### soal3.go

```go
package main

import "fmt"

func pola(n int) {
	if n == 0 {
		return
	}
	fmt.Print(n, " ")
	pola(n - 1)
	if n != 1 {
		fmt.Print(n, " ")
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	pola(n)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Allsche/109082500080_Dadi-Maulana-Muhammad/blob/main/modul5/output/output-soal4.png)

Program ini mencetak pola bilangan dari n turun ke 1 lalu kembali naik ke n menggunakan fungsi rekursif pola(n). Jika n == 0, rekursi berhenti. Sebelum pemanggilan rekursif, program mencetak nilai n, lalu memanggil pola(n-1) untuk menurunkan nilai hingga 1. Setelah pemanggilan selesai (saat kembali dari rekursi), program mencetak kembali nilai n kecuali saat n == 1, sehingga terbentuk pola simetris seperti 5 4 3 2 1 2 3 4 5.

### 3. Soal Latihan Modul 5E
Buatlah program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan ganjil.
Masukan terdiri dari sebuah bilangan bulat positif N.
Keluaran terdiri dari barisan bilangan ganjil dari 1 hingga N.
#### soal3.go

```go
package main

import "fmt"

func ganjil(n int, i int) {
	if i > n {
		return
	}
	if i%2 != 0 {
		fmt.Print(i, " ")
	}
	ganjil(n, i+1)
}

func main() {
	var n int
	fmt.Scan(&n)
	ganjil(n, 1)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Allsche/109082500080_Dadi-Maulana-Muhammad/blob/main/modul5/output/output-soal5.png)

Program ini mencetak bilangan ganjil dari 1 hingga n menggunakan fungsi rekursif ganjil(n, i), dengan i dimulai dari 1. Jika i melebihi n, rekursi berhenti. Pada setiap langkah, program mengecek apakah i adalah bilangan ganjil (i % 2 != 0), dan jika iya maka dicetak. Fungsi kemudian memanggil dirinya sendiri dengan i+1 untuk melanjutkan ke angka berikutnya sampai mencapai batas n.

### 3. Soal Latihan Modul 5F
Buatlah program yang mengimplementasikan rekursif untuk mencari hasil pangkat dari dua
buah bilangan.
Masukan terdiri dari bilangan bulat x dan y.
Keluaran terdiri dari hasil x dipangkatkan y.
Catatan: diperbolehkan menggunakan asterik "*", tapi dilarang menggunakan import "math".
#### soal3.go

```go
package main

import "fmt"

func pangkat(x int, y int) int {
	if y == 0 {
		return 1
	}
	return x * pangkat(x, y-1)
}

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	fmt.Println(pangkat(x, y))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Allsche/109082500080_Dadi-Maulana-Muhammad/blob/main/modul5/output/output-soal6.png)

Program ini menghitung hasil perpangkatan x^y menggunakan fungsi rekursif pangkat(x, y). Kondisi dasar adalah ketika y == 0, hasilnya 1 karena bilangan apa pun berpangkat 0 bernilai 1. Untuk kondisi lainnya, fungsi mengembalikan hasil x * pangkat(x, y-1), yang berarti mengalikan x secara berulang sebanyak y kali melalui rekursi. Program utama membaca input x dan y, lalu mencetak hasil perhitungan tersebut.