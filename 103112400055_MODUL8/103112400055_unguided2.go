//Feros Pedrosa

package main

import "fmt"

func main() {
	var input int
	votes := make([]int, 21) // Indeks 1-20 untuk nomor calon

	fmt.Println("Masukkan suara (akhiri dengan 0):")
	count := 0
	validVotes := 0

	for {
		fmt.Scan(&input)
		if input == 0 {
			break
		}
		count++
		if input >= 1 && input <= 20 {
			votes[input]++
			validVotes++
		}
	}

	ketua, suaraKetua := 0, 0
	wakil, suaraWakil := 0, 0

	for i := 1; i <= 20; i++ {
		suaraSaatIni := votes[i]
		if suaraSaatIni == 0 {
			continue
		}

		if suaraSaatIni > suaraKetua {
			wakil, suaraWakil = ketua, suaraKetua
			ketua, suaraKetua = i, suaraSaatIni
		} else if suaraSaatIni == suaraKetua {
			if i < ketua {
				wakil, suaraWakil = ketua, suaraKetua
				ketua, suaraKetua = i, suaraSaatIni
			} else if suaraWakil < suaraSaatIni || (suaraWakil == suaraSaatIni && i < wakil) {
				wakil, suaraWakil = i, suaraSaatIni
			}
		} else if suaraSaatIni > suaraWakil {
			wakil, suaraWakil = i, suaraSaatIni
		} else if suaraSaatIni == suaraWakil && i < wakil {
			wakil, suaraWakil = i, suaraSaatIni
		}
	}

	fmt.Printf("Suara masuk: %d\n", count)
	fmt.Printf("Suara sah: %d\n", validVotes)
	fmt.Printf("Ketua RT: %d\n", ketua)
	fmt.Printf("Wakil ketua: %d\n", wakil)
}
