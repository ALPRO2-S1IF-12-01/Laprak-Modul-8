package main
// 103112400078 MOHAMMAD REYHAN ARETHA FATIN
import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	data := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}

	sortArray(data)

	letak := binarySearch(data, k)
	if letak == -1 {
		fmt.Print("TIDAK ADA")
	} else {
		fmt.Print(letak)
	}
}

func sortArray(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func binarySearch(arr []int, x int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == x {
			return mid
		} else if arr[mid] < x {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
