package main

import "fmt"

func main() {
	var n, tot, sah int
	c := make(map[int]int)

	for {
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		tot++
		if n >= 1 && n <= 20 {
			sah++
			c[n]++
		}
	}

	fmt.Println("Suara masuk:", tot)
	fmt.Println("Suara sah:", sah)
	for i := 1; i <= 20; i++ {
		if c[i] > 0 {
			fmt.Printf("%d: %d\n", i, c[i])
		}
	}
}

