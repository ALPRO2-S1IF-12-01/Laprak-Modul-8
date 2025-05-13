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
		suaraPerCalon    [21]int
	)

	//hitung suara sah
	for i := 0; i < total; i++ {
		s := suara[i]
		if s >= 1 && s <= 20 {
			suaraPerCalon[s]++
			jumlahSuaraValid++
		}
	}

	//calon dengan suara terbanyak (ketua RT)
	var ketuaRT, wakilKetua int
	var suaraKetua, suaraWakil int

	for calon := 1; calon <= 20; calon++ {
		if suaraPerCalon[calon] > suaraKetua {
			wakilKetua = ketuaRT
			suaraWakil = suaraKetua
			ketuaRT = calon
			suaraKetua = suaraPerCalon[calon]
		} else if suaraPerCalon[calon] > suaraWakil {
			wakilKetua = calon
			suaraWakil = suaraPerCalon[calon]
		}
	}

	fmt.Printf("Suara masuk: %d\n", total)
	fmt.Printf("Suara sah: %d\n", jumlahSuaraValid)
	fmt.Printf("Ketua RT: %d\n", ketuaRT)
	fmt.Printf("Wakil ketua: %d\n", wakilKetua)
}
