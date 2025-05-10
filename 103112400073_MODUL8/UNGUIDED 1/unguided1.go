//Muhammad Zaky Mubarok
package main

import (
	"fmt"
	"sort"
)


func seqSearch(arr []int, target int) bool {
	for _, nilai := range arr {
		if nilai == target {
			return true
		}
	}
	return false
}


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
	
	calonKetua := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
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
		if len(calonKetua) <= 10 {
			if seqSearch(calonKetua, suara) {
				jumlahSuaraSah++
				perolehanSuara[suara]++
			}
		} else {
			if binSearch(calonKetua, suara) {
				jumlahSuaraSah++
				perolehanSuara[suara]++
			}
		}
	}

	
	fmt.Println("Suara masuk:", jumlahSuara)
	fmt.Println("Suara sah:", jumlahSuaraSah)
	fmt.Println()

	
	for calon := 1; calon <= 20; calon++ {
		if perolehanSuara[calon] > 0 {
			fmt.Printf("%d: %d\n", calon, perolehanSuara[calon])
		}
	}
}