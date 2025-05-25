package main

//103112400050

import "fmt"

func main() {
	var nomorPilihan int
	suaraMasuk := 0
	suaraSah := 0
	hitungSuaraCalon := make([]int, 21)
	for {
		_, err := fmt.Scan(&nomorPilihan)
		if err != nil || nomorPilihan == 0 {
			break
		}
		suaraMasuk++
		if nomorPilihan >= 1 && nomorPilihan <= 20 {
			suaraSah++
			hitungSuaraCalon[nomorPilihan]++
		}
	}
	fmt.Printf("Suara masuk: %d\n", suaraMasuk)
	fmt.Printf("Suara sah: %d\n", suaraSah)
	for i := 1; i <= 20; i++ {
		if hitungSuaraCalon[i] > 0 {
			fmt.Printf("%d: %d\n", i, hitungSuaraCalon[i])
		}
	}
}
