package main

import "fmt"

const NMAX = 1000000

var data [NMAX]int

func isiArray(n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}

func posisi(n, k int) int {
	kiri := 0
	kanan := n - 1
	indeksDitemukan := -1

	for kiri <= kanan {
		tengah := kiri + (kanan-kiri)/2
		if data[tengah] == k {
			indeksDitemukan = tengah
			break
		} else if data[tengah] < k {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	} //103112400049 Hisyam Nurdiatmoko
	return indeksDitemukan
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	isiArray(n)
	hasilPosisi := posisi(n, k)
	if hasilPosisi != -1 {
		fmt.Println(hasilPosisi)
	} else {
		fmt.Println("TIDAK ADA")
	}
}
