package main

import "fmt"

const MAKS = 1000000
var data [MAKS]int

func isiArray(n int) {
    for i := 0; i < n; i++ {
        fmt.Scan(&data[i])
    }
}

func cariPosisi(n, x int) int {
    kiri := 0
    kanan := n - 1
    
    for kiri <= kanan {
        tengah := (kiri + kanan) / 2
        if data[tengah] == x {
            return tengah
        } else if data[tengah] < x {
            kiri = tengah + 1
        } else {
            kanan = tengah - 1
        }
    }
    return -1
}

func main() {
    var n, x int
    fmt.Scan(&n, &x)
    
    isiArray(n)
    pos := cariPosisi(n, x)
    
    if pos != -1 {
        fmt.Println(pos)
    } else {
        fmt.Println("TIDAK ADA")
    }
}