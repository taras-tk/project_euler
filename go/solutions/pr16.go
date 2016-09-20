// 2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.
// 
// What is the sum of the digits of the number 2^1000?

package solutions

import (
	"fmt"
	"math/big"
	"strconv"
)

func intToString(num *big.Int) string {
	return num.String()
}

func powBigInt(base int64, exp int64) *big.Int {
	big_base := new(big.Int)
	big_base.SetInt64(base)

	big_exp := new(big.Int)
	big_exp.SetInt64(exp)

	res := new(big.Int)
	res.Exp(big_base, big_exp, nil)
	return res
}

func sumDigits(str string) (cnt int, res int) {
	for _, v := range(str) {
		d, err := strconv.Atoi(string(v))
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		} else {
			res += d
			cnt += 1
		}
	}
	return
}

func Pr16() {
	fmt.Printf("Pow:%d\n", powBigInt(2, 1000))

	cnt, res := sumDigits("123")
	fmt.Printf("sum:%d, %d\n", cnt, res)

	cnt, res = sumDigits(intToString(powBigInt(2, 15)))
	fmt.Printf("sum2:%d, %d\n", cnt, res)

	cnt, res = sumDigits(intToString(powBigInt(2, 1000)))
	fmt.Printf("sum3:%d, %d\n", cnt, res)
}

//todos@linux-ew09:~/tmp/project_euler/go> go run main.go 
//Pow:10715086071862673209484250490600018105614048117055336074437503883703510511249361224931983788156958581275946729175531468251871452856923140435984577574698574803934567774824230985421074605062371141877954182153046474983581941267398767559165543946077062914571196477686542167660429831652624386837205668069376
//sum:3, 6
//sum2:5, 26
//sum3:302, 1366
