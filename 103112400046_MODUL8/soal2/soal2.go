// ABID FADHILAH MUSTOFA
// 103112400046
package main

import (
	"fmt"
)

func pilkart() {
	var target, suaraMasuk, suaraSah int
	jumlahVote := make([]int, 20)

	for {
		_, err := fmt.Scan(&target)
		if err != nil || target == 0 {
			break
		}
		suaraMasuk++
		if target >= 1 && target <= 20 {
			suaraSah++
			jumlahVote[target-1]++
		}
	}

	fmt.Println("Suara masuk:", suaraMasuk)
	fmt.Println("Suara sah:", suaraSah)

	jumKetua1, jumKetua2 := 0, 0
	cekKetua1, cekKetua2 := -1, -1

	for i, vote := range jumlahVote {
		if vote > jumKetua1 || (vote == jumKetua1 && i < cekKetua1) {
			jumKetua2, cekKetua2 = jumKetua1, cekKetua1
			jumKetua1, cekKetua1 = vote, i
		} else if vote > jumKetua2 || (vote == jumKetua2 && i < cekKetua2) {
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
