package main

import "fmt"

func main() {
	var s, total, sah int
	suara := make(map[int]int)

	for {
		fmt.Scan(&s)
		if s == 0 {
			break
		}
		total++
		if s >= 1 && s <= 20 {
			sah++
			suara[s]++
		}
	}

	fmt.Println("Suara masuk:", total)
	fmt.Println("Suara sah:", sah)

	var ketua, wakil int
	max1, max2 := 0, 0

	for i := 1; i <= 20; i++ {
		skrg := suara[i]
		if skrg > max1 {
			max2 = max1
			max1 = skrg
			ketua = i
		} else if skrg > max2 && i != ketua {
			max2 = skrg
			wakil = i
		}
	}

	fmt.Printf("Ketua RT: %d\n", ketua)
	if max2 > 0 {
		fmt.Printf("Wakil: %d\n", wakil)
	} else {
		fmt.Println("Wakil: Tidak ada")
	}
}