//Anastasia Adinda Narendra Indrianto
//103112400085
package main

import "fmt"

const NMAX = 1000000
var data [NMAX]int

func main() {
	var anastasiaN, k int
	fmt.Scan(&anastasiaN, &k)             
	isiArray(anastasiaN)                 
	idx := posisi(anastasiaN, k)         

	if idx == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(idx)
	}
}

func isiArray(anastasiaN int) {
	for i := 0; i < anastasiaN; i++ {
		fmt.Scan(&data[i])
	}
}

func posisi(anastasiaN, k int) int {
	low := 0
	high := anastasiaN - 1

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
