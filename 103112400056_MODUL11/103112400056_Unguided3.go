// Felix Pedrosa V

package main

import "fmt"

const NMAX = 1000000
var arrayData [NMAX]int

func main() {
	var jumlah, target int
	fmt.Scan(&jumlah, &target)

	isiData(jumlah)
	hasil := cariPosisi(jumlah, target)

	if hasil == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(hasil)
	}
}

func isiData(jml int) {
	for i := 0; i < jml; i++ {
		fmt.Scan(&arrayData[i])
	}
}

func cariPosisi(jml, cari int) int {
	kiri := 0
	kanan := jml - 1

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		if arrayData[tengah] == cari {
			return tengah
		} else if arrayData[tengah] < cari {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return -1
}