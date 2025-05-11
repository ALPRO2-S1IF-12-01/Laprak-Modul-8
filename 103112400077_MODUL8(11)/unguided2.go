package main

import "fmt"

func main() {
	data := []int{7, 19, 3, 2, 78, 3, -1, -3, 18, 19, 0}

	input := 0
	valid := 0
	var hitungSuara [21]int 

	for _, suara := range data {
		if suara == 0 {
			break
		}
		input++
		if suara >= 1 && suara <= 20 {
			valid++
			hitungSuara[suara]++
		}
	}

	max1, max2 := 0, 0
	ketua, wakil := 0, 0

	for i := 1; i <= 20; i++ {
		jumlah := hitungSuara[i]
		if jumlah > max1 {
			max2 = max1
			wakil = ketua

			max1 = jumlah
			ketua = i
		} else if jumlah == max1 && i < ketua {
			max2 = max1
			wakil = ketua

			ketua = i
		} else if jumlah > max2 && jumlah < max1 {
			max2 = jumlah
			wakil = i
		} else if jumlah == max2 && jumlah != 0 && i < wakil {
			wakil = i
		}
	}

	fmt.Printf("Suara masuk: %d\n", input)
	fmt.Printf("Suara sah: %d\n", valid)
	fmt.Printf("Ketua RT: %d\n", ketua)
	fmt.Printf("Wakil Ketua: %d\n", wakil)
}
