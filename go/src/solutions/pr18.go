// By starting at the top of the triangle below and moving to adjacent numbers on the row below, the maximum total from top to bottom is 23.
//
//    3
//   7 4
//  2 4 6
// 8 5 9 3
//
// That is, 3 + 7 + 4 + 9 = 23.
//
// Find the maximum total from top to bottom of the triangle below:
//
// 75
// 95 64
// 17 47 82
// 18 35 87 10
// 20 04 82 47 65
// 19 01 23 75 03 34
// 88 02 77 73 07 63 67
// 99 65 04 28 06 16 70 92
// 41 41 26 56 83 40 80 70 33
// 41 48 72 33 47 32 37 16 94 29
// 53 71 44 65 25 43 91 52 97 51 14
// 70 11 33 28 77 73 17 78 39 68 17 57
// 91 71 52 38 17 14 91 43 58 50 27 29 48
// 63 66 04 68 89 53 67 30 73 16 69 87 40 31
// 04 62 98 27 23 09 70 98 73 93 38 53 60 04 23
//
// NOTE: As there are only 16384 routes, it is possible to solve this problem by trying every route. However, Problem 67, is the same challenge with a triangle containing one-hundred rows; it cannot be solved by brute force, and requires a clever method! ;o)

package solutions

import (
	"fmt"
)

type Element struct {
	weight int;
	index []int;
}

//    3
//   7 4
//  2 4 6
// 8 5 9 3
func walk(arr [][]int) (result [][]Element) {
	for row := 0; row < len(arr[len(arr)-1]); row++ {
		res := []Element{}
		for col := 0; col < len(arr[row]); col++ {
			var sum []int
			if row == 0 {
				// single parent
				sum = []int{arr[row][col]}
			} else if col == 0 {
				// single parent
				for _, weight := range(Element(result[row - 1][col]).index) {
					sum = append(sum, weight + arr[row][col])
				}
			} else if col == (len(arr[row]) - 1) {
				// single parent
				for _, weight := range(Element(result[row - 1][col - 1]).index) {
					sum = append(sum, weight + arr[row][col])
				}
			} else {
				// two parents
				for _, weight := range(Element(result[row - 1][col - 1]).index) {
					sum = append(sum, weight + arr[row][col])
				}

				for _, weight := range(Element(result[row - 1][col]).index) {
					sum = append(sum, weight + arr[row][col])
				}
			}
			elt := Element{weight: arr[row][col], index: sum}
			res = append(res, elt)
		}
		result = append(result, res)
	}
	return
}

func find_max(arr [][]Element) (max int) {
	i := len(arr) - 1
	for j := 0; j < len(arr[len(arr)-1]); j++ {
		for _, sum := range(Element(arr[i][j]).index) {
			if sum > max {
				max = sum
			}
		}
	}
	return
}

func Pr18() {
	numarr := [][]int{[]int{3}, []int{7, 4}, []int{2, 4, 6}, []int{8, 5, 9, 3}}
	fmt.Println(walk(numarr))
	fmt.Println(find_max(walk(numarr)))

	// 75
	// 95 64
	// 17 47 82
	// 18 35 87 10
	// 20 04 82 47 65
	// 19 01 23 75 03 34
	// 88 02 77 73 07 63 67
	// 99 65 04 28 06 16 70 92
	// 41 41 26 56 83 40 80 70 33
	// 41 48 72 33 47 32 37 16 94 29
	// 53 71 44 65 25 43 91 52 97 51 14
	// 70 11 33 28 77 73 17 78 39 68 17 57
	// 91 71 52 38 17 14 91 43 58 50 27 29 48
	// 63 66 04 68 89 53 67 30 73 16 69 87 40 31
	// 04 62 98 27 23 09 70 98 73 93 38 53 60 04 23
	numarr = [][]int{
		[]int{75},
		[]int{95, 64},
		[]int{17, 47, 82},
		[]int{18, 35, 87, 10},
		[]int{20,  4, 82, 47, 65},
		[]int{19,  1, 23, 75,  3, 34},
		[]int{88,  2, 77, 73,  7, 63, 67,},
		[]int{99, 65,  4, 28,  6, 16, 70, 92,},
		[]int{41, 41, 26, 56, 83, 40, 80, 70, 33,},
		[]int{41, 48, 72, 33, 47, 32, 37, 16, 94, 29,},
		[]int{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14,},
		[]int{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57,},
		[]int{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48,},
		[]int{63, 66, 04, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31},
		[]int{04, 62, 98, 27, 23,  9, 70, 98, 73, 93, 38, 53, 60,  4, 23},
	}
	fmt.Println(walk(numarr))
	fmt.Println(find_max(walk(numarr)))
	// Answer: 1074
}
