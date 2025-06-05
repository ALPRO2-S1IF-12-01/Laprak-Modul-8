package main

import (
	"fmt"
	"sort"
)

func binSearch(arr []int, target int) bool {
	bawah := 0
	atas := len(arr) - 1
	for bawah <= atas {
		tengah := (bawah + atas) / 2
		if arr[tengah] == target {
			return true
		} else if target < arr[tengah] {
			atas = tengah - 1
		} else {
			bawah = tengah + 1
		}
	}
	return false
}

func main() {

	calonKetua := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	var dataSuara []int
	var suara int

	fmt.Println("Masukkan suara pemilih (akhiri dengan angka 0):")
	for {
		fmt.Scan(&suara)
		if suara == 0 {
			break
		}
		dataSuara = append(dataSuara, suara)
	}

	jumlahSuara := len(dataSuara)
	jumlahSuaraSah := 0
	perolehanSuara := make(map[int]int)

	sort.Ints(calonKetua)

	for _, suara := range dataSuara {
		if binSearch(calonKetua, suara) {
			jumlahSuaraSah++
			perolehanSuara[suara]++
		}
	}

	var daftarCalon []int
	for calon := range perolehanSuara {
		daftarCalon = append(daftarCalon, calon)
	}

	sort.Slice(daftarCalon, func(i, j int) bool {
		return perolehanSuara[daftarCalon[i]] > perolehanSuara[daftarCalon[j]] ||
			(perolehanSuara[daftarCalon[i]] == perolehanSuara[daftarCalon[j]] &&
				daftarCalon[i] < daftarCalon[j])
	})

	fmt.Println("Suara masuk:", jumlahSuara)
	fmt.Println("Suara sah:", jumlahSuaraSah)

	if len(daftarCalon) > 0 {
		fmt.Println("\nKetua RT:", daftarCalon[0])
		if len(daftarCalon) > 1 {
			fmt.Println("Wakil Ketua RT:", daftarCalon[1])
		} else {
			fmt.Println("Wakil Ketua RT: Belum ada cukup suara")
		}
	} else {
		fmt.Println("\nBelum ada suara yang sah untuk menentukan Ketua RT.")
	}
}
