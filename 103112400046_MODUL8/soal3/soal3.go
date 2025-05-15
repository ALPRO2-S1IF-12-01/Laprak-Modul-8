// ABID FADHILAH MUSTOFA
// 103112400046
package main

import "fmt"

const NMAX = 10000
var data [NMAX]int

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	isiArray(n)

	letak := posisi(n, k)
	if letak == -1 {
		fmt.Print("TIDAK ADA")
	} else {
		fmt.Print(letak)
	}
}

func isiArray(n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}

func posisi(n, target int) int {
	low, high := 0, n-1
	for low <= high {
		mid := (low + high) / 2
		if data[mid] == target {
			return mid
		} else if target < data[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

