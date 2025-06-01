//Nama : Anggun Wahyu Widiyana (103112480280)
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	const maxCalon = 20
	var suara [maxCalon]int
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
			if num >= 1 && num <= maxCalon {
				suara[num-1]++
				totalValid++
			}
		}
	}

	// Cari suara terbanyak (ketua)
	max1 := -1
	ketua := 0
	for i := 0; i < maxCalon; i++ {
		if suara[i] > max1 {
			max1 = suara[i]
			ketua = i + 1
		} else if suara[i] == max1 && (i+1) < ketua {
			ketua = i + 1
		}
	}

	// Cari suara terbanyak kedua (wakil), nomor terkecil dan bukan ketua
	max2 := -1
	wakil := 0
	for i := 0; i < maxCalon; i++ {
		if i+1 == ketua {
			continue
		}
		if suara[i] > max2 {
			max2 = suara[i]
			wakil = i + 1
		} else if suara[i] == max2 && (i+1) < wakil {
			wakil = i + 1
		}
	}

	fmt.Printf("Suara masuk: %d\n\n", totalData)
	fmt.Printf("Suara sah: %d\n\n", totalValid)
	fmt.Printf("Ketua RT: %d\n\n", ketua)
	fmt.Printf("Wakil ketua: %d\n", wakil)
}