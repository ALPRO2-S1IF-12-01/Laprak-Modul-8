package main

import (
	"fmt"
	"sort"
)

func sequentialSearch(arr []float64, target float64) (int,int) {
	iterations := 0
	for i, val := range arr {
		iterations++
		fmt.Printf("Sequential Step %d: cek arr[%d] = %.1f\n", iterations, i, val)
		if val == target {
			return i, iterations
		}
	}
	return -1, iterations
}

func binarySearch(arr []float64, target float64) (int,int) {
	iterations := 0
	low := 0
	high := len(arr) - 1

	for low <= high {
		iterations++
		mid := (low + high) / 2
		fmt.Printf("Binary Step %d: cek arr[%d] = %.1f\n", iterations, mid, arr[mid])
		if arr[mid] == target {
			return mid, iterations
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1, iterations
}

func main() {
	data :=[]float64{2, 7, 9, 1, 5, 6, 18, 13, 25, 20}
	target := 13.0

	fmt.Println("Sequential Search (data tidak perlu urut):")
	idxSeq, iterSeq := sequentialSearch(data, target)
	if idxSeq != -1 {
		fmt.Printf("Hasil: Ditemukan pada index %d dalam %d langkah\n\n", idxSeq, iterSeq)
	} else {
		fmt.Printf("Hasil: Tidak ditemukan setelah %d langkah\n", iterSeq)
	}

	sort.Float64s(data)
	fmt.Println("\nBinary Search (setelah data di urutkan):")
	fmt.Println("Data urut:", data)

	idxBin, iterBin := binarySearch(data, target)
	if idxBin != -1 {
		fmt.Printf("Hasil: Ditemukan pada index %d dalam %d langkah\n", idxBin, iterBin)
	} else {
		fmt.Printf("Hasil: Tidak ditemukan setelah %d langkah\n", iterBin)
	}
	
}