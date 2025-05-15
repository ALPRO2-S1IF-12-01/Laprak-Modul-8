// M. DAVI ILYAS RENALDO_103112400062
package main

import (
	"fmt"
)

func main() {
	var suara [21]int 
	var angka int
	suaraMasuk := 0
	suaraSah := 0

	for {
		_, err := fmt.Scan(&angka)
		if err != nil || angka == 0 {
			break
		}

		suaraMasuk++

		if angka >= 1 && angka <= 20 {
			suara[angka]++
			suaraSah++
		}
	}
	fmt.Printf("Suara masuk: %d\n", suaraMasuk)
	fmt.Printf("Suara sah: %d\n", suaraSah)

	for i := 1; i <= 20; i++ {
		if suara[i] > 0 {
			fmt.Printf("%d: %d\n", i, suara[i])
		}
	}
}