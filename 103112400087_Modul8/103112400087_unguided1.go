// 103112400087_Muhammad Shabrian
package main

import "fmt"

func main() {
	var suara, totalMasuk, totalSah int
	var count [21]int
	for {
		fmt.Scan(&suara)
		if suara == 0 {
			break
		}
		totalMasuk++
		if suara >= 1 && suara <= 20 {
			count[suara]++
			totalSah++
		}
	}

	fmt.Println("Suara masuk:", totalMasuk)
	fmt.Println("Suara sah:", totalSah)
	for i := 1; i <= 20; i++ {
		if count[i] > 0 {
			fmt.Printf("%d: %d\n", i, count[i])
		}
	}
}
