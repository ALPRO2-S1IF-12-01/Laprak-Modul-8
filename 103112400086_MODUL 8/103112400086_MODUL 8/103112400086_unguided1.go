package main

import "fmt"

func main() {
	var (
		dataSuara []int
		a         int
		total     int
	)

	for {
		fmt.Scan(&a)
		if a == 0 {
			break
		}
		dataSuara = append(dataSuara, a)
		total++
	}

	hitungHasilPemilihan(dataSuara, total)
}

func hitungHasilPemilihan(suara []int, total int) {
	var (
		jumlahSuaraValid int
		tercatat         [21]bool
	)

	//hitung suara sah
	for i := 0; i < total; i++ {
		s := suara[i]
		if s >= 1 && s <= 20 {
			if !tercatat[s] {
				tercatat[s] = true
			}
			jumlahSuaraValid++
		}
	}

	fmt.Printf("Suara masuk : %d\n", total)
	fmt.Printf("Suara sah : %d\n", jumlahSuaraValid)

	for calon := 1; calon <= 20; calon++ {
		if tercatat[calon] {
			jumlah := hitungSuaraCalon(suara, total, calon)
			fmt.Printf("%d : %d\n", calon, jumlah)
		}
	}
}

func hitungSuaraCalon(data []int, panjang int, calon int) int {
	jumlah := 0
	for i := 0; i < panjang; i++ {
		if data[i] == calon {
			jumlah++
		}
	}
	return jumlah
}
