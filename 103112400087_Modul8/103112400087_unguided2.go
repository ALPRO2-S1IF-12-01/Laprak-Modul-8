// 103112400087_Muhammad Shabrian
package main

import "fmt"

func main() {
	var suara int
	var totalMasuk, totalSah int
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

	var ketua, wakil int
	for i := 1; i <= 20; i++ {
		if count[i] > count[ketua] || (count[i] == count[ketua] && i < ketua) {
			wakil = ketua
			ketua = i
		} else if count[i] > count[wakil] || (count[i] == count[wakil] && i < wakil && i != ketua) {
			wakil = i
		}
	}

	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakil)
}
