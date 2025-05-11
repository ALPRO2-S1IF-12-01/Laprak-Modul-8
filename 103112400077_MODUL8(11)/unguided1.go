package main

import "fmt"

func main() {
	data := []int{7, 19, 3, 2, 78, 3, 1, -3, 18, 19, 0}

	input := 0
	valid := 0
	var hitungSuara [21]int

	for _, suara := range data {
		if suara == 0 {
			break
		}
		input++
		if suara >= 1 && suara <= 20 {
			valid++
			hitungSuara[suara]++
		}
	}

	fmt.Printf("Suara masuk: %d\n", input)
	fmt.Printf("Suara sah: %d\n", valid)

	for calon := 1; calon <= 20; calon++ {
		if hitungSuara[calon] > 0 {
			fmt.Printf("%d: %d\n", calon, hitungSuara[calon])
		}
	}
}
