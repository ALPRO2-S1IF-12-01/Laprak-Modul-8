// Achmad Zulvan Nur Hakim 103112400070
package main

import "fmt"

type PilihanKaRT struct {
	totalSuara   int
	suaraSah     int
	hasilPilih map[int]int
}

func prosesSuara(data []int) (PilihanKaRT, int) {
	p := PilihanKaRT{
		totalSuara:   0,
		suaraSah:     0,
		hasilPilih: make(map[int]int),
	}

	for i := 1; i <= 20; i++ {
		p.hasilPilih[i] = 0
	}

	iterasi := 0

	for _, suara := range data {
		iterasi++
		p.totalSuara++

		if suara == 0 {
			p.totalSuara--
			break
		}

		if suara >= 1 && suara <= 20 {
			p.suaraSah++
			p.hasilPilih[suara]++
		}
	}

	return p, iterasi
}

func main() {	
	var input int
	var data []int
	
	for {
		fmt.Scan(&input)
		data = append(data, input)
		if input == 0 {
			break
		}
	}

	hasilPungutan, iterasi := prosesSuara(data)

	fmt.Printf("\nTotal iterasi: %d\n", iterasi)
	fmt.Printf("Suara masuk: %d\n", hasilPungutan.totalSuara)
	fmt.Printf("Suara sah: %d\n", hasilPungutan.suaraSah)
	
	fmt.Println("Hasil:")
	for i := 1; i <= 20; i++ {
		if hasilPungutan.hasilPilih[i] > 0 {
			fmt.Printf("%d: %d suara\n", i, hasilPungutan.hasilPilih[i])
		}
	}
}