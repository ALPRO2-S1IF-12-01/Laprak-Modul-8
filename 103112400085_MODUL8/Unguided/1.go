// Anastasia Adinda Narendra Indianto
// 103112400085
package main

import (
	"fmt"
	"sort"
)

func main() {
	var anastasiaInput int
	var totalSuara, suaraSah int
	var suara [21]int     // indeks 1-20 untuk calon
	var calonAda [21]bool // menandai calon mana saja yang mendapat suara

	for {
		fmt.Scan(&anastasiaInput)

		if anastasiaInput == 0 {
			break
		}

		totalSuara++

		if anastasiaInput >= 1 && anastasiaInput <= 20 {
			suara[anastasiaInput]++
			suaraSah++
			calonAda[anastasiaInput] = true
		}
	}

	fmt.Printf("Suara masuk: %d\n", totalSuara)
	fmt.Printf("Suara sah: %d\n", suaraSah)

	var daftarCalon [20]int
	idx := 0
	for i := 1; i <= 20; i++ {
		if calonAda[i] {
			daftarCalon[idx] = i
			idx++
		}
	}

	sort.Ints(daftarCalon[:idx])

	for i := 0; i < idx; i++ {
		c := daftarCalon[i]
		fmt.Printf("%d: %d\n", c, suara[c])
	}
}
