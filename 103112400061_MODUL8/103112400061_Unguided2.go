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
	
	selectedrt(T)
}

func selectedrt(T [N]int) {

	var (
		sahCounter int = 0
		allSuara int = 0
		winnerKetua, winnerWakil int
		tempCounter1, tempCounter2 int
		bilPrinted [21]bool
		counter [21]int
	)

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
	
	for i := 0; i < allSuara; i++ {
		val := T[i]
		if val >= 1 && val <= 20 { // 103112400061_Keishin	
			counter[val]++
		}
	}

	for i := 1; i <= 20; i++ {
		if counter[i] > tempCounter1 {
			winnerWakil = winnerKetua
			tempCounter2 = tempCounter1

			winnerKetua = i
			tempCounter1 = counter[i]
		} else if counter[i] >= tempCounter2 && counter[i] <= tempCounter1 {
			winnerWakil = i
			tempCounter2 = counter[i]
		}
	}

	fmt.Printf("Suara masuk: %d\n", allSuara)
	fmt.Printf("Suara sah: %d\n", sahCounter)
	fmt.Printf("Ketua RT: %d\n", winnerKetua)
	fmt.Printf("Wakil Ketua RT: %d\n", winnerWakil)
}