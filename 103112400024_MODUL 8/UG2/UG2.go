package main
// 103112400024 Setyo Nugroho
import "fmt"

func pilkart() {
	var input int
	var totalMasuk, totalSah int
	var suara [21]int

	for {
		fmt.Scan(&input)
		if input == 0 {
			break
		}
		totalMasuk++
		if input >= 1 && input <= 20 {
			suara[input]++
			totalSah++
		}
	}

	fmt.Printf("Suara masuk: %d\n", totalMasuk)
	fmt.Printf("Suara sah: %d\n", totalSah)

	ketua := 0
	maxSuara := 0
	for i := 1; i <= 20; i++ {
		if suara[i] > maxSuara || (suara[i] == maxSuara && i < ketua) {
			ketua = i
			maxSuara = suara[i]
		}
	}

	wakil := 0
	secondMax := 0
	for i := 1; i <= 20; i++ {
		if i != ketua {
			if suara[i] > secondMax || (suara[i] == secondMax && (wakil == 0 || i < wakil)) {
				wakil = i
				secondMax = suara[i]
			}
		}
	}

	fmt.Printf("Ketua RT: %d\n", ketua)
	fmt.Printf("Wakil ketua: %d\n", wakil)
}

func main() {
	pilkart()
}