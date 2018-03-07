package solutions

// Starting with the number 1 and moving to the right in a clockwise direction a 5 by 5 spiral is formed as follows:

// 21 22 23 24 25
// 20  7  8  9 10
// 19  6  1  2 11
// 18  5  4  3 12
// 17 16 15 14 13

// It can be verified that the sum of the numbers on the diagonals is 101.
// What is the sum of the numbers on the diagonals in a 1001 by 1001 spiral formed in the same way?

import (
	"fmt"
)

func genMatrix(dimension int) (result [][]int) {
    // 21 22 23 24 25
    // 20  7  8  9 10
    // 19  6  1  2 11
    // 18  5  4  3 12
	// 17 16 15 14 13

	result = make([][]int, dimension)
    for i := range result {
        result[i] = make([]int, dimension)
	}

	maxI := dimension - 1
	minI := 0
	currentI := maxI
	maxJ := dimension - 1
	minJ := 0
	currentJ := minJ
	for currentNum := dimension*dimension; currentNum > 0; {
		// right to left
		for currentI = maxI; currentI >= minI; currentI-- {
			result[minJ][currentI] = currentNum
			currentNum--
		}
		// top to bottom
		for currentJ = minJ+1; currentJ <= maxJ; currentJ++ {
			result[currentJ][minI] = currentNum
			currentNum--
		}
		// left to right
		for currentI = minI+1; currentI <= maxI; currentI++ {
			result[maxJ][currentI] = currentNum
			currentNum--
		}
		// bottom to top
		for currentJ = maxJ-1; currentJ >= (minJ+1); currentJ-- {
			result[currentJ][maxJ] = currentNum
			currentNum--
		}

		maxI--
		minI++
		maxJ--
		minJ++
	}
	return
}

func sumDiags(matrix [][]int) (result int) {
	maxLen := len(matrix[0])
	result = 0
	j := 0
	for i := 0; i < maxLen; i++ {
		result += matrix[i][j]
		j++
	}
	j = 0
	for i := maxLen-1; i >= 0; i-- {
		if i != j {
			// do not count number twice
			result += matrix[i][j]
		}
		j++
	}
	return
}

func Pr28() {
	matrix := genMatrix(1001)
	// fmt.Println(matrix)
	fmt.Println(sumDiags(matrix))
}

// Tarass-MBP-3:src todos$ time go run main.go
// 669171001

// real    0m0.349s
// user    0m0.293s
// sys     0m0.100s