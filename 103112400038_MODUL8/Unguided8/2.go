//Nama  : Hakan Ismail Afnan 
//NIM   : 103112400038
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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
	validCalon := make([]int, maxCalon)
	for i := 1; i <= maxCalon; i++ {
		validCalon[i-1] = i
	}
	sort.Ints(validCalon)

	countSuara := make(map[int]int)
	totalData, totalValid := 0, 0

	fmt.Println("Masukkan suara (akhiri dengan 0):")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()
		data := strings.Fields(input)

		for _, v := range data {
			num, err := strconv.Atoi(v)
			if err != nil {
				continue
			}
			if num == 0 {
				break
			}
			totalData++
			if binarySearch(validCalon, num) {
				countSuara[num]++
				totalValid++
			}
		}
	}

	type calon struct {
		no, suara int
	}
	var calonList []calon
	for no, suara := range countSuara {
		calonList = append(calonList, calon{no, suara})
	}

	sort.Slice(calonList, func(i, j int) bool {
		if calonList[i].suara == calonList[j].suara {
			return calonList[i].no < calonList[j].no
		}
		return calonList[i].suara > calonList[j].suara
	})

	ketua := calonList[0].no
	wakil := calonList[1].no

	fmt.Printf("Suara masuk: %d\n\n", totalData)
	fmt.Printf("Suara sah: %d\n\n", totalValid)
	fmt.Printf("Ketua RT: %d\n\n", ketua)
	fmt.Printf("Wakil ketua: %d\n", wakil)
}