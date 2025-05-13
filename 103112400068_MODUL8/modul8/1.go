package main

import (
	"fmt"
	"sort"
)

func sequentialSearch(arr []float64, target float64) (int, int) {
	iteration := 0
	for i, val := range arr {
		iteration++
		fmt.Printf("sequential Step %d: cek arr[%d] = %.1f\n", iteration, i, val)
		if val == target {
			return i, iteration
		}
	}
	return -1, iteration
}
func binarySearch(arr []float64, target float64) (int, int) {
	iterations := 0
	low := 0
	high := len(arr) - 1

	for low <= high {
		iterations++
		mid := (low + high) / 2
		fmt.Printf("binary step %d: cek arr[%d]) = %.1f\n", iterations, mid, arr[mid])
		if arr[mid] == target {
			return mid, iterations
		} else if target < arr[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1, iterations
}
func main() {
	data := []float64{2, 7, 9, 1, 5, 6, 18, 13, 25, 20}
	target := 13.0
	fmt.Println("sequential search ( data tidak perlu urut):")
	idxSeq, iterSeq := sequentialSearch(data, target)
	if idxSeq != 1 {
		fmt.Printf("hasil : Ditemukan di indeks %d dalam %d langkah\n", idxSeq, iterSeq)
	} else {
		fmt.Printf("tidak ditemukan hasil dalam %d langkah\n", iterSeq)
	}
	sort.Float64s(data)
	fmt.Println("binary Seach(setelah data diurutkan):")
	fmt.Println("Data terurut:", data)
	idxBin, iterBin := binarySearch(data, target)
	if idxBin != -1 {
		fmt.Printf("Hasil ditemukan di indeks %d dalam %d lankah \n", idxBin, iterBin)
	} else {

		fmt.Printf(" tidak ditemukan dalam %d lankah \n", iterBin)
	}
}
