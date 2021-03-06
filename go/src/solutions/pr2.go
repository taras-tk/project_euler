//Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:
//
//1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
//
//By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.

package solutions

import (
	"fmt"
)

var fibo_arr []uint64

func fibo(val uint64, max_val uint64) uint64 {
	if val > max_val {
		return 0
	}

	var res uint64
	if val == 1 {
		res = 1
	} else if val == 2 {
		res = 2
	} else {
		res = fibo_arr[val - 2] + fibo_arr[val - 3]
	}

	if res >= 4000000 {
		return 0
	}

	fibo_arr = append(fibo_arr, res)
	return fibo(val + 1, max_val)
}

func Pr2() {
	fibo(1, 100)
	fmt.Println("res ", fibo_arr)

	var sum uint64
	for _,x := range fibo_arr {
		if x % 2 == 0 {
			sum = sum + x
			fmt.Println("even ", x)
		}
	}
	fmt.Println("sum ", sum)
}
