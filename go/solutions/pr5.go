// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
// 
// 	What is the smallest positive number that is evenly divisible (divisible with no remainder) by all of the numbers from 1 to 20?

// Resolution:
// https://en.wikipedia.org/wiki/Least_common_multiple
// The lcm will be the product of multiplying the highest power of each prime number together. 
// Here I use Sieve to calculate the largest prime.

package solutions

import (
	"fmt"
	"math"
)

func sieve2(n int) []int {
	var arr []bool = make([]bool, n)
	var res []int

	for i,_ := range arr {
		arr[i] = true
	}

	var sqrt = int(math.Sqrt(float64(n)))

	for i := 2; i <= sqrt; i = i+1 {
		if arr[i] == true {
			q := 0
			for j := i*i; j < n; j = i*i + q*i {
				q = q+1
				arr[j] = false
			}
		}
	}
	
	for i, v := range arr {
		if v == true && i != 0 && i != 1 {
			res = append(res, i)
		}
	}

	return res
}

func trial_division2(n int) []int {
	var result []int
	if n < 2 {
		return result
	}
	
	new_n := int(math.Pow(float64(n), 0.5) + 1)
	for _, v := range(sieve2(new_n)) {
		if v * v > n {
			break
		}
		for ; n % v == 0; {
			result = append(result, v)
			n = n / v
		}
	}
	if n > 1 {
		result = append(result, n)
	}
	return result
}

func prime_powers(n int) (result map[int]int) {
	result = make(map[int]int)
	seq := trial_division2(n)
	for _, val := range(seq) {
		_, ok := result[val]
		if !ok {
			result[val] = 1
		} else {
			result[val] = result[val] + 1
		}
	}
	return
}

func Pr5() {
	seq := []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	var res map[int]int
	res = make(map[int]int)

	for _, v := range(seq) {
		pow := prime_powers(v)
		for key, val := range(pow) {
			v, ok := res[key]
			if !ok {
				res[key] = val
			} else {
				if val > v {
					res[key] = val
				}
			}
		}
		
		fmt.Printf("%v => %v => %v\n", v, trial_division2(v), prime_powers(v))
	}

	fmt.Printf("%v\n", res)

	num := 1
	for key, val := range(res) {
		num = num * int(math.Pow(float64(key), float64(val)))
	}

	fmt.Printf("num => %v\n", num)
}
