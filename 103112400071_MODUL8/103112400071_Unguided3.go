// Raihan Adi Arba
// 103112400071

package main

import "fmt"

const NMAX = 1000000

var data [NMAX]int

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	isiArray(n)
	idx := posisi(n, k)

	if idx == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(idx)
	}
}

func isiArray(n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}

func posisi(n, k int) int {
	kiri, kanan := 0, n-1
	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		switch {
		case data[tengah] == k:
			return tengah
		case data[tengah] < k:
			kiri = tengah + 1
		default:
			kanan = tengah - 1
		}
	}
	return -1
}
