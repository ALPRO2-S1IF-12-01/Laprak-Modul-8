// M. DAVI ILYAS RENALDO_103112400062
package main

import (
	"fmt"
	"sort"
)

type Pilkart struct {
	totalSuara   int
	suaraSah     int
	hasilPilih map[int]int
}

func prosesSuara(data []int) (Pilkart, int) {
	p := Pilkart{
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

func tentukanPemenang(p Pilkart) (int, int) {
	type Calon struct {
		nomor int
		suara int
	}
	
	var noCalon []Calon
	
	for nomor, suara := range p.hasilPilih {
		if suara > 0 {
			noCalon = append(noCalon, Calon{nomor, suara})
		}
	}
	
	sort.Slice(noCalon, func(i, j int) bool {
		if noCalon[i].suara == noCalon[j].suara {
			return noCalon[i].nomor < noCalon[j].nomor
		}
		return noCalon[i].suara > noCalon[j].suara
	})

	var ketua, wakil int
	
	if len(noCalon) > 0 {
		ketua = noCalon[0].nomor
	}
	
	if len(noCalon) > 1 {
		wakil = noCalon[1].nomor
	}
	
	return ketua, wakil
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

	hasilPungutan, _ := prosesSuara(data)

	fmt.Printf("\nSuara masuk: %d\n", hasilPungutan.totalSuara)
	fmt.Printf("Suara sah: %d\n", hasilPungutan.suaraSah)
	
	ketua, wakil := tentukanPemenang(hasilPungutan)
	fmt.Printf("Ketua RT: %d\n", ketua)
	fmt.Printf("Wakil ketua: %d\n", wakil)
}