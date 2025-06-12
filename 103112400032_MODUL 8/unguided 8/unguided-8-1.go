// daffa tsaqifna f
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const maxCalon = 20

func main() {
	var suaraMasuk, suaraSah int
	var hasil [maxCalon + 1]int // indeks 1-20 dipakai

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Masukkan suara (akhiri dengan 0):")
	scanner.Scan()
	input := scanner.Text()
	tokens := strings.Fields(input)

	for _, token := range tokens {
		num, err := strconv.Atoi(token)
		if err != nil {
			continue
		}
		if num == 0 {
			break
		}
		suaraMasuk++
		if num >= 1 && num <= maxCalon {
			suaraSah++
			hasil[num]++
		}
	}

	fmt.Printf("Suara masuk: %d\n", suaraMasuk)
	fmt.Printf("Suara sah: %d\n", suaraSah)
	for i := 1; i <= maxCalon; i++ {
		if hasil[i] > 0 {
			fmt.Printf("%d: %d\n", i, hasil[i])
		}
	}
}
