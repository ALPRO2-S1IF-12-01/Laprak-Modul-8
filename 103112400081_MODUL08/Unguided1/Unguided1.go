//RYANAKEYLANOVIANTOWIDODO
//103112400081
//MAAFCOPASLAPTOPNGEHANK

package main

import (
	"fmt"
	"sort"
)

func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	validCandidates := make([]int, 20)
	for i := range validCandidates {
		validCandidates[i] = i + 1
	}
	sort.Ints(validCandidates)

	voteCounts := make(map[int]int)
	totalVotes := 0
	var input int

	fmt.Println("Masukkan suara satu per satu (akhiri dengan 0):")
	for {
		_, err := fmt.Scan(&input)
		if err != nil || input == 0 {
			break
		}
		totalVotes++
		index := binarySearch(validCandidates, input)
		if index != -1 {
			voteCounts[input]++
		}
	}

	fmt.Printf("Suara masuk: %d\n", totalVotes)
	fmt.Printf("Suara sah: %d\n", len(voteCounts))

	for candidate, count := range voteCounts {
		fmt.Printf("%d: %d\n", candidate, count)
	}
}
