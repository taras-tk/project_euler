package solutions

// Surprisingly there are only three numbers that can be written as the sum of fourth powers of their digits:

// 1634 = 14 + 64 + 34 + 44
// 8208 = 84 + 24 + 04 + 84
// 9474 = 94 + 44 + 74 + 44

// As 1 = 14 is not a sum it is not included.

// The sum of these numbers is 1634 + 8208 + 9474 = 19316.

// Find the sum of all the numbers that can be written as the sum of fifth powers of their digits.

import (
	"fmt"
	"math"
	"strconv"
)

func sumNthPowerFor(number int, n int) (result int) {
	strNumber := strconv.Itoa(number)
	// fmt.Println(strNumber)
	for _, digitStr := range strNumber {
		// fmt.Printf("%c\n", digitStr)
		digitInt, _ := strconv.Atoi(string(digitStr))
		// fmt.Printf("%d", digitInt)
		result += int(math.Pow(float64(digitInt), float64(n)))
	}
	return
}

func bruteSearchSums(maxNum int, n int) (result int) {
	for i := 2; i <= maxNum; i++ {
		if i == sumNthPowerFor(i, n) {
			fmt.Println(i)
			result += i
		}
	}
	return
}

func Pr30() {
	// fmt.Println(sumNthPowerFor(1634, 4))
	// fmt.Println(sumNthPowerFor(8208, 4))
	// fmt.Println(sumNthPowerFor(9474, 4))
	result := bruteSearchSums(999999, 4)
	fmt.Printf("result(4)=%d\n", result)
	result = bruteSearchSums(9999999, 5)
	fmt.Printf("result(5)=%d\n", result)
}

// Tarass-MBP-3:src todos$ time go run main.go
// 1634
// 8208
// 9474
// result(4)=19316
// 4150
// 4151
// 54748
// 92727
// 93084
// 194979
// result(5)=443839

// real	0m7.061s
// user	0m7.018s
// sys	0m0.141s
