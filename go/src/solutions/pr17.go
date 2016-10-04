//If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.
//
//If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?
//
//
//NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and forty-two) contains 23 letters and 115 (one hundred and fifteen) contains 20 letters. The use of "and" when writing out numbers is in compliance with British usage.

package solutions

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

const TEN = 10
const ONE_HUNDRED = 100
const ONE_THOUSAND = 1000
const ONE_MILLION = 1000000
const ONE_BILLION = 1000000000           //         1.000.000.000 (9)
const ONE_TRILLION = 1000000000000       //     1.000.000.000.000 (12)
const ONE_QUADRILLION = 1000000000000000 // 1.000.000.000.000.000 (15)
const MAX = 9007199254740992             // 9.007.199.254.740.992 (15)

var LESS_THAN_TWENTY = []string{
	"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten",
	"eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen",
}

var TENTHS_LESS_THAN_HUNDRED = []string{
	"zero", "ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety",
}

func generateWords(number int, arguments ...[]string) string {
	var remainder int
	var word string
	var words []string

	if len(arguments) == 0 {
		words = []string{}
	} else {
		words = arguments[0]
	}

	// We’re done
	if (number == 0) {

		if len(words) == 0 {
			return string("zero")
		} else {
			result := strings.Join(words, " ")

			re := regexp.MustCompile(`,$`)
			result = re.ReplaceAllString(result, "")

			re = regexp.MustCompile(` and$`)
			result = re.ReplaceAllString(result, "")

			return result
		}
	}

	// If negative, prepend "minus"
	if (number < 0) {
		words = append(words, "minus")
		number = int(math.Abs(float64(number)))
	}

	if (number < 20) {
		remainder = 0
		word = LESS_THAN_TWENTY[number]

	} else if (number < ONE_HUNDRED) {
		remainder = number % TEN
		word = TENTHS_LESS_THAN_HUNDRED[int(math.Floor(float64(number / TEN)))]

		// In case of remainder, we need to handle it here to be able to add the “-”
		if (remainder > 0) {
			word += "-" + LESS_THAN_TWENTY[remainder]
			remainder = 0
		}

	} else if (number < ONE_THOUSAND) {
		remainder = number % ONE_HUNDRED
		word = generateWords(int(math.Floor(float64(number / ONE_HUNDRED)))) + " hundred and"

	} else if (number < ONE_MILLION) {
		remainder = number % ONE_THOUSAND
		word = generateWords(int(math.Floor(float64(number / ONE_THOUSAND)))) + " thousand,"

	} else if (number < ONE_BILLION) {
		remainder = number % ONE_MILLION
		word = generateWords(int(math.Floor(float64(number / ONE_MILLION)))) + " million,"

	} else if (number < ONE_TRILLION) {
		remainder = number % ONE_BILLION
		word = generateWords(int(math.Floor(float64(number / ONE_BILLION)))) + " billion,"

	} else if (number < ONE_QUADRILLION) {
		remainder = number % ONE_TRILLION
		word = generateWords(int(math.Floor(float64(number / ONE_TRILLION)))) + " trillion,"

	} else if (number <= MAX) {
		remainder = number % ONE_QUADRILLION
		word = generateWords(int(math.Floor(float64(number / ONE_QUADRILLION)))) + " quadrillion,"
	}

	words = append(words, word)
	return generateWords(remainder, words)
}

func Pr17() {
	var numarr []string
	result := 0
	for num := 1; num <= 1000; num++ {
		str := generateWords(num, []string{})
		re := regexp.MustCompile(`(\s|-)`)
		parts := re.Split(str, -1)

		for i := 0; i < len(parts); i++ {
			numarr = append(numarr, parts[i])
			result += len(parts[i])
		}
	}
	fmt.Printf("%d", result)

	// /home/todos/go/go1.7.1/bin/go run /home/todos/tmp/project_euler/go/src/main.go
	// 21124

	//fmt.Printf("%s\n", generateWords(0, []string{}))
	//fmt.Printf("%s\n", generateWords(1, []string{}))
	//fmt.Printf("%s\n", generateWords(78, []string{}))
	//fmt.Printf("%s\n", generateWords(99, []string{}))
	//fmt.Printf("%s\n", generateWords(299, []string{}))
	//fmt.Printf("%s\n", generateWords(300, []string{}))
	//fmt.Printf("%s\n", generateWords(999, []string{}))
	//fmt.Printf("%s\n", generateWords(1000, []string{}))
	//fmt.Printf("%s\n", generateWords(1000999, []string{}))
	//fmt.Printf("%s\n", generateWords(111222, []string{}))
	//fmt.Printf("%s\n", generateWords(-111222, []string{}))
	//fmt.Printf("%s\n", generateWords(-222, []string{}))
}
