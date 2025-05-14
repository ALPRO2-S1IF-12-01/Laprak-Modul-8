package main

import "fmt"

func main() {
	var num, total, sah int
	suara := [21]int{}

	for {
		_, err := fmt.Scan(&num)
		if err != nil || num == 0 {
			break
		}
		total++
		if num >= 1 && num <= 20 {
			sah++
			suara[num]++
		}
	}

	ketua, wakil := 0, 0
	max1, max2 := 0, 0

	for i := 1; i <= 20; i++ {
		if suara[i] > max1 {
			max2 = max1
			wakil = ketua
			max1 = suara[i]
			ketua = i
		} else if suara[i] > max2 {
			max2 = suara[i]
			wakil = i
		}
	}

	fmt.Printf("suara yg masuk: %d\nsuara yg sah: %d\n", total, sah)
	fmt.Printf("Ketua RT: %d\nWakil ketua: %d\n", ketua, wakil)
}