// Savila Nur Fadilla
// 103112400031
package main

import (
	"fmt"
	"sort"
)

func main() {
	var input int
	suaraMasuk := 0
	suaraSah := 0
	suara := make(map[int]int)

	for {
		fmt.Scan(&input)
		if input == 0 {
			break
		}

		suaraMasuk++

		if input >= 1 && input <= 20 {
			suara[input]++
			suaraSah++
		}
	}
	fmt.Println("Suara masuk:", suaraMasuk)
	fmt.Println("Suara sah:", suaraSah)

	var calon []int
	for k := range suara {
		calon = append(calon, k)
	}
	sort.Ints(calon)

	for _, c := range calon {
		fmt.Printf("%v: %v\n", c, suara[c])
	}
}