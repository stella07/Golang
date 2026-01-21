package arrays

import "fmt"

func LargestNum() {
	fmt.Println("Find the largest element in the given array")

	var arr = [5]int{1, 5, 7, 2, 3}
	var largestNum = arr[0]
	for i := 0; i <= len(arr)-1; i++ {
		if arr[i] > largestNum {
			largestNum = arr[i]
		}
	}

	fmt.Println(largestNum)
}
