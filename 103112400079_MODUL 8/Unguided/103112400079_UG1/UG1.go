package main

// Muhammad Faris Rachmadi
// 103112400079
import "fmt"

func pilkart() {
	var target int
	var suaraMasuk, suaraSah int = 0, 0
	var jumlahVote [21]int // Array untuk menyimpan suara calon 1-20

	for {
		fmt.Scan(&target)
		if target == 0 {
			break
		}
		if target >= 1 && target <= 20 {
			suaraSah++
			jumlahVote[target]++
		}
		suaraMasuk++
	}

	fmt.Printf("Suara masuk: %d\n", suaraMasuk)
	fmt.Printf("Suara sah: %d\n", suaraSah)

	for i := 1; i <= 20; i++ {
		if jumlahVote[i] > 0 {
			fmt.Printf("%d: %d\n", i, jumlahVote[i])
		}
	}
}

func main() {
	pilkart()
}
