package main

import "fmt"

func main() {
	fmt.Println("Program Pencarian Elemen")
	var n, k int
	_, err := fmt.Scan(&n, &k)

	if err != nil || n <= 0 {
		fmt.Println("Input tidak valid.")
		return
	}

	data := make([]int, n)

	for i := 0; i < n; i++ {
		_, err := fmt.Scan(&data[i])
		if err != nil {
			fmt.Println("Input elemen tidak valid.")
			return
		}
	}

	pos := -1
	for i := 0; i < n; i++ {
		if data[i] == k {
			pos = i
			break
		}
	}

	if pos != -1 {
		fmt.Printf("%d\n", pos)
	} else {
		fmt.Println("TIDAK ADA")
	}
}
