package main

import "fmt"

func main() {
	const MAX_CANDIDATE = 20
	var suaraMasuk, suaraSah int
	var suara [MAX_CANDIDATE + 1]int
	var input int

	for {
		fmt.Scan(&input)
		if input == 0 {
			break
		}
		suaraMasuk++
		if input >= 1 && input <= MAX_CANDIDATE {
			suara[input]++
			suaraSah++
		}
	}

	fmt.Printf("Suara masuk: %d\n", suaraMasuk)
	fmt.Printf("Suara sah: %d\n", suaraSah)

	for i := 1; i <= MAX_CANDIDATE; i++ {
		if suara[i] > 0 {
			fmt.Printf("%d: %d\n", i, suara[i])
		}
	}

	ketua, wakil := 0, 0
	max1, max2 := -1, -1

	for i := 1; i <= MAX_CANDIDATE; i++ {
		if suara[i] > max1 {
			max2 = max1
			wakil = ketua
			max1 = suara[i]
			ketua = i
		} else if suara[i] == max1 && i < ketua {
			max2 = max1
			wakil = ketua
			ketua = i
		} else if suara[i] > max2 && suara[i] < max1 {
			max2 = suara[i]
			wakil = i
		} else if suara[i] == max2 && i < wakil {
			wakil = i
		}
	}

	fmt.Printf("Ketua RT: %d\n", ketua)
	fmt.Printf("Wakil ketua: %d\n", wakil)
}
