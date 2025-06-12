package main

import "fmt"

func main() {
	fmt.Println("Program Penghitungan Suara Pemilihan\n==================================")
	fmt.Println("Masukkan nomor calon (1-20) atau 0 untuk mengakhiri input.")

	var suara int
	hitungSuara := [21]int{}
	totalSuaraMasuk := 0
	totalSuaraSah := 0

	for {
		fmt.Print("\nNomor calon: ")
		_, err := fmt.Scan(&suara)
		if err != nil {
			fmt.Println("Input tidak valid! Masukkan angka saja.")
			var dummy string
			fmt.Scanln(&dummy) 
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
			fmt.Println("Nomor calon tidak valid! (1-20)")
		}
	}

	fmt.Println("\nHasil Penghitungan Suara\n=======================")
	fmt.Printf("Total suara masuk: %d\n", totalSuaraMasuk)
	fmt.Printf("Total suara sah: %d\n", totalSuaraSah)
	fmt.Printf("Total suara tidak sah: %d\n", totalSuaraMasuk-totalSuaraSah)

	ketua, wakil := 0, 0
	max1, max2 := 0, 0

	for i := 1; i <= 20; i++ {
		if hitungSuara[i] > max1 {
			max2 = max1
			wakil = ketua
			max1 = hitungSuara[i]
			ketua = i
		} else if hitungSuara[i] > max2 {
			max2 = hitungSuara[i]
			wakil = i
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