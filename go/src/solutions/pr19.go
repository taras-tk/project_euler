// You are given the following information, but you may prefer to do some research for yourself.

// 1 Jan 1900 was a Monday.
// Thirty days has September,
// April, June and November.
// All the rest have thirty-one,
// Saving February alone,
// Which has twenty-eight, rain or shine.
// And on leap years, twenty-nine.
// A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.
// How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?

// Google play: https://play.golang.org/p/rfAhyvpWXE

package solutions

import (
	"fmt"
	"time"
)

func Pr19() {
	sundays := 0
	for year := 1901; year <= 2000; year++ {
		month := time.January
		for m := 1; m <= 12; m++ {
			if time.Date(year, month, 1, 0, 0, 0, 0, time.UTC).Weekday() == time.Sunday {
				sundays++
			}

			month++
		}
	}
	fmt.Printf("%d sundays\n", sundays)
	// 171 sundays
}

func Pr19_2() {
	t0 := time.Date(1901, time.January, 1, 0, 0, 0, 0, time.UTC)
	t1 := time.Date(2000, time.December, 31, 0, 0, 0, 0, time.UTC)

	diff := t1.Sub(t0)

	// convert diff to days
	days := int(diff.Hours() / 24)

	result := (int(t0.Weekday()-1) + days) / 7

	fmt.Printf("%d days, %d => %d \n", days, t0.Weekday(), result)
}
