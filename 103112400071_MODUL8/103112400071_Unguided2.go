// Raihan Adi Arba
// 103112400071

package main

import "fmt"

func main() {
	const N = 20
	var (
		suara      [N + 1]int
		input      int
		totalSuara int
		suaraSah   int
		ketua      int
		wakil      int
	)

	for {
		fmt.Scan(&input)
		if input == 0 {
			break
		}
		totalSuara++

		if input >= 1 && input <= N {
			suara[input]++
			suaraSah++
		}
	}

	fmt.Printf("Suara masuk: %d\n", totalSuara)
	fmt.Printf("Suara sah: %d\n", suaraSah)

	for i := 1; i <= N; i++ {
		if suara[i] > suara[ketua] || (suara[i] == suara[ketua] && i < ketua) {
			wakil = ketua
			ketua = i
		} else if suara[i] > suara[wakil] || (suara[i] == suara[wakil] && i < wakil && i != ketua) {
			wakil = i
		}
	}

	fmt.Printf("Ketua RT: %d\n", ketua)
	fmt.Printf("Wakil ketua: %d\n", wakil)
}
