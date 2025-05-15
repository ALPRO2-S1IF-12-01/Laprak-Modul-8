package main

import "fmt"

func main() {
    var input int
    masuk := 0
    sah := 0
    hasil := make(map[int]int)

    for {
        fmt.Scan(&input)
        masuk++

        if input == 0 {
            masuk-- 
            break
        }

        if input >= 1 && input <= 20 {
            sah++
            hasil[input]++
        }
    }

    ketua, wakil := cariPemenang(hasil)

    fmt.Printf("Suara masuk: %d\n", masuk)
    fmt.Printf("Suara sah: %d\n", sah)
    fmt.Printf("Ketua RT: %d\n", ketua)
    fmt.Printf("Wakil ketua: %d\n", wakil)
}

func cariPemenang(hasil map[int]int) (int, int) {
    max1, max2 := 0, 0  
    ketua, wakil := 0, 0

    for nomor, suara := range hasil {
        if suara > max1 {
            max2 = max1
            wakil = ketua
            max1 = suara
            ketua = nomor
        } else if suara > max2 && suara < max1 {
            max2 = suara
            wakil = nomor
        } else if suara == max1 {
            if nomor < ketua {
                max2 = max1
                wakil = ketua
                max1 = suara
                ketua = nomor
            } else if nomor < wakil || wakil == 0 {
                wakil = nomor
            }
        } else if suara == max2 {
            if nomor < wakil || wakil == 0 {
                wakil = nomor
            }
        }
    }

    return ketua, wakil
}