package main

import "fmt"

const MAX = 100000
var data [MAX]int

func isiArray(n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}

func posisi(k, n int) int {
	for i := 0; i < n; i++ {
		if data[i] == k {
			return i
		}
	}
	return -1 
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	isiArray(n)
	pos := posisi(k, n)

	if pos == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(pos)
	}
}
