//If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
//
//Find the sum of all the multiples of 3 or 5 below 1000.

//Answer:
//todos@linux-ew09:~/tmp/project_euler/go> echo 10|go run main.go 
//23
//todos@linux-ew09:~/tmp/project_euler/go> echo 1000|go run main.go 
//233168

package solutions

import (
	"fmt"
)

func Pr1() {
	var result_list []int
	var num int
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		fmt.Println("Error: %v", err)
		return
	}
	for v := 1; v < num; v++ {
		if v % 3 == 0 || v % 5 == 0 {
			result_list = append(result_list, v)
		}
	}
	//fmt.Println(result_list)

	result_num := 0
	for _, v := range result_list {
		result_num += v
	}
	fmt.Println(result_num)
}
