package main

import "fmt"

func main() {
	var suara int
	hitungSuara := make(map[int]int)
	totalSuaraMasuk := 0
	totalSuaraSah := 0
	totalSuaraTidakSah := 0

	fmt.Println("Program Penghitungan Suara Pemilihan")
	fmt.Println("==================================")

	fmt.Println("\nMasukkan nomor calon (1-20)")
	fmt.Println("Masukkan 0 untuk mengakhiri input")

	for {
		fmt.Print("\nNomor calon: ")
		_, err := fmt.Scan(&suara)
		if err != nil {
			fmt.Println("Input tidak valid! Masukkan angka saja.")
			var clearInput string
			fmt.Scanln(&clearInput) 
			continue
		}

		if suara == 0 {
			break
		}

		totalSuaraMasuk++

		if suara >= 1 && suara <= 20 {
			hitungSuara[suara]++
			totalSuaraSah++
		} else {
			totalSuaraTidakSah++
			fmt.Println("Nomor calon tidak valid! (1-20)")
		}
	}

	fmt.Println("\nHasil Penghitungan Suara")
	fmt.Println("=======================")
	fmt.Printf("Total suara masuk: %d\n", totalSuaraMasuk)
	fmt.Printf("Total suara sah: %d\n", totalSuaraSah)
	fmt.Printf("Total suara tidak sah: %d\n", totalSuaraTidakSah)

	ketua, wakil := 0, 0
	max1, max2 := 0, 0

	for calon, jumlah := range hitungSuara {
		if jumlah > max1 {
			max2 = max1
			wakil = ketua
			max1 = jumlah
			ketua = calon
		} else if jumlah > max2 {
			max2 = jumlah
			wakil = calon
		}
	}

	if totalSuaraSah > 0 {
		fmt.Println("\nPerolehan Suara per Calon:")
		for i := 1; i <= 20; i++ {
			if hitungSuara[i] > 0 {
				fmt.Printf("Calon %2d: %d suara\n", i, hitungSuara[i])
			}
		}

		fmt.Println("\nHasil Akhir:")
		fmt.Printf("Ketua RT terpilih: Calon %d (%d suara)\n", ketua, max1)
		fmt.Printf("Wakil Ketua terpilih: Calon %d (%d suara)\n", wakil, max2)
	} else {
		fmt.Println("\nTidak ada suara sah yang masuk")
	}
}
