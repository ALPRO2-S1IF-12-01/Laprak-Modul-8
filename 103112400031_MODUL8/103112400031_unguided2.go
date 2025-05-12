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

	type calon struct {
		nomor, jumlah int
	}
	var data []calon
	for k, v := range suara {
		data = append(data, calon{k, v})
	}

	sort.Slice(data, func(a, b int) bool {
		if data[a].jumlah == data[b].jumlah {
			return data[a].nomor < data[b].nomor
		}
		return data[a].jumlah > data[b].jumlah
	})

	if len(data) > 0 {
		fmt.Println("Ketua RT:", data[0].nomor)
	}
	if len(data) > 1 {
		fmt.Println("Wakil ketua:", data[1].nomor)
	}
}