package main
// 103112400084
// NUFAIL ALAUDDIN TSAQIF
import "fmt"

func pilkart() {
	var target int
	var suaraMasuk, suaraSah int
	jumlahVote := make([]int, 20) 

	for {
		_, err := fmt.Scan(&target)
		if err != nil || target < 0 {
			continue
		}
		if target == 0 {
			break
		}
		suaraMasuk++
		if target >= 1 && target <= 20 {
			suaraSah++
			jumlahVote[target-1]++ 
		}
	}

	fmt.Printf("\nSuara masuk: %d\n", suaraMasuk)
	fmt.Printf("Suara sah: %d\n", suaraSah)

	fmt.Println("\nHasil Perhitungan Suara:")
	for i, suara := range jumlahVote {
		if suara > 0 { 
			fmt.Printf("Calon %d: %d suara\n", i+1, suara)
		}
	}
}

func main() {
	pilkart()
}
