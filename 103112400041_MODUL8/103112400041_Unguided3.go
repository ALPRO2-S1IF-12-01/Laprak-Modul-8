//BERTHA ADELA
//103112400041

package main
import "fmt"
const NMAX = 1000000
var data [NMAX]int
func main(){
	var n, k int
	fmt.Scanln(&n, &k)
	isiArray(n)
	posisi(n, k)
	if posisi(n, k) != -1 {
		fmt.Println(posisi(n,k))
	} else {
		fmt.Println("TIDAK ADA")
	}
}
func isiArray(n int){
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}
func posisi(n, k int) int {
	low := 0
	high := n-1

	for low <= high {
		mid := (low + high)/2
		if data[mid] == k {
			return mid
		} else if k < data[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}