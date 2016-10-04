//The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
//
//Find the sum of all the primes below two million.

// Articles used in this algorithm:
// https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes

package solutions

import (
	"math"
	"fmt"
)

func sieve10(n uint64) (res []uint64) {
	var arr []bool = make([]bool, n)

	for i,_ := range arr {
		arr[i] = true
	}

	fmt.Println("allocated bool array")

	var sqrt = uint64(math.Sqrt(float64(n)))

	var i uint64
	for i = 2; i <= sqrt; i = i+1 {
		if arr[i] == true {
			var q uint64 = 0
			for j := i*i; j < n; j = i*i + q*i {
				q = q+1
				arr[j] = false
			}
		}
	}

	fmt.Println("built bool array")

	var len_a uint64 = uint64(len(arr))
	for i = 0; i < len_a; i++ {
		if arr[i] == true && i != 0 && i != 1 {
			res = append(res, i)
			if i >= n {
				break
			}
		}
	}

	return res
}

func Pr10() {
	var num uint64 = 2000000
	svs := sieve10(num)
	var sum uint64
	for _, v := range(svs) {
		sum += v
	}

	//fmt.Printf("%v: %v\n", num, svs)
	fmt.Printf("%v\n", sum)
}

//time go run main.go
//allocated bool array
//built bool array
//142913828922
//
//real    0m0.138s
//user    0m0.117s
//sys     0m0.021s
