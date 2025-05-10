// JESIKA METANIA RAHMA ARIFIN
// 103112400080

package main

import "fmt"

func main() {
	var suara [21]int
	var angka int
	jumlahMasuk := 0
	jumlahSah := 0

	for {
		fmt.Scan(&angka)
		if angka == 0 {
			break
		}
		jumlahMasuk++

		if angka >= 1 && angka <= 20 {
			suara[angka]++
			jumlahSah++
		}
	}

	fmt.Println("Suara masuk:", jumlahMasuk)
	fmt.Println("Suara sah:", jumlahSah)

	// Cari ketua (suara terbanyak) dan wakil (terbanyak kedua)
	ketua := 0
	wakil := 0
	for i := 1; i <= 20; i++ {
		if suara[i] > suara[ketua] || (suara[i] == suara[ketua] && i < ketua) {
			wakil = ketua
			ketua = i
		} else if (suara[i] > suara[wakil] || (suara[i] == suara[wakil] && i < wakil)) && i != ketua {
			wakil = i
		}
	}

	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil Ketua RT:", wakil)
}
