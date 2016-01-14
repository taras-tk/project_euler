//A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
//
//Find the largest palindrome made from the product of two 3-digit numbers.

package main


import (
	"fmt"
	"strconv"
)

func Reverse(s string) (result string) {
  for _,v := range s {
    result = string(v) + result
  }
  return 
}

func IsPalindrome(n int) bool {
	s := strconv.Itoa(n)
	return s == Reverse(s)
}

func PalindromesBruteForce(n1 int, n1_min int, n2 int, n2_min int) (max [][]int) {
	for i := n1; i > n1_min; i = i - 1 {
		for j := n2; j > n2_min; j = j - 1 {
			product := i * j
			if IsPalindrome(product) {
				res := []int{i, j, product}
				max = append(max, res)
			}
		}
	}
	return
}

func LargestPalindromeBruteForce(n1 int, n1_min int, n2 int, n2_min int) (max int, ci int, cj int) {
	for _, v := range(PalindromesBruteForce(n1, n1_min, n2, n2_min)) {
		if v[2] > max {
			ci, cj, max = v[0], v[1], v[2]
		}
	}
	return
}

func main() {
	max, i, j := LargestPalindromeBruteForce(999, 800, 999, 800)
	fmt.Printf("%v %v %v", max, i, j)
}
