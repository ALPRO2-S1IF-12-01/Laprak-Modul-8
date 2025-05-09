package main

import (
	"fmt"
)
/*
Dimas Ramadhani
103112400065
*/
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
	for i := 0; i < len(calonKetua); i++ {
		if jumlahVote[i] > 0 {
			fmt.Printf("%d: %d\n", i+1, jumlahVote[i])
		}
	}
}
func main() {
	pilkart()
}
