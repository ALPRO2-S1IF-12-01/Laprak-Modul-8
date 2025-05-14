package main

import "fmt"

func main() {
	var suara int
	hitungSuara := make(map[int]int)
	totalSuaraMasuk := 0
	totalSuaraSah := 0

	fmt.Println("NAMA: MULIA AKBAR NANDA PRATAMA\nNIM: 103112400034")
	fmt.Println("silahkan masukkan nomor calon:")

	for {
		fmt.Scan(&suara)
		totalSuaraMasuk++

		if suara == 0 {
			totalSuaraMasuk--
			break
		}

		if suara >= 1 && suara <= 20 {
			hitungSuara[suara] = hitungSuara[suara] + 1
			totalSuaraSah++
		}
	}

	fmt.Println("Suara masuk:", totalSuaraMasuk)
	fmt.Println("Suara sah:", totalSuaraSah)

	for calon, jumlah := range hitungSuara {
		if jumlah > 0 {
			fmt.Printf("%d: %d\n", calon, jumlah)
		}
	}
}