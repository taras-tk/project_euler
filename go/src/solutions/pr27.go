package solutions


import (
	"fmt"
	"math"
)

// for small prime
func isPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value) / 2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func formula(a int, b int, n int) (result int) {
	result = n * n + a * n + b
	return
}

func numPrimes(a int, b int) (result int) {
	for n := 0; n <= int(math.Floor(math.Pow(2, 32))); n++ {
		x := formula(a, b, n)
		if isPrime(x) {
			result = n + 1
		} else {
			break
		}
	}
	return
}

func generatePrimes(a int, b int) (result [][]int) {
	for x_a := -a+1; x_a < a; x_a++ {
		for x_b := -b; x_b <= b; x_b++ {
			d := []int{x_a, x_b, numPrimes(x_a, x_b)}
			result = append(result, d)
		}
	}
	return
}

func maxPrimes(data [][]int) (result []int) {
	max := 0
	for i := 0; i < len(data); i++ {
		if max < data[i][2] {
			max = data[i][2]
			result = data[i]
		}
	}
	return
}

func Pr27() {
	fmt.Println(maxPrimes(generatePrimes(10, 10)))
	fmt.Println(maxPrimes(generatePrimes(100, 100)))
	max := maxPrimes(generatePrimes(1000, 1000))
	fmt.Println(max[0] * max[1])
}

//Tarass-MBP-3:src todos$ time go run main.go
//[-3 7 6]
//[-15 97 48]
//-59231
//
//real    0m2.859s
//user    0m3.264s
//sys     0m0.315s
