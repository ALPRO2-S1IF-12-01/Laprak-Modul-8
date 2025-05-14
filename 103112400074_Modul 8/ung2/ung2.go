// Ahmad Ruba'i
// 103112400074
package main

import "fmt"

func main() {
    var suaraMasuk, suaraSah int
    var suara [21]int 
    var x int

    selesai := false
    for !selesai {
        n, _ := fmt.Scan(&x)
        if n == 0 {
            selesai = true
        } else if x == 0 {
            selesai = true
        } else {
            suaraMasuk++
            if x >= 1 && x <= 20 {
                suaraSah++
                suara[x]++
            }
        }
    }

    fmt.Println("Suara masuk:", suaraMasuk)
    fmt.Println("Suara sah:", suaraSah)

    max := 0
    for i := 1; i <= 20; i++ {
        if suara[i] > max {
            max = suara[i]
        }
    }

    var pemenang []int
    for i := 1; i <= 20; i++ {
        if suara[i] == max && max > 0 {
            pemenang = append(pemenang, i)
        }
    }

    ketua := 0
    if len(pemenang) > 0 {
        ketua = pemenang[0]
    }

    wakil := 0
    if len(pemenang) > 1 {
        wakil = pemenang[1]
    } else {
        max2 := 0
        for i := 1; i <= 20; i++ {
            if suara[i] > max2 && suara[i] < max {
                max2 = suara[i]
            }
        }
        for i := 1; i <= 20; i++ {
            if suara[i] == max2 && max2 > 0 {
                wakil = i
                break
            }
        }
    }

    fmt.Println("Ketua RT:", ketua)
    fmt.Println("Wakil ketua RT:", wakil)
}