package main

import "fmt"

func main() {
	var N = 5
	Pattern1(N)
}

/*
Square Fill Pattern

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

/*
Right Half Pyramid Pattern

*
**
***
****
*****

*/

func Pattern2(N int) {
	for i := 1; i <= N; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

/*
 Reverse Right Half Pyramid Pattern

*****
****
***
**
*

*/

func Pattern3(N int) {
	for i := 1; i <= N; i++ {
		for j := 1; j <= N-i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

/*
Number Increasing Pyramid Pattern

1
12
123
1234
12345

*/
func Pattern4(N int) {
	for i := 1; i <= N; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		fmt.Println("")
	}
}

/*
Number-Increasing Reverse Pyramid Pattern

12345
1234
123
12
1

*/
func Pattern5(N int) {
	for i := 1; i <= N; i++ {
		for j := 1; j <= N-i+1; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
}

/*
Right Pascal’s Triangle

*
**
***
****
*****
****
***
**
*

*/
func Pattern6(N int) {
	Pattern2(N)
	Pattern3(N - 1)
}

/*
K Pattern

*****
****
***
**
*
**
***
****
*****


*/
func Pattern7(N int) {
	for i := 1; i <= N; i++ {
		for j := 1; j <= N-i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := 1; i <= N-1; i++ {
		for j := 1; j <= i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/*
Zero-One Triangle Pattern

1
0 1
1 0 1
0 1 0 1
1 0 1 0 1

*/
func Pattern8(N int) {
	for i := 1; i <= N; i++ {
		for j := 1; j <= i; j++ {
			if (i+j)%2 == 0 {
				fmt.Print("1 ")
			} else {
				fmt.Print("0 ")
			}

		}
		fmt.Println("")
	}
}

/*
Square Hollow Pattern

******
*    *
*    *
*    *
*    *
******

*/
func Pattern9(N int) {
	for i := 0; i <= N; i++ {
		for j := 0; j <= N; j++ {
			if i == 0 || j == 0 || j == N || i == N {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}

/*

Number-Changing Pyramid Pattern

1
2 3
4 5 6
7 8 9 10
11 12 13 14 15

*/
func Pattern10(N int) {
	num := 1
	for i := 1; i <= N; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(num, " ")
			num++
		}
		fmt.Println("")
	}
}

/*
Left Half Pyramid Pattern

    *
   **
  ***
 ****
*****


*/
func Pattern11(N int) {
	for i := 1; i <= N; i++ {
		for j := 1; j <= N-i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/*
Reverse Left Half Pyramid Pattern

*****
 ****
  ***
   **
    *

*/
func Pattern12(N int) {
	for i := 1; i <= N; i++ {
		for j := 1; j <= i-1; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= N-i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

/*
Triangle Star Pattern

    *
   ***
  *****
 *******
*********

*/
func Pattern13(N int) {
	for i := 1; i <= N; i++ {
		for j := 1; j <= N-i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/*

Number Triangle Pattern

    1
   2 2
  3 3 3
 4 4 4 4
5 5 5 5 5

*/
func Pattern14(N int) {
	for i := 1; i <= N; i++ {
		for j := 1; j <= N-i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print(i, " ")
		}
		fmt.Println("")
	}
}

/*
Reverse Number Triangle Pattern

1 2 3 4 5
 2 3 4 5
  3 4 5
   4 5
    5

*/
func Pattern15(N int) {
	for i := 1; i <= N; i++ {
		for j := 1; j <= i-1; j++ {
			fmt.Print(" ")
		}
		for j := i; j <= N; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}

/*
Rhombus Pattern

    *****
   *****
  *****
 *****
*****

*/
func Pattern16(N int) {
	for i := 1; i <= N; i++ {
		for j := 1; j <= N-i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= N; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/*
Palindrome Triangle Pattern

   	1
   2 1 2
  3 2 1 2 3
 4 3 2 1 2 3 4
5 4 3 2 1 2 3 4 5

*/
func Pattern17(N int) {
	for i := 1; i <= N; i++ {
		for j := 1; j <= N-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < i; j++ {
			fmt.Print(i-j, " ")
		}
		for j := 1; j <= i-1; j++ {
			fmt.Print(j+1, " ")
		}
		fmt.Println("")
	}
}

/*
Diamond Star Pattern
    *
   ***
  *****
 *******
*********
 *******
  *****
   ***
    *

*/
func Pattern18(N int) {
	Pattern13(N)
	for i := 1; i < N; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= (N-i)*2-1; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

/*
Butterfly Star Pattern
    *
   ***
  *****
 *******
*********
 *******
  *****
   ***
    *

*/
func Pattern19(N int) {

	for i := 1; i <= N; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		for j := 1; j <= (N-i)*2; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}

	for i := 0; i < N; i++ {
		for j := 1; j <= N-i; j++ {
			fmt.Print("*")
		}
		for j := 1; j <= i*2; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= N-i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}

}

/*
Mirror Image Triangle Pattern

1 2 3 4 5
 2 3 4 5
  3 4 5
   4 5
    5
   4 5
  3 4 5
 2 3 4 5
1 2 3 4 5

*/
func Pattern20(N int) {

	for i := 1; i <= N; i++ {
		for j := 1; j <= i-1; j++ {
			fmt.Print(" ")
		}
		for j := i; j <= N; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}

	for i := 1; i < N; i++ {
		for j := 1; j <= N-i-1; j++ {
			fmt.Print(" ")
		}
		for j := N - i; j <= N; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}

/*
Hollow Triangle Pattern
    *
   * *
  *   *
 *     *
*********

*/

func Pattern21(N int) {
	for i := 1; i <= N; i++ {
		for j := 1; j <= N-i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			if i == N || j == 1 || j == 2*i-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

/*
Hollow Reverse Triangle Pattern

*********
 *     *
  *   *
   * *
    *

*/

func Pattern22(N int) {
	for i := 1; i <= N; i++ {
		for j := 1; j <= i-1; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*(N-i+1)-1; j++ {
			if j == 2*(N-i+1)-1 || j == 1 || i == 1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

/*
Hollow Diamond Pyramid

    *
   * *
  *   *
 *     *
*       *
 *     *
  *   *
   * *
    *

*/
func Pattern23(N int) {
	for i := 1; i <= N; i++ {
		for j := 1; j <= N-i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*(i)-1; j++ {
			if j == 1 || j == 2*(i)-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Println()
	}
	for i := 1; i < N; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= (N-i)*2-1; j++ {
			if j == 1 || j == (N-i)*2-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Println()
	}
}

/*
Hollow Hourglass Pattern


* * * * *
 *     *
  *   *
   * *
    *
   * *
  *   *
 *     *
* * * * *

*/
func Pattern24(N int) {
	for i := 1; i <= N; i++ {
		for j := 1; j <= i-1; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= N-i+1; j++ {
			if i == 1 || j == 1 || j == N-i+1 {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}

		}
		fmt.Println()
	}

	for i := 1; i < N; i++ {
		for j := 1; j <= N-i-1; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i+1; j++ {
			if i == N-1 || j == 1 || j == i+1 {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}

		}
		fmt.Println()
	}
}

/*
X - Star Pattern

*       *
 *     *
  *   *
   * *
    *
   * *
  *   *
 *     *
*       *

*/
func Pattern25(N int) {
	for i := 1; i <= N; i++ {
		for j := 1; j <= i-1; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= N-i+1; j++ {
			if j == 1 || j == N-i+1 {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}

		}
		fmt.Println()
	}

	for i := 1; i < N; i++ {
		for j := 1; j <= N-i-1; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i+1; j++ {
			if j == 1 || j == i+1 {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}

		}
		fmt.Println()
	}
}
