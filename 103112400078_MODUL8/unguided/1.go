package main
//103112400078 MOHAMMAD REYHAN ARETHA FATIN 
import "fmt"

func pilkart() {
	var target int
	var suaraMasuk, suaraSah int = 0, 0
	jumlahVote := make(map[int]int)

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
			jumlahVote[target]++
		}
	}

	fmt.Printf("\nSuara masuk: %d\n", suaraMasuk)
	fmt.Printf("Suara sah: %d\n", suaraSah)

	fmt.Println("\nHasil Perhitungan Suara:")
	for calon, suara := range jumlahVote {
		fmt.Printf("%d: %d\n", calon, suara)
	}
}

func main() {
	pilkart()
}
