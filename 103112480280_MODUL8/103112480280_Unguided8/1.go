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
	totalData := 0
	totalValid := 0

	fmt.Println("Masukkan data suara (akhiri dengan 0):")
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

	fmt.Printf("Suara masuk: %d\n\n", totalData)
	fmt.Printf("Suara sah: %d\n\n", totalValid)
	for i, count := range suara {
		if count > 0 {
			fmt.Printf("%d:%d\n", i+1, count)
		}
	}
}