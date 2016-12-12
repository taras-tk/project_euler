// Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
// If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable numbers.

// For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284. The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.

// Evaluate the sum of all the amicable numbers under 10000.

// todos@linux-ew09:~/tmp/project_euler/go/src> go run main.go
// 220 [1 2 4 5 10 11 20 22 44 55 110] 284
// 284 [1 2 4 71 142] 220
// map[5:1 6:6 7:1 8:7 9:4 1:0 2:1 3:1 4:3 10:8]
// 284
// 220
// map[6368:6232 220:284 284:220 2620:2924 5564:5020 5020:5564 6232:6368 1210:1184 1184:1210 2924:2620]
// 31626

package solutions

import (
	"fmt"
)

func properDivisors(n int) (result []int) {
	for i := 1; i <= n; i++ {
		if n%i == 0 && i != n {
			result = append(result, i)
		}
	}
	return
}

func sumElements(elements []int) (sum int) {
	for _, el := range elements {
		sum += el
	}
	return
}

func buildElementMap(maxNum int) (result map[int]int) {
	result = make(map[int]int)
	for el := 1; el <= maxNum; el++ {
		result[el] = sumElements(properDivisors(el))
	}
	return
}

func findAmicableNumbers(input map[int]int) (result map[int]int) {
	result = make(map[int]int)
	for i, el := range input {
		// fmt.Println(i, el)
		if input[el] == i && el != i {
			result[i] = el
			result[el] = i
		}
	}
	return
}

func findAmicableSum(input map[int]int) (result int) {
	var amicables []int
	for i, el := range input {
		// fmt.Println(i, el)
		if input[el] == i && el != i {
			amicables = append(amicables, el)
		}
	}
	// fmt.Println(amicables)
	for _, el := range amicables {
		result += el
	}
	return
}

func Pr21() {
	fmt.Println(220, properDivisors(220), sumElements(properDivisors(220)))
	fmt.Println(284, properDivisors(284), sumElements(properDivisors(284)))
	fmt.Println(buildElementMap(10))
	fmt.Println(buildElementMap(220)[220])
	fmt.Println(buildElementMap(284)[284])
	fmt.Println(findAmicableNumbers(buildElementMap(10000)))
	fmt.Println(findAmicableSum(buildElementMap(10000)))
}
