//RYANAKEYLANOVIANTOWIDODO
//103112400081
//MAAFCOPASLAPTOPNGEHANK

package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	data := make([]int, n)
	for i := range data {
		fmt.Scan(&data[i])
	}

	pos := findPosition(data, k)

	if pos == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(pos)
	}
}

func findPosition(arr []int, target int) int {
	for i, val := range arr {
		if val == target {
			return i
		}
	}
	return -1
}
