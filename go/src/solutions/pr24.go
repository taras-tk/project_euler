// A permutation is an ordered arrangement of objects. For example, 3124 is one possible permutation of the digits 1, 2, 3 and 4.
// If all of the permutations are listed numerically or alphabetically, we call it lexicographic order. The lexicographic permutations of 0, 1 and 2 are:
//
// 012   021   102   120   201   210
//
// What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?

package solutions

import (
	"fmt"
	"sort"
	"strings"
)

// This function finds the index of the smallest character
// which is greater than 'first' and is present in str[l..h]
func findCeil(str []byte, first byte, l int, h int) (ceilIndex int) {
	ceilIndex = l

	// Now iterate through rest of the elements and find
	// the smallest character greater than 'first'
	for i := l + 1; i <= h; i++ {
		if str[i] > first && str[i] < str[ceilIndex] {
			ceilIndex = i
		}
	}
	return
}

type sortBytes []byte

func (s sortBytes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortBytes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortBytes) Len() int {
	return len(s)
}

// Print all permutations of str in sorted order
func sortedPermutations(str []byte, maxPermIndex int) {

	// Get size of string
	size := len(str)

	// Sort the string in increasing order
	sort.Sort(sortBytes(str))

	// Print permutations one by one
	isFinished := false
	permIndex := 0
	for isFinished == false {

		// fmt.Println(str)

		permIndex++
		if permIndex >= maxPermIndex {
			fmt.Println("Max reached: ", str)
			result := []string{}
			for _, c := range str {
				result = append(result, string(c))
			}
			fmt.Println("Readable max: ", strings.Join(result, ""))
			break
		}

		// Find the rightmost character which is smaller than its next
		// character. Let us call it 'first char'
		var i int
		for i = size - 2; i >= 0; i-- {
			if str[i] < str[i+1] {
				break
			}
		}

		// If there is no such chracter, all are sorted in decreasing order,
		// means we just printed the last permutation and we are done.
		if i == -1 {
			isFinished = true
			fmt.Println("Done")
		} else {

			// Find the ceil of 'first char' in right of first character.
			// Ceil of a character is the smallest character greater than it
			ceilIndex := findCeil(str, str[i], i+1, size-1)

			// Swap first and second characters
			str[i], str[ceilIndex] = str[ceilIndex], str[i]

			// Sort the string on right of 'first char'
			indexLow := i + 1
			indexHigh := indexLow + size - i - 1

			if indexLow < size && indexHigh > 0 && indexLow < indexHigh {
				sort.Sort(sortBytes(str[indexLow:indexHigh]))
			}
		}
	}
}

func Pr24() {
	// fmt.Println(findCeil([]byte{'k', 'f', 'a', 'g', 'h', 'c'}, 'c', 0, 5))
	sortedPermutations([]byte{'0', '1', '2'}, 5)
	sortedPermutations([]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}, 1000000)
}

// todos@linux-ew09:~/tmp/project_euler/go> time go run src/main.go
// Max reached:  [50 48 49]
// Readable max:  201
// Max reached:  [50 55 56 51 57 49 53 52 54 48]
// Readable max:  2783915460

// real    0m0.350s
// user    0m0.373s
// sys     0m0.047s
