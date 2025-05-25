package main

//103112400050
import (
	"fmt"
)

func main() {
	const MaxCalon = 20
	var suara [MaxCalon + 1]int
	var input, total, sah int

	for {
		fmt.Scan(&input)
		if input == 0 {
			break
		}
		total++
		if input >= 1 && input <= MaxCalon {
			suara[input]++
			sah++
		}
	}

	fmt.Printf("Suara masuk: %d\n", total)
	fmt.Printf("Suara sah: %d\n", sah)

	var ketua, wakil int
	for i := 1; i <= MaxCalon; i++ {
		if suara[i] > suara[ketua] || (suara[i] == suara[ketua] && i < ketua) {
			wakil = ketua
			ketua = i
		} else if (suara[i] > suara[wakil] || (suara[i] == suara[wakil] && i < wakil)) && i != ketua {
			wakil = i
		}
	}

	fmt.Printf("Ketua RT: %d\n", ketua)
	fmt.Printf("Wakil ketua: %d\n", wakil)
}
