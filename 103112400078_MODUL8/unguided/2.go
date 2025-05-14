package main
// 103112400078 MOHAMMAD REYHAN ARETHA FATIN
import "fmt"
func pilkart() {
	var target, suaraMasuk, suaraSah int
	jumlahVote := make([]int, 20) 
	for {
		_, err := fmt.Scan(&target)
		if err != nil || target == 0 {
			break
		}
		if target >= 1 && target <= 20 {
			suaraMasuk++
			suaraSah++
			jumlahVote[target-1]++
		} else {
			suaraMasuk++
		}
	}

	fmt.Println("Suara masuk:", suaraMasuk)
	fmt.Println("Suara sah:", suaraSah)

	var ketua, wakilKetua int
	maxVote, secondMaxVote := 0, 0

	for i, vote := range jumlahVote {
		if vote > maxVote {
			secondMaxVote = maxVote
			wakilKetua = ketua
			maxVote = vote
			ketua = i + 1
		} else if vote > secondMaxVote {
			secondMaxVote = vote
			wakilKetua = i + 1
		}
	}

	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakilKetua)
}

func main() {
	pilkart()
}
