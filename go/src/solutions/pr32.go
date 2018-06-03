// We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly once; for example, the 5-digit number, 15234, is 1 through 5 pandigital.
//
// The product 7254 is unusual, as the identity, 39 × 186 = 7254, containing multiplicand, multiplier, and product is 1 through 9 pandigital.
//
// Find the sum of all products whose multiplicand/multiplier/product identity can be written as a 1 through 9 pandigital.
// HINT: Some products can be obtained in more than one way so be sure to only include it once in your sum.

package solutions

import (
	"fmt"
	"strconv"
)

func swap(a []rune, i int, j int) {
	c := a[i]
	a[i] = a[j]
	a[j] = c
}

// This function takes four parameters:
// 1. Array of runes
// 2. Starting index of the array
// 3. Ending index of the array
// 4. Result array
func permute(a []rune, l int, r int, result *[][]int) {
	if l == r {
		// fmt.Printf("- %+v\n", string(a))
		findId(string(a), result)
	} else {
		for i := l; i <= r; i++ {
			swap(a, l, i)
			permute(a, l+1, r, result)
			swap(a, l, i) //backtrack
		}
	}
}

func findId(str string, result *[][]int) {
	// fmt.Printf("- %s\n", str)
	for i := 1; i < len(str)-1; i++ {
		for k := i + 1; k < len(str); k++ {
			left, _ := strconv.Atoi(str[0:i])
			right, _ := strconv.Atoi(str[i:k])
			ident, _ := strconv.Atoi(str[k:])
			// fmt.Println(left, right, ident)
			if left*right == ident {
				*result = append(*result, []int{left, right, ident})
			}
		}
	}
	// fmt.Println(*result)
	return
}

func findProducts(in [][]int) (result []int) {
	for _, x := range in {
		if len(result) == 0 {
			result = append(result, x[2])
		} else {
			included := false
			for _, r := range result {
				if r == x[2] {
					included = true
					break
				}
			}
			if !included {
				result = append(result, x[2])
			}
		}
	}
	return
}

func findSum(in []int) (result int) {
	for _, x := range in {
		result += x
	}
	return
}

func Pr32() {
	var str string = "123456789"
	var n int = len(str)
	var result [][]int
	permute([]rune(str), 0, n-1, &result)
	fmt.Println(result)
	fmt.Println("---")
	prod := findProducts(result)
	fmt.Println(prod)
	fmt.Println("---")
	fmt.Println(findSum(prod))
}

// ➜  src git:(master) ✗ go run main.go
// [[12 483 5796] [138 42 5796] [157 28 4396] [159 48 7632] [1738 4 6952] [186 39 7254] [18 297 5346] [1963 4 7852] [198 27 5346] [27 198 5346] [28 157 4396] [297 18 5346] [39 186 7254] [42 138 5796] [4 1738 6952] [4 1963 7852] [483 12 5796] [48 159 7632]]
// ---
// [5796 4396 7632 6952 7254 5346 7852]
// ---
// 45228
