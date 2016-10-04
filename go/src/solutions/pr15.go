//Starting in the top left corner of a 2×2 grid, and only being able to move to the right and down, there are exactly 6 routes to the bottom right corner.
//
//How many such routes are there through a 20×20 grid?

// Links:
// http://joaoff.com/2008/01/20/a-square-grid-path-problem/
// http://mathforum.org/library/drmath/view/70459.html
// http://betterexplained.com/articles/navigate-a-grid-using-combinations-and-permutations/
// http://math.stackexchange.com/questions/636128/calculating-the-number-of-possible-paths-through-some-squares

// The answer is binomial coefficient: (2 * n)! / n! * n!

package solutions

import (
	"fmt"
	"math/big"
	"math/rand"
	"strconv"
	"strings"
)

func intToStr(num int64, digits int) string {
	//return big.NewInt(num).String()
	str := strconv.FormatInt(num, 2)
	fmt_str := fmt.Sprintf("%%0%ds", digits)
	return fmt.Sprintf(fmt_str, str)
}

func isValid(str string, r_len int, d_len int) bool {
	r_count := strings.Count(str, "0") // r
	d_count := strings.Count(str, "1") // d
	return r_count == r_len && d_count == d_len
}

func bruteSearch(r_len int, d_len int) (result int64) {
	// way too slow, after 428 min no end for 20x20 (Intel(R) Core(TM) i7-4700HQ CPU @ 2.40GHz):

	// todos@linux-ew09:~/tmp/project_euler/go> time go run main.go
	// max is 1099511627776
	// ^Cexit status 2
	//
	// real    428m11.279s
	// user    435m44.713s
	// sys     4m18.684s

	// For 2x2:

	// todos@linux-ew09:~/tmp/project_euler/go> time go run main.go
	// max is 16
	// result: 6
	//
	// real    0m0.180s
	// user    0m0.155s
	// sys     0m0.024s

	// For 10x10:

	// todos@linux-ew09:~/tmp/project_euler/go> time go run main.go
	// max is 1048576
	// result: 184756
	//
	// real    0m1.158s
	// user    0m1.138s
	// sys     0m0.046s

	var max int64
	max = 1 << uint32(r_len+d_len)
	fmt.Printf("max is %d\n", max)
	var i int64
	for i = 0; i < max; i++ {
		str := intToStr(i, r_len+d_len)
		//fmt.Printf("str: %s\n", str)
		if isValid(str, r_len, d_len) {
			result += 1
		}
	}
	return
}

func generatePerm(src string) (dest []string) {
	// Generate permutations of the src string
	// http://www.geeksforgeeks.org/write-a-c-program-to-print-all-permutations-of-a-given-string/
	// https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle
	perm := rand.Perm(len(src))
	for i, v := range perm {
		fmt.Println(i, v)
		//dest[v] = src[i]
	}

	return
}

// http://introcs.cs.princeton.edu/java/23recursion/Permutations.java.html
func sedgePerm(s string) (num int) {
	var N int = len(s)
	var a []rune
	var res [][]rune
	var perms []string
	for i := 0; i < N; i++ {
		a = append(a, rune(s[i]))
	}
	sedgePerm2(a, N, &perms)

	for _, v := range res {
		perms = append(perms, string(v))
	}
	num = len(perms)
	return
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func sedgePerm2(a []rune, n int, perms *[]string) {
	if n == 1 {
		str_a := string(a)
		if !stringInSlice(str_a, *perms) {
			*perms = append(*perms, string(a))
		}
		return
	}
	for i := 0; i < n; i++ {
		sedgeSwap(a, i, n-1)
		sedgePerm2(a, n-1, perms)
		sedgeSwap(a, i, n-1)
	}
}

func sedgeSwap(a []rune, i int, j int) {
	c := a[i]
	a[i] = a[j]
	a[j] = c
}

// https://www.quora.com/What-is-the-fastest-algorithm-to-find-all-possible-permutations-of-a-string
// https://en.wikipedia.org/wiki/Heap%27s_algorithm
func wpSwap(a []rune, i int, j int) {
	c := a[i]
	a[i] = a[j]
	a[j] = c
}

func wpPerm(a []rune, n int, res *[][]rune) {
	if n == 0 {
		//fmt.Println(string(a))
		*res = append(*res, []rune(string(a)))
		return
	}

	for i := 0; i < n; i++ {
		wpPerm(a, n-1, res)
		if n%2 == 0 {
			wpSwap(a, i, n-1)
		} else {
			wpSwap(a, 0, n-1)
		}
	}
}

func wpDeDup(src [][]rune) (dst [][]rune) {
	for _, v := range src {
		if !strIn(v, dst) {
			dst = append(dst, v)
		}
	}
	return
}

func strIn(s []rune, src [][]rune) bool {
	for _, v := range src {
		if sliceEq(v, s) {
			return true
		}
	}
	return false
}

func sliceEq(s []rune, d []rune) bool {
	// Equality for slices is not defined
	for i, v := range s {
		if v != d[i] {
			return false
		}
	}
	return true
}

// Calculates factorial
func fact(n int64, max int64, res *int64) {
	if *res == 0 {
		*res = 1
	}

	if n >= max {
		*res *= n
		return
	}

	*res *= n
	fact(n+1, max, res)
}

func fact_big(max int64, res *big.Int) {
	res.MulRange(1, max)
	return
}

// Factorial wrapper
func factorial(n int64) (res *big.Int) {
	res = new(big.Int)
	fact_big(n, res)
	return
}

func binomial(n int64) (res *big.Int) {
	a := factorial(2 * n)
	b := factorial(n)
	c := new(big.Int)
	c.Mul(b, b)
	//fmt.Println(c)
	res = new(big.Int)
	res.Div(a, c)
	return
}

func Pr15() {
	//fmt.Println("result:", (bruteSearch(20, 20)))
	//fmt.Println("result:", (bruteSearch(4, 4)))
	//fmt.Println("perm:", generatePerm("RRDD"))
	//fmt.Println("perm:", sedgePerm("RRRRRDDDDD"))

	//s1 := "1234"
	//s2 := []rune(s1)
	//fmt.Println(string(s2))
	//wpSwap(s2, 1, 2)
	//fmt.Println(string(s2))

	//s1 := "RRRRDDDD"
	//s2 := []rune(s1)
	//res := [][]rune{}
	//wpPerm(s2, len(s2), &res)
	//fmt.Printf("res: %+v\n", res)
	//fmt.Printf("dedup: %+v\n", len(wpDeDup(res)))

	//var res int64
	//var n int64 = 3
	//fact(1, n, &res)
	//fmt.Printf("fact: %d %d\n", n, res)
	//fmt.Printf("fact: %d %d\n", n, factorial(n))

	fmt.Printf("binomial: %d\n", binomial(20))
}

//todos@linux-ew09:~/tmp/project_euler/go> go run main.go
//binomial: 137846528820
