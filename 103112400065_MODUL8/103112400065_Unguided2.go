package main

import (
	"fmt"
)

func pilkart() {
	var target int
	var suaraMasuk, suaraSah int = 0, 0
	calonKetua := make([]int, 20)
	jumlahVote := make([]int, 20)
	for i := 0; i < 20; i++ {
		calonKetua[i] = i + 1
	}
	for {
		_, err := fmt.Scan(&target)
		if err != nil {
			break
		}
		if target == 0 {
			break
		}
		suaraMasuk++
		for i, val := range calonKetua {
			if val == target {
				suaraSah++
				jumlahVote[i]++
				break
			}
		}
	}
	fmt.Println("Suara masuk:", suaraMasuk)
	fmt.Println("Suara sah:", suaraSah)
	/*
	   Dimas Ramadhani
	   103112400065
	*/
	jumKetua1, jumKetua2 := 0, 0
	cekKetua1, cekKetua2 := -1, -1
	for i, vote := range jumlahVote {
		if vote > jumKetua1 || (vote == jumKetua1 && cekKetua1 > i) {
			jumKetua2, cekKetua2 = jumKetua1, cekKetua1
			jumKetua1, cekKetua1 = vote, i
		} else if vote > jumKetua2 || (vote == jumKetua2 && cekKetua2 > i) {
			jumKetua2, cekKetua2 = vote, i
		}
	}
	if cekKetua1 != -1 {
		fmt.Println("Ketua RT:", cekKetua1+1)
	}
	if cekKetua2 != -1 {
		fmt.Println("Wakil ketua:", cekKetua2+1)
	}
}
func main() {
	pilkart()
}
