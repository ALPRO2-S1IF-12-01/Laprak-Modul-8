//  DWI OKTA SURYANINGRUM
//  103112400066

package main

import (
	"fmt"
)

func main() {
	// map untuk nyimpen jumlah suara tiap kandidat dengan range (1-20)
	hitung := make(map[int]int)

	masuk := 0 // jumlah semua suara yang masuk
	sah := 0   // jumlah suara yang sah

	var input int
	for {
		fmt.Scan(&input)

		if input == 0 {
			break // kalau ketemu 0, berhenti baca input
		}

		masuk++ // hitung total suara masuk

		if input >= 1 && input <= 20 {
			hitung[input]++ // suara valid, masukin ke map
			sah++
		}
	}

	// menampilkan hasil akhir
	fmt.Println("Suara masuk:", masuk)
	fmt.Println("Suara sah:", sah)

	// cetak hasil masing-masing kandidat (yang dapet suara aja)
	for i := 1; i <= 20; i++ {
		if jumlah, ada := hitung[i]; ada {
			fmt.Printf("%d: %d\n", i, jumlah)
		}
	}
}
