package main
import (
	"fmt"
	"sort"

)

func sequentialSearch(arr []float64, target float64) (int, int) {
	itterations := 0
	for i, val := range arr {
		itterations++
		fmt.Printf("Sequential Step %d: cek arr[%d] = %.1f\n", itterations, i, val)
		if val == target {
			return i, itterations
		}
	}
	return -1, itterations
}

func binarySearch(arr []float64, target float64) (int, int) {
	itterations := 0
	low := 0
	high := len(arr) - 1
	for low <= high {
		itterations++
		mid := (low + high) / 2
		fmt.Printf("Binary Step %d: cek arr[%d] = %.1f\n", itterations, mid, arr[mid])
		if arr[mid] == target {
			return mid, itterations
		} else if target < arr[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
			}
			return -1, itterations
	}
func main() {

	//Array awal
	data := []float64{2, 7, 9, 1, 5, 6, 18, 13,25,20}
	target := 13.0

	fmt.Println("Sequential Search (data tidak perlu urut):")
	idxSeq, iterSeq := sequentialSearch(data, target)
	if idxSeq != -1 {
		fmt.Printf("Hasil : Ditemukan di indeks %d dalam %d langkah\n\n", idxSeq, iterSeq)
	} else {
		fmt.Printf("Hasil : Tidak ditemukan dalam %d langkah\n", iterSeq)
	}

	//Binary search perlu array diurutkan
	sort.Float64s(data)
	fmt.Println("Binary Search (setelah data diurutkan):")
	fmt.Println("Data terurut:", data)

	idxBin, iterBin := binarySearch(data, target)
	if idxBin != -1 {
		fmt.Printf("Hasil : Ditemukan di indeks %d dalam %d langkah\n", idxBin, iterBin)
	} else {
		fmt.Printf("Hasil : Tidak ditemukan dalam %d langkah\n", iterBin)
	}

	
}