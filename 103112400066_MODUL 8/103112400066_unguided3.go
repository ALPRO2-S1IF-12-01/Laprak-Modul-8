//  DWI OKTA SURYANINGRUM
//  103112400066

package main

import (
	"fmt"
)

const NMAX = 1000000

// arrau global untuk menyimpan data
var data [NMAX]int


func main() {
	var n, k int

	// membaca n (jumlah data) dan k (angka yang ingin dicari)
	fmt.Scan(&n, &k)

	// memanggil fungsi isiArray untuk mengisi array data sebanyak n angka
	isiArray(n)

	// mencari posisi k dalam array
	pos := posisi(n, k)

	// mwnampilkan hasil pencarian
	if pos == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(pos)
	}
}

// fungsi unruk membaca n data dan menyimpannya ke dalam array
func isiArray(n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i]) // Baca satu per satu dan simpan ke array
	}
}

// fungsi untuk mencari posisi k dalam array data
func posisi(n, k int) int {
	// kareba data sudah terurut, kita bisa pakai binary search
	low := 0
	high := n - 1

	for low <= high {
		mid := (low + high) / 2

		if data[mid] == k {
			return mid // Ketemu, langsung return posisi
		} else if data[mid] < k {
			low = mid + 1 // Cari di kanan
		} else {
			high = mid - 1 // Cari di kiri
		}
	}

	return -1 // Kalau sampai sini, berarti tidak ketemu
}
