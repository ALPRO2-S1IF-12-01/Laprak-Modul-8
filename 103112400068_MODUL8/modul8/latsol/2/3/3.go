package main

import "fmt"

func main() {
	const MAX_CANDIDATE = 20
	var suaraMasuk, suaraSah int
	var suara [MAX_CANDIDATE + 1]int
	var input int

	for {
		fmt.Scan(&input)
		if input == 0 {
			break
		}
		suaraMasuk++
		if input >= 1 && input <= MAX_CANDIDATE {
			suara[input]++
			suaraSah++
		}
	}

	fmt.Printf("Suara masuk: %d\n", suaraMasuk)
	fmt.Printf("Suara sah: %d\n", suaraSah)

	for i := 1; i <= MAX_CANDIDATE; i++ {
		if suara[i] > 0 {
			fmt.Printf("%d: %d\n", i, suara[i])
		}
	}
}
