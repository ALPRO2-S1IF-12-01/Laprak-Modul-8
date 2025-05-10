// JESIKA METANIA RAHMA ARIFIN
// 103112400080

package main

import "fmt"

const NMAX = 1000000
var data [NMAX]int

func main() {
	var n, k int
	fmt.Scan(&n, &k)         // baca jumlah data dan angka yang dicari
	isiArray(n)              // isi array dengan n bilangan
	pos := posisi(n, k)      // cari posisi k

	if pos != -1 {
		fmt.Println(pos)
	} else {
		fmt.Println("TIDAK ADA")
	}
}

func isiArray(n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}

func posisi(n, k int) int {
	low := 0
	high := n - 1

	for low <= high {
		mid := (low + high) / 2
		if data[mid] == k {
			return mid
		} else if data[mid] < k {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
