//BERTHA ADELA
//103112400041
package main
import "fmt"
func main() {
	var pilihan int
	var arr [100]int
	pilihan = -1
	jumlahMasukan := 0
	jumlahSah := 0
	var frekuensi [21]int

	for pilihan != 0 {
		fmt.Scan(&pilihan)
		if pilihan != 0 {
			arr[jumlahMasukan] = pilihan
			jumlahMasukan = jumlahMasukan + 1

			if pilihan >= 1 && pilihan <= 20 {
				frekuensi[pilihan] = frekuensi[pilihan] + 1
				jumlahSah = jumlahSah + 1
			}
		}
	}
	fmt.Printf("Suara masuk: %d\n", jumlahMasukan)
	fmt.Printf("Suara sah: %d\n", jumlahSah)
	
	for i := 1; i <= 20; i++ {
		frekuensi[i] = frekuensi[i]
	}

	max1, max2 := 0, 0
	ketua, wakil := 0, 0

	for i := 1; i <= 20; i++ {
		if frekuensi[i] > max1 {
			max2 = max1
			wakil = ketua

			max1 = frekuensi[i]
			ketua = i
		} else if frekuensi[i] > max2 {
			max2 = frekuensi[i]
			wakil = i
		}
	}

	if ketua != 0 {
		fmt.Printf("Ketua RT: %d\n", ketua)
	}
	if wakil != 0 {
		fmt.Printf("Wakil Ketua: %d\n", wakil)
	}
}
