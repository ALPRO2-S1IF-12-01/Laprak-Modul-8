package main

import "fmt"

const NMAX = 1000000

var data [NMAX]int

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	isiArray(n)

	pos := posisi(n, k)
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
	for i := 0; i < n; i++ {
		if data[i] == k {
			return i
		}
	}
	return -1
}

