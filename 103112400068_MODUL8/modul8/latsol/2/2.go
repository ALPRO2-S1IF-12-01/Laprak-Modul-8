package main

import "fmt"

func main() {
	var angka float64
	var temp int
	fmt.Scan(&angka)
	temp = int(angka) + 1
	cek := false
	for !cek {
		angka = angka + 0.1
		fmt.Printf("%.1f\n", angka)
		cek = angka > float64(temp)-0.11
	}
	fmt.Println(temp)
}
