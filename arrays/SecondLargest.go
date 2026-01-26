package arrays

import (
	"fmt"
	"math"
)

var arr = []int{1, 1, 2, 2, 2, 3, 3}

func SecondSmallestLargest() {
	SecondLargest()
}

func SecondLargest() {
	if len(arr) < 2 {
		fmt.Println("Second max cannot be found")
		return
	}
	largest := math.MinInt64
	secondLargest := math.MinInt64

	for i := 0; i < len(arr); i++ {
		if arr[i] > largest {
			secondLargest = largest
			largest = arr[i]
		} else if arr[i] > secondLargest && arr[i] != largest {
			secondLargest = arr[i]
		}
	}
	fmt.Println("Second largest = ", secondLargest)
}
