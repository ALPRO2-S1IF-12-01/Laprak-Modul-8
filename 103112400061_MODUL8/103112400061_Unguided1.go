package main

import (
	"fmt"
)

const N int = 1000

func main() {
	var (
		T [N]int
		x int = -1
	)

	for i := 0; x != 0; i++{
		fmt.Scan(&x)
		T[i] = x
	}
	
	pilkart(T)
}

func pilkart(T [N]int) {

	var (
		sahCounter int = 0
		allSuara int = 0
		bilPrinted [21]bool
	) // 103112400061_Keishin

	for i := 0; T[i] != 0; i++ {
		allSuara++
	}

	for i := 0; i < allSuara; i++{
		val := T[i]
		if val >= 1 && val <= 20 {
			if !bilPrinted[val] {
				bilPrinted[val] = true
			}
			sahCounter++
		}
	}

	fmt.Printf("Suara masuk: %d\n", allSuara)
	fmt.Printf("Suara sah: %d\n", sahCounter)

	for i := 1; i <= 20; i++ {
		if bilPrinted[i] {
			tempCounter := 0
			
			for j := 0; j < allSuara; j++ {
				if T[j] == i {
					tempCounter++
				}
			}
			fmt.Printf("%d: %d\n", i, tempCounter)
		}
	}
}