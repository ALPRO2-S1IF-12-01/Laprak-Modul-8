//Muhammad Gamel Al Ghifari
//103112400028
package main

import (
	"fmt"
)

func main() {
	var n int
	vote := make(map[int]int)
	total, sah := 0, 0

	fmt.Println("Input suara (akhiri dengan 0):")
	for {
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		total++
		if n >= 1 && n <= 20 {
			vote[n]++
			sah++
		}
	}

	fmt.Println("Suara masuk:", total)
	fmt.Println("Suara sah:", sah)
	for i := 1; i <= 20; i++ {
		if vote[i] > 0 {
			fmt.Printf("%d: %d\n", i, vote[i])
		}
	}
}