# <h1 align="center">Laporan Praktikum Modul 3</h1>
<p align="center">Dadi Maulana Muhammad - 109082500080</p>

## Unguided 

### Soal Latihan Modul 3.1
Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika
diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng
untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian
membantu Jonas? (tidak tentunya ya :p)
Masukan terdiri dari empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi,
dengan syarat a ≥ c dan b ≥ d.
Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi a
terhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d.
Catatan: permutasi (P) dan kombinasi (C) dari n terhadap r (n ≥ r) dapat dihitung dengan
menggunakan persamaan berikut!

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

func permutation(n, r int) int {
	return factorial(n) / factorial(n-r)
}

func combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	fmt.Println(permutation(a, c), combination(a, c))
	fmt.Println(permutation(b, d), combination(b, d))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Allsche/109082500080_Dadi-Maulana-Muhammad/blob/main/modul3/output/output-soal1.png)

Program Go tersebut membaca empat bilangan bulat a, b, c, dan d, kemudian menghitung nilai permutasi dan kombinasi menggunakan fungsi bantu. Fungsi factorial digunakan untuk menghitung faktorial suatu bilangan dengan perulangan, yang kemudian dimanfaatkan oleh fungsi permutation untuk menghitung P(n,r)=n!/(n−r)! dan fungsi combination untuk menghitung C(n,r)=n!/(r!(n−r)!). Setelah menerima input, program mencetak hasil permutasi dan kombinasi untuk pasangan a terhadap c pada baris pertama, serta b terhadap d pada baris kedua, sehingga menghasilkan dua baris output sesuai perhitungan matematika kombinasi dan permutasi.

### Soal Latihan Modul 3.2
Diberikan tiga buah fungsi matematika yaitu f (x) = x^2, g (x) = x − 2 dan h (x) = x + 1. Fungsi komposisi (fogoh)(x) artinya adalah f(g(h(x))). Tuliskan f(x), g(x) dan h(x) dalam bentuk function.
Masukan terdiri dari sebuah bilangan bulat a, b dan c yang dipisahkan oleh spasi.
Keluaran terdiri dari tiga baris. Baris pertama adalah (fogoh)(a), baris kedua (gohof)(b),
dan baris ketiga adalah (hofog)(c)!
#### soal2.go

```go
package main

import "fmt"

func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	fmt.Println(f(g(h(a))))
	fmt.Println(g(h(f(b))))
	fmt.Println(h(f(g(c))))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Allsche/109082500080_Dadi-Maulana-Muhammad/blob/main/modul3/output/output-soal2.png)

Program ini membaca tiga bilangan bulat a, b, dan c, lalu menghitung hasil dari tiga fungsi komposisi berdasarkan fungsi dasar yang telah ditentukan yaitu f(x) = x², g(x) = x − 2, dan h(x) = x + 1. Komposisi pertama (fogoh)(a) berarti menghitung 𝑓(𝑔(ℎ(𝑎))), yaitu nilai a diproses oleh fungsi h, kemudian hasilnya dimasukkan ke g, dan akhirnya ke f. Komposisi kedua (gohof)(b) berarti g(h(f(b))), dan komposisi ketiga (hofog)(c) berarti h(f(g(c))). Setiap hasil komposisi dicetak pada baris yang berbeda sesuai urutan yang diminta.

### Soal Latihan Modul 3.3
[Lingkaran] Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius
r. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y)
berdasarkan dua lingkaran tersebut.
Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat
dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik
sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan
bilangan bulat.
Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik
di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".
#### soal3.go

```go
package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt((a-c)*(a-c) + (b-d)*(b-d))
}

func didalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}

func main() {
	var cx1, cy1, r1 float64
	var cx2, cy2, r2 float64
	var x, y float64

	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	dalam1 := didalam(cx1, cy1, r1, x, y)
	dalam2 := didalam(cx2, cy2, r2, x, y)

	if dalam1 && dalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Allsche/109082500080_Dadi-Maulana-Muhammad/blob/main/modul3/output/output-soal3.png)

Program ini membaca input berupa dua lingkaran (masing-masing terdiri dari titik pusat dan radius) serta satu titik sembarang, lalu menentukan posisi titik tersebut terhadap kedua lingkaran. Fungsi jarak digunakan untuk menghitung jarak antara titik pusat lingkaran dan titik sembarang menggunakan rumus Euclidean, sedangkan fungsi didalam memeriksa apakah jarak tersebut kurang dari atau sama dengan radius lingkaran sehingga titik berada di dalam lingkaran. Hasil pengecekan untuk kedua lingkaran disimpan dalam variabel boolean, kemudian program menentukan output apakah titik berada di dalam kedua lingkaran, hanya di salah satu, atau di luar keduanya sesuai kondisi yang terpenuhi.