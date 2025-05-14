//Ahmad Ruba'i
//103112400074
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

    fmt.Printf("Suara masuk: %d\n", suaraMasuk)
    fmt.Printf("Suara sah: %d\n", suaraSah)
    for i := 1; i <= 20; i++ {
        if suara[i] > 0 {
            fmt.Printf("%d: %d\n", i, suara[i])
        }
    }
}