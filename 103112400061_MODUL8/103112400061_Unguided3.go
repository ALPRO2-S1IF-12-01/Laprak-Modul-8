package main

import "fmt"

const NMAX int = 1000000
var data [NMAX]int

func main() {
	var n, k int

	fmt.Scan(&n, &k)
	isiArray(n)
	indeks := posisi(n, k)
	if indeks != -1 {
		fmt.Println(indeks)
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
	left, right := 0, n

	for left <= right {
		mid := (left + right) / 2
		if data[mid] == k {
			return mid
		} else if k < data[mid] {
			right = mid -1
		} else {
			left = mid + 1
		}
	}

	return -1
}