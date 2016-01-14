//A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
//
//a^2 + b^2 = c^2
//For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.
//
//There exists exactly one Pythagorean triplet for which a + b + c = 1000.
//Find the product abc.

package main

import (
	"fmt"
)

func find_pythagorean_triplet(max_sum int) (a int, b int, c int) {
	for i := 1; i < max_sum; i++ {
		for j := i; j < max_sum; j++ {
			for k := j; k < max_sum; k++ {
				if is_pythagorean_triplet(i, j, k) {
					if i + j + k == max_sum {
						a, b, c = i, j, k
						return
					}
				}
			}
		}
	}
	return
}

func is_pythagorean_triplet(a int, b int, c int) bool {
	if a > b && a > c && b > c {
		fmt.Printf("Input error: %v, %v, %v\n", a, b, c)
		return false
	}
	
	return a*a + b*b == c*c
}

func main() {
	fmt.Printf("%v\n", is_pythagorean_triplet(3,4,4))
	a, b, c := find_pythagorean_triplet(12)
	fmt.Printf("%v, %v, %v\n", a, b, c)
	a, b, c = find_pythagorean_triplet(1000)
	fmt.Printf("%v * %v * %v = %v\n", a, b, c, a*b*c)
}
