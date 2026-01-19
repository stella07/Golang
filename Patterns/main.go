package main

import "fmt"

func main() {
	fmt.Println("Pattern Programs")
	var N = 5
	Pattern1(N)
}

/*
SQUARE FILL PATTERN

*****
*****
*****
*****
*****

*/
func Pattern1(N int) {
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}

}
