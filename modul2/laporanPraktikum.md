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
[penjelasan]
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
		fmt.Printf("Percobaan %d: ", i)
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
[penjelasan]
Program Go tersebut digunakan untuk mencatat hasil 5 kali percobaan dengan memasukkan 4 warna dari tabung reaksi pada setiap percobaan. Program membaca input warna secara berurutan, lalu memeriksa apakah susunannya sama dengan urutan yang ditentukan yaitu “merah”, “kuning”, “hijau”, dan “ungu”. Pemeriksaan dilakukan menggunakan perulangan sebanyak lima kali, dan setiap kali urutan warna tidak sesuai maka variabel penanda keberhasilan akan diubah menjadi false. Setelah seluruh percobaan selesai, program akan menampilkan hasil akhir berupa true jika semua percobaan memiliki urutan warna yang benar, atau false jika ada satu atau lebih percobaan yang tidak sesuai dengan urutan yang ditentukan.