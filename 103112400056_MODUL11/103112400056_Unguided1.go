// Felix Pedrosa V

package main

import (
	"fmt"
	"sort"
)

// Fungsi pencarian biner sederhana
func cariIndex(data []int, cari int) bool {
	kiri := 0
	kanan := len(data) - 1
	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		if data[tengah] == cari {
			return true
		} else if data[tengah] < cari {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return false
}

// Fungsi pengecekan apakah input suara sah
func isValid(suara int, daftar []int) bool {
	return cariIndex(daftar, suara)
}

func main() {
	var suara int
	kandidatSah := make([]int, 20)
	for i := range kandidatSah {
		kandidatSah[i] = i + 1
	}
	sort.Ints(kandidatSah)

	hitungSuara := make([]int, 21)
	jumlahMasukan := 0
	jumlahSah := 0

	for {
		_, err := fmt.Scan(&suara)
		if err != nil {
			break
		}
		if suara == 0 {
			break
		}
		jumlahMasukan++
		if isValid(suara, kandidatSah) {
			hitungSuara[suara]++
			jumlahSah++
		}
	}

	fmt.Printf("Suara masuk: %d\n", jumlahMasukan)
	fmt.Printf("Suara sah: %d\n", jumlahSah)

	for i := 1; i <= 20; i++ {
		if hitungSuara[i] > 0 {
			fmt.Printf("%d: %d\n", i, hitungSuara[i])
		}
	}
}