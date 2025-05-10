package main

import (
	"fmt"
	"sort"
)

func binarySearch(arr []int, target int) bool {
	low := 0
	high := len(arr) - 1
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
	var input int
	validCandidates := make([]int, 20)
	for i := 0; i < 20; i++ {
		validCandidates[i] = i + 1
	}
	sort.Ints(validCandidates)

	voteCounts := make([]int, 21) 
	totalVotes := 0
	validVotes := 0

	fmt.Println("Masukkan suara satu per satu (akhiri dengan 0):")
	for {
		_, err := fmt.Scan(&input)
		if err != nil {
			break
		}
		if input == 0 {
			break
		}
		totalVotes++
		if binarySearch(validCandidates, input) {
			voteCounts[input]++
			validVotes++
		}
	}

	fmt.Printf("Suara masuk: %d\n", totalVotes)
	fmt.Printf("Suara sah: %d\n", validVotes)

	for i := 1; i <= 20; i++ {
		if voteCounts[i] > 0 {
			fmt.Printf("%d: %d\n", i, voteCounts[i])
		}
	}
}
