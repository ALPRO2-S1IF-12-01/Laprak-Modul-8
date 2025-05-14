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

	for i := 1; i <= 20; i++ {
		if suara[i] > 0 {
			fmt.Printf("%d: %d\n", i, suara[i])
		}
	}
}

func main() {
	pilkart()
}