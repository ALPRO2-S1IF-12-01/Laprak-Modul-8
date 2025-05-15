// ABID FADHILAH MUSTOFA
// 103112400046
package main

import (
	"fmt"
)

func pilkart() {
	var target int
	suaraMasuk, suaraSah := 0, 0
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

	for i, v := range jumlahVote {
		if v > 0 {
			fmt.Printf("%d: %d\n", i+1, v)
		}
	}
}

func main() {
	pilkart()
}