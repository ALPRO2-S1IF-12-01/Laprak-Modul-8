//Muhammad Gamel Al Ghifari
//103112400028
package main

import (
	"fmt"
	"sort"
)

func main() {
	var x, total, sah int
	v := make(map[int]int)

	for fmt.Scan(&x); x != 0; fmt.Scan(&x) {
		total++
		if x >= 1 && x <= 20 {
			v[x]++
			sah++
		}
	}

	type pair struct{ no, jml int }
	var data []pair
	for i := 1; i <= 20; i++ {
		if v[i] > 0 {
			data = append(data, pair{i, v[i]})
		}
	}

	sort.Slice(data, func(i, j int) bool {
		if data[i].jml == data[j].jml {
			return data[i].no < data[j].no
		}
		return data[i].jml > data[j].jml
	})

	fmt.Println("Suara masuk:", total)
	fmt.Println("Suara sah:", sah)
	if len(data) > 0 {
		fmt.Println("Ketua RT:", data[0].no)
	}
	if len(data) > 1 {
		fmt.Println("Wakil ketua:", data[1].no)
	}
}