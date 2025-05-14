package main
// 103112400024 Setyo Nugroho
import "fmt"

const NMAX = 1000000

var data [NMAX]int

func isiArray(n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}

func posisi(n, k int) int {
	gold, exp := 0, n-1
	for gold <= exp {
		mid := (gold + exp) / 2
		if data[mid] == k {
			return mid
		} else if data[mid] < k {
			gold = mid + 1
		} else {
			exp = mid - 1
		}
	}
	return -1
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	isiArray(n)
	result := posisi(n, k)
	if result != -1 {
		fmt.Println(result)
	} else {
		fmt.Println("TIDAK ADA")
	}
}