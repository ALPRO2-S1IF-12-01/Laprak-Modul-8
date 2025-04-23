// Raihan Adi Arba
// 103112400071

package main

import (
	"fmt"
)

const jumlahCalon = 20

func inisialisasiCalon() []int {
	calon := make([]int, jumlahCalon)
	for i := 0; i < jumlahCalon; i++ {
		calon[i] = i + 1
	}
	return calon
}

func hitungSuara(calon []int) (totalMasuk int, totalSah int, jumlahVote []int) {
	jumlahVote = make([]int, jumlahCalon)
	var suara int

	for {
		_, err := fmt.Scan(&suara)
		if err != nil || suara == 0 {
			break
		}

		totalMasuk++

		for i := 0; i < jumlahCalon; i++ {
			if suara == calon[i] {
				jumlahVote[i]++
				totalSah++
				break
			}
		}
	}

	return totalMasuk, totalSah, jumlahVote
}

func tampilkanHasil(totalMasuk, totalSah int, calon []int, jumlahVote []int) {
	fmt.Println("Suara masuk:", totalMasuk)
	fmt.Println("Suara sah:", totalSah)
	for i := 0; i < jumlahCalon; i++ {
		if jumlahVote[i] > 0 {
			fmt.Printf("%d: %d\n", calon[i], jumlahVote[i])
		}
	}
}

func main() {
	calon := inisialisasiCalon()
	totalMasuk, totalSah, jumlahVote := hitungSuara(calon)
	tampilkanHasil(totalMasuk, totalSah, calon, jumlahVote)
}
