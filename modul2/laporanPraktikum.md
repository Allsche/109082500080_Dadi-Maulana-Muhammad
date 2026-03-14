# <h1 align="center">Laporan Praktikum Modul 2 - REVIEW ALGORITMA DAN PEMROGRAMAN 1 </h1>
<p align="center">Dadi Maulana Muhammad - 109082500080</p>

## Unguided 

### 1. Soal Latihan Modul 2A
Telusuri program berikut dengan cara mengkompilasi dan mengeksekusi program. Silakan
masukan data yang sesuai sebanyak yang diminta program. Perhatikan keluaran yang
diperoleh. Coba terangkan apa sebenarnya yang dilakukan program tersebut?
#### soal1.go

```go
package main

import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp            string
	)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Allsche/109082500080_Dadi-Maulana-Muhammad/blob/main/modul2/output/output-soal1.png)

Program Go tersebut meminta pengguna memasukkan tiga buah string yang disimpan dalam variabel satu, dua, dan tiga, lalu menampilkan ketiga nilai tersebut sebagai output awal. Setelah itu program melakukan proses rotasi nilai variabel menggunakan variabel sementara temp, di mana nilai satu disimpan terlebih dahulu ke temp, kemudian satu diisi dengan nilai dua, dua diisi dengan nilai tiga, dan tiga diisi kembali dengan nilai yang ada di temp. Proses ini menyebabkan urutan nilai bergeser ke kiri dari (satu, dua, tiga) menjadi (dua, tiga, satu), kemudian hasil akhirnya ditampilkan sebagai output akhir.

### 2. Soal Latihan Modul 2B
Siswa kelas IPA di salah satu sekolah menengah atas di Indonesia sedang mengadakan
praktikum kimia. Di setiap percobaan akan menggunakan 4 tabung reaksi, yang mana
susunan warna cairan di setiap tabung akan menentukan hasil percobaan. Siswa diminta
untuk mencatat hasil percobaan tersebut. Percobaan dikatakan berhasil apabila susunan
warna zat cair pada gelas 1 hingga gelas 4 secara berturutan adalah ‘merah’, ‘kuning’,
‘hijau’, dan ‘ungu’ selama 5 kali percobaan berulang.
Buatlah sebuah program yang menerima input berupa warna dari ke 4 gelas reaksi
sebanyak 5 kali percobaan. Kemudian program akan menampilkan true apabila urutan
warna sesuai dengan informasi yang diberikan pada paragraf sebelumnya, dan false
untuk urutan warna lainnya.
Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergaris bawah
adalah input/read):
#### soal2.go

```go
package main

import "fmt"

func main() {
	var a, b, c, d string
	berhasil := true

	for i := 1; i <= 5; i++ {
		fmt.Print("Percobaan ", i, ": ")
		fmt.Scan(&a, &b, &c, &d)

		if !(a == "merah" && b == "kuning" && c == "hijau" && d == "ungu") {
			berhasil = false
		}
	}

	fmt.Println("BERHASIL:", berhasil)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Allsche/109082500080_Dadi-Maulana-Muhammad/blob/main/modul2/output/output-soal2.png)

Program Go tersebut digunakan untuk mencatat hasil 5 kali percobaan dengan memasukkan 4 warna dari tabung reaksi pada setiap percobaan. Program membaca input warna secara berurutan, lalu memeriksa apakah susunannya sama dengan urutan yang ditentukan yaitu “merah”, “kuning”, “hijau”, dan “ungu”. Pemeriksaan dilakukan menggunakan perulangan sebanyak lima kali, dan setiap kali urutan warna tidak sesuai maka variabel penanda keberhasilan akan diubah menjadi false. Setelah seluruh percobaan selesai, program akan menampilkan hasil akhir berupa true jika semua percobaan memiliki urutan warna yang benar, atau false jika ada satu atau lebih percobaan yang tidak sesuai dengan urutan yang ditentukan.

### 3. Soal Latihan Modul 2C
PT POS membutuhkan aplikasi perhitungan biaya kirim berdasarkan berat parsel. Maka,
buatlah program BiayaPos untuk menghitung biaya pengiriman tersebut dengan ketentuan
sebagai berikut!
Dari berat parsel (dalam gram), harus dihitung total berat dalam kg dan sisanya (dalam
gram). Biaya jasa pengiriman adalah Rp. 10.000,- per kg. Jika sisa berat tidak kurang dari 500
gram, maka tambahan biaya kirim hanya Rp. 5,- per gram saja. Tetapi jika kurang dari 500
gram, maka tambahan biaya akan dibebankan sebesar Rp. 15,- per gram. Sisa berat (yang
kurang dari 1kg) digratiskan biayanya apabila total berat ternyata lebih dari 10kg.
Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergaris bawah
adalah input/read):
#### soal3.go

```go
package main

import "fmt"

func main() {
	var gram, kg, sisa, biayaKg, biayaSisa, total int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&gram)

	kg = gram / 1000
	sisa = gram % 1000

	biayaKg = kg * 10000

	if kg > 10 {
		biayaSisa = 0
	} else {
		if sisa >= 500 {
			biayaSisa = sisa * 5
		} else {
			biayaSisa = sisa * 15
		}
	}

	total = biayaKg + biayaSisa

	fmt.Println("Detail berat: ", kg, "kg + ", sisa, " gr")
	fmt.Println("Detail biaya: Rp. ", biayaKg, " + Rp. ", biayaSisa)
	fmt.Println("Total biaya: Rp. ", total)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Allsche/109082500080_Dadi-Maulana-Muhammad/blob/main/modul2/output/output-soal3.png)

Program Go tersebut digunakan untuk menghitung biaya pengiriman parsel berdasarkan berat dalam gram. Program terlebih dahulu membaca input berat parsel, kemudian menghitung jumlah kilogram dan sisa gram dari berat tersebut menggunakan operasi pembagian dan modulus. Biaya dasar pengiriman dihitung sebesar Rp10.000 per kilogram, sedangkan sisa berat yang kurang dari 1 kg akan dikenakan biaya tambahan, yaitu Rp5 per gram jika sisa berat ≥ 500 gram atau Rp15 per gram jika sisa berat < 500 gram. Namun, jika total berat parsel lebih dari 10 kg, maka sisa gram tersebut tidak dikenakan biaya tambahan. Setelah semua perhitungan selesai, program menampilkan detail berat (kg dan gram), detail biaya per bagian, serta total biaya pengiriman.