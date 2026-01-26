package arrays

import "fmt"

func LargestNum() {

	var arr = []int{2, 4, 100, 3, 0}
	max := arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	fmt.Println("Largest Num = ", max)
}
