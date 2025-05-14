package main

// Muhammad Faris Rachmadi
// 103112400079

import "fmt"

func pilkart() {
	var target int
	var suaraMasuk, suaraSah int
	var jumlahVote [21]int

	fmt.Println("Masukkan suara (akhiri dengan angka 0):")
	for {
		fmt.Scan(&target)
		if target == 0 {
			break
		}
		suaraMasuk++

		if target >= 1 && target <= 20 {
			suaraSah++
			jumlahVote[target]++
		}
	}

	fmt.Printf("Suara masuk: %d\n", suaraMasuk)
	fmt.Printf("Suara sah: %d\n", suaraSah)

	var max1, max2 int
	var calonKetua, calonWakil int

	for i := 1; i <= 20; i++ {
		if jumlahVote[i] > max1 {
			max2 = max1
			calonWakil = calonKetua
			max1 = jumlahVote[i]
			calonKetua = i
		} else if jumlahVote[i] > max2 {
			max2 = jumlahVote[i]
			calonWakil = i
		}
	}

	fmt.Printf("Ketua RT: %d\n", calonKetua)
	fmt.Printf("Wakil ketua: %d\n", calonWakil)
}

func main() {
	pilkart()
}
