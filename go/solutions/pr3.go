//The prime factors of 13195 are 5, 7, 13 and 29.
//
//What is the largest prime factor of the number 600851475143 ?

// Articles used in this algorithm:
// https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
// https://en.wikipedia.org/wiki/Trial_division

package solutions

import (
	"fmt"
	"math"
)

func sieve(n int) []int {
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

func trial_division(n int) []int {
	var result []int
	if n < 2 {
		return result
	}
	
	new_n := int(math.Pow(float64(n), 0.5) + 1)
	for _, v := range(sieve(new_n)) {
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

func Pr3() {
	fmt.Println(sieve(30))
	fmt.Println(trial_division(792))
	fmt.Println(trial_division(600851475143))
}
