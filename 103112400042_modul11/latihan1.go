// M.HANIF AL FAIZ
// NIM: 103112400042
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

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

	var calonTerpilih []int
	for calon := range suaraCalon {
		calonTerpilih = append(calonTerpilih, calon)
	}
	sort.Ints(calonTerpilih)

	fmt.Printf("Suara masuk: %d\n", totalSuara)
	fmt.Printf("Suara sah: %d\n", suaraValid)

	for _, calon := range calonTerpilih {
		fmt.Printf("%d: %d\n", calon, suaraCalon[calon])
	}
}
