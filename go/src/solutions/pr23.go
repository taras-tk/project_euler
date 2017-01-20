// A perfect number is a number for which the sum of its proper divisors is exactly equal to the number. For example, the sum of the proper divisors of 28 would be 1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect number.
//
// A number n is called deficient if the sum of its proper divisors is less than n and it is called abundant if this sum exceeds n.
//
// As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest number that can be written as the sum of two abundant numbers is 24. By mathematical analysis, it can be shown that all integers greater than 28123 can be written as the sum of two abundant numbers. However, this upper limit cannot be reduced any further by analysis even though it is known that the greatest number that cannot be expressed as the sum of two abundant numbers is less than this limit.
//
// Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers.

package solutions

import (
	"fmt"
)

// >>> def proper_divs2(n):
// ...     return {x for x in range(1, (n + 1) // 2 + 1) if n % x == 0 and n != x}

func proper_division(n int) (result []int) {
	for x := 1; x < ((n+1)/2 + 1); x++ {
		if n%x == 0 && n != x {
			result = append(result, x)
		}
	}
	return result
}

func sumPDivs(n int) (result []int) {
	result = append(result, 1)
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			result = append(result, i)
		}
	}
	return
}

// >>> def divisors(n):
//     l=[]
//     for i in range(1, n//2 +1):
//         if n % i==0:
//             l.append(i)
//     l.append(n)
//     return(l)
// >>> divisors(360)

func divisors(n int) (result []int) {
	for i := 1; i < (n/2 + 1); i++ {
		if n%i == 0 {
			result = append(result, i)
		}
	}
	return
}

func divisors2(n int) (result []int) {
	for i := 1; i < n; i++ {
		if n%i == 0 && n != i {
			result = append(result, i)
		}
	}
	return
}

func get_perfect_numbers(n int) (result []int) {
	for i := 1; i < n; i++ {
		sum_divs := 0
		divs := proper_division(i)
		for j := 0; j < len(divs); j++ {
			sum_divs += divs[j]
		}
		if sum_divs == i {
			result = append(result, i)
		}
	}
	return
}

func get_abundant_numbers(n int) (result []int) {
	for i := 1; i <= n; i++ {
		sum_divs := 0
		divs := proper_division(i)
		for j := 0; j < len(divs); j++ {
			sum_divs += divs[j]
		}
		if sum_divs > i {
			result = append(result, i)
		}
	}
	return
}

func isEqualToSumOfTwo(numbers []int, value int) (result bool) {
	for j := 0; j < len(numbers); j++ {
		if numbers[j] > value {
			return
		}

		for k := 0; k < len(numbers); k++ {
			sum := numbers[j] + numbers[k]
			if sum > value {
				break
			}

			if sum == value {
				// fmt.Printf("Eq S2: %d + %d = %d\n", numbers[j], numbers[k], sum)
				result = true
				return
			}
		}
	}
	return
}

func calc_result(n int) (result int64) {
	numbers := get_abundant_numbers(n)
	// fmt.Println(numbers)
	for i := 0; i <= n; i++ {
		if !isEqualToSumOfTwo(numbers, i) {
			result += int64(i)
			// fmt.Printf("Add: %d\n", i)
		}
	}
	return
}

func Pr23() {
	fmt.Println(proper_division(16))
	fmt.Println(sumPDivs(16))
	fmt.Println(divisors(16))
	fmt.Println(divisors2(16))
	fmt.Println(proper_division(12))
	fmt.Println(proper_division(28))

	//fmt.Println(get_perfect_numbers(28123))
	fmt.Println(calc_result(50))
	fmt.Println(calc_result(28123))
}

// todos@linux-ew09:~/tmp/project_euler/go> time go run src/main.go
// [1 2 4 8]
// [1 2 4 8]
// [1 2 4 8]
// [1 2 4 8]
// [1 2 3 4 6]
// [1 2 4 7 14]
// 891
// 4179871
//
// real    0m10.087s
// user    0m10.049s
// sys     0m0.041s
