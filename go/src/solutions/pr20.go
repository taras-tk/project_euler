// n! means n × (n − 1) × ... × 3 × 2 × 1
//
// For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
// and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.
//
// Find the sum of the digits in the number 100!

// Wrong answer:
// todos@linux-ew09:~/tmp/project_euler/go/src> python -c "from functools import reduce; dig = '7629259731982352384'; sum = reduce(lambda x, y: x+y, [ int(x) for x in dig ], 0); print 'sum={}'.format(sum)"
// sum=95

// Based on online factorial calculator:
// todos@linux-ew09:~/tmp/project_euler/go/src> python -c "from functools import reduce; dig = '93326215443944152681699238856266700490715968264381621468592963895217599993229915608941463976156518286253697920827223758251185210916864000000000000000000000000'; sum = reduce(lambda x, y: x+y, [ int(x) for x in dig ], 0); print 'sum={}'.format(sum)"
// sum=648

// todos@linux-ew09:~/tmp/project_euler/go/src> go run main.go
// 10 -> 1
// 101 -> 2
// 1101 -> 3
// 10 -> fact 3628800 -> sum 27
// 100 -> fact 93326215443944152681699238856266700490715968264381621468592963895217599993229915608941463976156518286253697920827223758251185210916864000000000000000000000000 -> sum 648

package solutions

import (
	"fmt"
	"math/big"
	"strconv"
)

// Calculates factorial for int64
func factInt64(n int64, max int64, res *int64) {
	if *res == 0 {
		*res = 1
	}

	if n >= max {
		*res *= n
		return
	}

	*res *= n
	factInt64(n+1, max, res)
}

func numToSumInt64(n int64) (sum int) {
	s := strconv.Itoa(int(n))
	for _, s1 := range s {
		s2 := string(s1)
		x, _ := strconv.Atoi(s2)
		sum += x
	}
	return
}

func numToSumBigInt(n *big.Int) (sum int) {
	s := n.String()
	for _, s1 := range s {
		s2 := string(s1)
		x, _ := strconv.Atoi(s2)
		sum += x
	}
	return
}

func factBigInt(max int64, res *big.Int) {
	res.MulRange(1, max)
	return
}

// Factorial wrapper
func factWrapper(n int64) (res *big.Int) {
	res = new(big.Int)
	factBigInt(n, res)
	return
}

func Pr20() {
	//var n = 10
	fmt.Printf("%d -> %d\n", 10, numToSumInt64(10))
	fmt.Printf("%d -> %d\n", 101, numToSumInt64(101))
	fmt.Printf("%d -> %d\n", 1101, numToSumInt64(1101))
	var f1 int64
	var n1 int64 = 10
	factInt64(1, n1, &f1)
	fmt.Printf("%d -> fact %d -> sum %d\n", n1, f1, numToSumInt64(f1))

	var n2 int64 = 100
	f2 := factWrapper(n2)
	fmt.Printf("%d -> fact %d -> sum %d\n", n2, f2, numToSumBigInt(f2))
}
