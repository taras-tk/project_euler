//By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
//
//What is the 10 001st prime number?

// Articles used in this algorithm:
// https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes

// TBD: current implementation bfsieve uses 4G RAM (2**32),
// should be shortened using segmented sieve
// TBD: execution time is 1m18.413s should shorten as well

package main

import (
	"fmt"
	"math"
)


func bfsieve(cnt int) (res []uint64) {
	var n uint64 = math.MaxUint32
	var arr []bool = make([]bool, n)

	for i,_ := range arr {
		arr[i] = true
	}

	fmt.Println("Allocated bool array")

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

	fmt.Println("Built bool array")
	
	var cnt_rnt int
	var len_a uint64 = uint64(len(arr))
	for i = 0; i < len_a; i++ {
		if arr[i] == true && i != 0 && i != 1 {
			res = append(res, i)
			cnt_rnt += 1
			if cnt_rnt >= cnt {
				break
			}
		}
	}

	return res
}

func main() {
	res := bfsieve(10001)
	//fmt.Println(bfsieve(10001))
	len_r := len(res)
	fmt.Printf("%v %v", res[len_r-1], len_r)
}
