// M.HANIF AL FAIZ
// 103112400042
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Calon struct {
	Nomor int
	Suara int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	parts := strings.Fields(input)

	var (
		totalSuara int
		suaraValid int
		suaraCalon = make(map[int]int)
	)

	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			continue
		}

		totalSuara++

		if num == 0 {
			break
		}

		if num >= 1 && num <= 20 {
			suaraValid++
			suaraCalon[num]++
		}
	}

	var calons []Calon
	for nomor, suara := range suaraCalon {
		calons = append(calons, Calon{Nomor: nomor, Suara: suara})
	}

	sort.Slice(calons, func(i, j int) bool {
		if calons[i].Suara == calons[j].Suara {
			return calons[i].Nomor < calons[j].Nomor
		}
		return calons[i].Suara > calons[j].Suara
	})

	fmt.Printf("Suara masuk: %d\n", totalSuara)
	fmt.Printf("Suara sah: %d\n", suaraValid)

	if len(calons) == 0 {
		fmt.Println("Tidak ada calon yang mendapatkan suara")
	} else if len(calons) == 1 {
		fmt.Printf("Ketua RT: %d\n", calons[0].Nomor)
		fmt.Println("Wakil ketua: Tidak ada")
	} else {
		fmt.Printf("Ketua RT: %d\n", calons[0].Nomor)
		fmt.Printf("Wakil ketua: %d\n", calons[1].Nomor)
	}
}
