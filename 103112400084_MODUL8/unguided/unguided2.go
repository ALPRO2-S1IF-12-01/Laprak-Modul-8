package main
// 103112400084
// NUFAIL ALAUDDIN TSAQIF
import "fmt"

func pilkart() {
	var target, suaraMasuk, suaraSah int
	jumlahVote := make(map[int]int) 

	for {
		_, err := fmt.Scan(&target)
		if err != nil || target == 0 {
			break
		}
		if target >= 1 && target <= 20 {
			suaraMasuk++
			suaraSah++
			jumlahVote[target]++ 
		} else {
			suaraMasuk++ 
		}
	}

	fmt.Println("Suara masuk:", suaraMasuk)
	fmt.Println("Suara sah:", suaraSah)

	var ketua, wakilKetua int
	maxVote, secondMaxVote := 0, 0

	for calon, vote := range jumlahVote {
		if vote > maxVote {
			secondMaxVote = maxVote
			wakilKetua = ketua
			maxVote = vote
			ketua = calon
		} else if vote > secondMaxVote {
			secondMaxVote = vote
			wakilKetua = calon
		}
	}

	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakilKetua)
}

func main() {
	pilkart()
}
