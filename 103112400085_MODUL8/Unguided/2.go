// Anastasia Adinda Narendra Indrianto
// 103112400085
package main

import (
	"fmt"
	"sort"
)

func main() {
	var anastasiaInput int
	var totalSuara, suaraSah int
	var suara [21]int

	for {
		fmt.Scan(&anastasiaInput)
		if anastasiaInput == 0 {
			break
		}

		totalSuara++

		if anastasiaInput >= 1 && anastasiaInput <= 20 {
			suara[anastasiaInput]++
			suaraSah++
		}
	}

	fmt.Printf("Suara masuk: %d\n", totalSuara)
	fmt.Printf("Suara sah: %d\n", suaraSah)

	type calon struct {
		nomor, suara int
	}

	var calonSuara []calon
	for i := 1; i <= 20; i++ {
		if suara[i] > 0 {
			calonSuara = append(calonSuara, calon{nomor: i, suara: suara[i]})
		}
	}

	sort.Slice(calonSuara, func(i, j int) bool {
		if calonSuara[i].suara == calonSuara[j].suara {
			return calonSuara[i].nomor < calonSuara[j].nomor
		}
		return calonSuara[i].suara > calonSuara[j].suara
	})

	if len(calonSuara) > 1 {
		ketua := calonSuara[0]
		wakilKetua := calonSuara[1]
		fmt.Printf("Ketua RT: %d\n", ketua.nomor)
		fmt.Printf("Wakil ketua: %d\n", wakilKetua.nomor)
	} else if len(calonSuara) == 1 {
		ketua := calonSuara[0]
		fmt.Printf("Ketua RT: %d\n", ketua.nomor)
		fmt.Printf("Wakil ketua: Tidak ada calon lain\n")
	} else {
		fmt.Println("Tidak ada calon yang sah")
	}
}
