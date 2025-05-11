// Felix Pedrosa V

package main

import (
	"fmt"
	"sort"
)

type Kandidat struct {
	Nomor int
	Suara int
}

func main() {
	var suara int
	jumlahMasukan := 0
	jumlahSah := 0

	hitungSuara := make(map[int]int)

	for {
		_, err := fmt.Scan(&suara)
		if err != nil {
			break
		}
		if suara == 0 {
			break
		}
		jumlahMasukan++

		if suara >= 1 && suara <= 20 {
			hitungSuara[suara]++
			jumlahSah++
		}
	}

	// Buat slice dari hasil untuk disortir
	var hasil []Kandidat
	for nomor, jumlah := range hitungSuara {
		hasil = append(hasil, Kandidat{Nomor: nomor, Suara: jumlah})
	}

	// Urutkan berdasarkan suara terbanyak, jika sama urutkan berdasarkan nomor terkecil
	sort.Slice(hasil, func(i, j int) bool {
		if hasil[i].Suara == hasil[j].Suara {
			return hasil[i].Nomor < hasil[j].Nomor
		}
		return hasil[i].Suara > hasil[j].Suara
	})

	// Tampilkan hasil
	fmt.Printf("Suara masuk: %d\n", jumlahMasukan)
	fmt.Printf("Suara sah: %d\n", jumlahSah)

	if len(hasil) == 0 {
		fmt.Println("Tidak ada suara sah, tidak ada ketua dan wakil.")
	} else if len(hasil) == 1 {
		fmt.Printf("Ketua RT: %d\n", hasil[0].Nomor)
		fmt.Println("Wakil Ketua: Tidak tersedia (hanya satu calon sah).")
	} else {
		fmt.Printf("Ketua RT: %d\n", hasil[0].Nomor)
		fmt.Printf("Wakil Ketua: %d\n", hasil[1].Nomor)
	}
}
