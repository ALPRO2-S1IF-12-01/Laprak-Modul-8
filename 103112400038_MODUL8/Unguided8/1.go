// Nama  : Hakan Ismail Afnan
// NIM   : 103112400038
package main

import (
	"fmt"
	"sort"
)

func binarySearch(arr []int, target int) bool {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return true
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}

func main() {
	const maxCalon = 20
	data := []int{7, 19, 3, 2, 78, 3, 1, 3, 18, 19, 0} // data input langsung

	// Daftar calon sudah terurut 1..20
	validCalon := make([]int, maxCalon)
	for i := 1; i <= maxCalon; i++ {
		validCalon[i-1] = i
	}
	sort.Ints(validCalon) // sebenarnya sudah terurut, tapi untuk konsistensi

	countSuara := make(map[int]int)
	totalData := 0
	totalValid := 0

	for _, v := range data {
		if v == 0 {
			break
		}
		totalData++
		if binarySearch(validCalon, v) {
			countSuara[v]++
			totalValid++
		}
	}

	fmt.Printf("Suara masuk: %d\n\n", totalData)
	fmt.Printf("Suara sah: %d\n\n", totalValid)

	for i := 1; i <= maxCalon; i++ {
		if count, ok := countSuara[i]; ok {
			fmt.Printf("%d:%d\n", i, count)
		}
	}
}