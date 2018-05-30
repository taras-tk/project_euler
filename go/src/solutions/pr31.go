// In England the currency is made up of pound, £, and pence, p, and there are eight coins in general circulation:
//     1p, 2p, 5p, 10p, 20p, 50p, £1 (100p) and £2 (200p).
//
// It is possible to make £2 in the following way:
//     1×£1 + 1×50p + 2×20p + 1×5p + 1×2p + 3×1p
//
// How many different ways can £2 be made using any number of coins?

package solutions

import (
	"fmt"
)

func caseCalc(p1 int, p2 int, p5 int, p10 int, p20 int, p50 int, p100 int) (result int) {
	result = p1*1 + p2*2 + p5*5 + p10*10 + p20*20 + p50*50 + p100*100
	return
}

func calc() (result [][]int) {
	for p1 := 0; p1 <= 200; p1++ {
		for p2 := 0; p2 <= 100; p2++ {
			for p5 := 0; p5 <= 40; p5++ {
				for p10 := 0; p10 <= 20; p10++ {
					for p20 := 0; p20 <= 10; p20++ {
						for p50 := 0; p50 <= 4; p50++ {
							for p100 := 0; p100 <= 2; p100++ {
								calc := caseCalc(p1, p2, p5, p10, p20, p50, p100)
								if calc == 200 {
									slic1 := []int{p1, p2, p5, p10, p20, p50, p100}
									result = append(result, slic1)
								}
							}
						}
					}
				}
			}
		}

	}
	return
}

func Pr31() {
	calc1 := calc()
	// fmt.Println(calc1)
	fmt.Println(len(calc1))
	fmt.Printf("Result: %d\n", len(calc1)+1)
}

// ➜  src git:(master) ✗ time go run main.go
// 73681
// Result: 73682
// go run main.go  5,76s user 0,08s system 104% cpu 5,600 total
