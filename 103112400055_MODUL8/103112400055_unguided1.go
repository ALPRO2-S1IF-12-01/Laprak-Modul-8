//Feros Pedrosa

package main

import "fmt"

func main() {
	var input int
	votes := make([]int, 21)

	fmt.Println("Masukkan suara (akhiri dengan 0):")
	count := 0
	validVotes := 0

	for {
		fmt.Scan(&input)
		if input == 0 {
			break
		}
		if input >= 1 && input <= 20 {
			votes[input]++
			validVotes++
		}

		count++
	}

	fmt.Printf("Suara masuk: %d\n", count)
	fmt.Printf("Suara sah: %d\n", validVotes)

	for i := 1; i <= 20; i++ {
		if votes[i] > 0 {
			fmt.Printf("%d: %d\n", i, votes[i])
		}
	}
}
