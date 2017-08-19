package solutions

import (
	"fmt"
)

// Function to find length of period in 1/n
func getPeriod(n int) (count int) {
	// Find the (n+1)th remainder after decimal point
	// in value of 1/n
	rem := 1 // Initialize remainder
	for i := 1; i <= n+1; i++ {
		rem = (10 * rem) % n
	}

	// Store (n+1)th remainder
	var d = rem

	// Count the number of remainders before next
	// occurrence of (n+1)'th remainder 'd'
	count = 0
	for {
		rem = (10 * rem) % n
		count++
		if rem == d {
			break
		}
	}
	return
}

func getMax(d int) (dmax int, countmax int) {
	dmax = 0
	countmax = 0
	for i := 1; i <= d; i++ {
		count := getPeriod(i)
		if count > countmax {
			countmax = count
			dmax = i
		}
	}
	return
}

func Pr26() {
	fmt.Println(getPeriod(3))
	fmt.Println(getPeriod(6))
	fmt.Println(getPeriod(7))
	fmt.Println(getPeriod(210))
	fmt.Println("---")
	fmt.Println(getMax(1000))
}

// Tarass-MBP:go todos$ go run pr26.go
// 1
// 1
// 6
// 6
// ---
// 983 982

// http://www.geeksforgeeks.org/find-length-period-decimal-value-1n/

// // Function to find length of period in 1/n
// int getPeriod(int n)
// {
//    // Find the (n+1)th remainder after decimal point
//    // in value of 1/n
//    int rem = 1; // Initialize remainder
//    for (int i = 1; i <= n+1; i++)
//          rem = (10*rem) % n;
//
//    // Store (n+1)th remainder
//    int d = rem;
//
//    // Count the number of remainders before next
//    // occurrence of (n+1)'th remainder 'd'
//    int count = 0;
//    do {
//       rem = (10*rem) % n;
//       count++;
//    } while(rem != d);
//
//    return count;
// }
