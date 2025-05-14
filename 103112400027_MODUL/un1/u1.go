//RAJA_103112400027
package main

import (
    "fmt"
)

const N = 1000
const MAX_CALON = 20

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

    fmt.Println("Suara masuk:", totalMasuk)
    fmt.Println("Suara sah:", totalSah)
    for i := 1; i <= MAX_CALON; i++ {
        if jumlahSuara[i] > 0 {
            fmt.Printf("%d: %d\n", i, jumlahSuara[i])
        }
    }
}
