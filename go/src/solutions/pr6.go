//The sum of the squares of the first ten natural numbers is,
//
//1^2 + 2^2 + ... + 10^2 = 385
//The square of the sum of the first ten natural numbers is,
//
//(1 + 2 + ... + 10)^2 = 55^2 = 3025
//Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.
//
//Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

package solutions

import (
	"fmt"
)

func sum_of_squares(seq []int) (res int) {
	for _, v := range(seq) {
		res += v * v
	}
	return
}

func square_of_sum(seq []int) (res int) {
	for _, v := range(seq) {
		res += v
	}
	res *= res
	return
}

func Pr6() {
	var arr []int
	for i := 1; i <= 100; i++ {
		arr = append(arr, i)
	}
	//fmt.Printf("%v\n", arr)

	s1 := sum_of_squares(arr)
	s2 := square_of_sum(arr)
	fmt.Printf("%v - %v = %v\n", s2, s1, s2 - s1)
}
