//RAJA_103112400027
package main

import (
    "fmt"
)

const MAX_CALON = 20	
const N = 1000

func main() {
    var suara [N]int
    var jumlahSuara [MAX_CALON + 1]int
    var totalMasuk, totalSah int
    var input int

    for {
        fmt.Scan(&input)
        if input == 0 {
            break
        }
        suara[totalMasuk] = input
        totalMasuk++
    }

    for i := 0; i < totalMasuk; i++ {
        if suara[i] >= 1 && suara[i] <= MAX_CALON {
            jumlahSuara[suara[i]]++
            totalSah++
        }
    }

    ketua, wakil := -1, -1
    max1, max2 := 0, 0

    for i := 1; i <= MAX_CALON; i++ {
        if jumlahSuara[i] > max1 {
            max2 = max1
            wakil = ketua
            max1 = jumlahSuara[i]
            ketua = i
        } else if jumlahSuara[i] == max1 && ketua != -1 && i < ketua {
            wakil = ketua
            ketua = i
        } else if jumlahSuara[i] > max2 && jumlahSuara[i] < max1 {
            max2 = jumlahSuara[i]
            wakil = i
        } else if jumlahSuara[i] == max2 && i < wakil {
            wakil = i
        }
    }

    fmt.Println("Suara masuk:", totalMasuk)
    fmt.Println("Suara sah:", totalSah)
    fmt.Println("Ketua RT:", ketua)
    fmt.Println("Wakil ketua:", wakil)
}
