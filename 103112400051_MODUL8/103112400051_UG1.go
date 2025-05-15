package main

import "fmt"

func main() {
	var input int
	masuk := 0
	sah := 0
	hasil := make(map[int]int)

	for {
		fmt.Scan(&input)
		masuk++

		if input == 0 {
			masuk--
			break
		}

		if input >= 1 && input <= 20 {
			sah++
			hasil[input]++
		}
	}

	fmt.Printf("Suara masuk: %d\n", masuk)
	fmt.Printf("Suara sah: %d\n", sah)

	for i := 1; i <= 20; i++ {
		if c, e := hasil[i]; e {
			fmt.Printf("%d: %d\n", i, c)
		}
	}
}