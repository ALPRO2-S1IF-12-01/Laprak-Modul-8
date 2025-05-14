//  DWI OKTA SURYANINGRUM
//  103112400066

package main

import (
	"fmt"
	"sort"
)

func main() {
	// map buat nyimpen jumlah suara untuk masing-masing kandidat
	suara := make(map[int]int)

	totalMasuk := 0 // semua input yang masuk (selain 0)
	suaraSah := 0   // suara yang valid (antara 1–20)

	var input int
	for {
		fmt.Scan(&input)

		if input == 0 {
			break // 0 artinya selesai input
		}

		totalMasuk++

		if input >= 1 && input <= 20 {
			suara[input]++ // menghitung suara kandidat
			suaraSah++
		}
	}

	// menampilkan ringkasan suara
	fmt.Println("Suara masuk:", totalMasuk)
	fmt.Println("Suara sah:", suaraSah)

	// struct buat nampung pasangan kandidat dan jumlah suara
	type Pasangan struct {
		nomor  int
		jumlah int
	}

	var hasil []Pasangan

	// memasukkan ke slice untuk bisa diurutkan
	for nomor, jumlah := range suara {
		hasil = append(hasil, Pasangan{nomor, jumlah})
	}

	// mengurutkan berdasarkan jumlah suara terbanyak, lalu nomor terkecil
	sort.Slice(hasil, func(i, j int) bool {
		if hasil[i].jumlah == hasil[j].jumlah {
			return hasil[i].nomor < hasil[j].nomor
		}
		return hasil[i].jumlah > hasil[j].jumlah
	})

	// memastikan minimal ada 2 kandidat untuk ditampilkan
	if len(hasil) >= 2 {
		fmt.Println("Ketua RT:", hasil[0].nomor)
		fmt.Println("Wakil ketua:", hasil[1].nomor)
	} else if len(hasil) == 1 {
		fmt.Println("Ketua RT:", hasil[0].nomor)
		fmt.Println("Wakil ketua: TIDAK ADA")
	} else {
		fmt.Println("TIDAK ADA suara sah yang masuk.")
	}
}