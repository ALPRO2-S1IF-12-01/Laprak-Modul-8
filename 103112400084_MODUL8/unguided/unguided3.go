package main
// 103112400084
// NUFAIL ALAUDDIN TSAQIF
import "fmt"

const NMAX = 10000

var data [NMAX]int

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	isiArray(n)

	letak := posisi(n, k)
	if letak == -1 {
		fmt.Print("TIDAK ADA")
	} else {
		fmt.Print(letak)
	}
}

func isiArray(a int) {
	for i := 0; i < a; i++ {
		fmt.Scan(&data[i])
	}
}

func posisi(a, b int) int {
	low := 0
	high := a - 1
	for low <= high {
		mid := (low + high) / 2
		if data[mid] == b {
			return mid
		} else if b < data[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
