// Using names.txt (right click and 'Save Link/Target As...'), a 46K text file containing over five-thousand first names, begin by sorting it into alphabetical order. Then working out the alphabetical value for each name, multiply this value by its alphabetical position in the list to obtain a name score.

// For example, when the list is sorted into alphabetical order, COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list. So, COLIN would obtain a score of 938 × 53 = 49714.

// What is the total of all the name scores in the file?

// todos@linux-ew09:~/tmp/project_euler/go/src> time go run main.go
// Pr22:  COLIN 53
// readNames:  [MARY PATRICIA LINDA BARBARA ELIZABETH JENNIFER MARIA SUSAN MARGARET DOROTHY]
// namesAndAlphabeticalNumbers: [%!r(*solutions.NameNumber=&{AARON 1 49}) %!r(*solutions.NameNumber=&{ABBEY 2 35}) %!r(*solutions.NameNumber=&{ABBIE 3 19}) %!r(*solutions.NameNumber=&{ABBY 4 30}) %!r(*solutions.NameNumber=&{ABDUL 5 40}) %!r(*solutions.NameNumber=&{ABE 6 8}) %!r(*solutions.NameNumber=&{ABEL 7 20}) %!r(*solutions.NameNumber=&{ABIGAIL 8 41}) %!r(*solutions.NameNumber=&{ABRAHAM 9 44}) %!r(*solutions.NameNumber=&{ABRAM 10 35})]
// Result:  871198282

// real    0m0.260s
// user    0m0.282s
// sys     0m0.052s

package solutions

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func alphabeticalValue(name string) (result int) {
	for _, r := range name {
		index := int(r) + 1 - int('A')
		// fmt.Println(index)
		result += index
	}
	return
}

func readNames(path string) (result []string) {
	file, err := os.Open(path)
	if err != nil {
		panic("Cannot read file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var str string
	for scanner.Scan() {
		str = scanner.Text()
	}

	arr := strings.Split(str, ",")
	for i := 0; i < len(arr); i++ {
		s := arr[i]
		s = strings.TrimSuffix(s, "\"")
		s = strings.TrimPrefix(s, "\"")
		result = append(result, s)
	}
	return
}

type NameNumber struct {
	name   string
	index  int
	number int
}

func namesAndAlphabeticalNumbers(slice []string) (result []*NameNumber) {
	sort.Strings(slice)
	// fmt.Println(slice)
	// fmt.Println(slice[937])
	for i, name := range slice {
		st := new(NameNumber)
		st.name = name
		st.index = i + 1
		st.number = alphabeticalValue(name)
		// fmt.Printf("%v\n", st)
		result = append(result, st)
	}
	// fmt.Printf("%r\n", result[937:940])
	return
}

func scoreFromNameNumbers(narr []*NameNumber) (result int64) {
	for _, obj := range narr {
		result += int64(obj.index * obj.number)
	}
	return
}

func Pr22() {
	name := "COLIN"
	fmt.Println("Pr22: ", name, alphabeticalValue(name))
	names := readNames("./solutions/pr22_names.txt")
	fmt.Println("readNames: ", names[0:10])
	// fmt.Printf("namesAndAlphabeticalNumbers: %r", namesAndAlphabeticalNumbers(names[0:10]))
	// fmt.Printf("namesAndAlphabeticalNumbers: %r", namesAndAlphabeticalNumbers(names))
	// namesAndAlphabeticalNumbers(names)
	alphanums := namesAndAlphabeticalNumbers(names)
	fmt.Printf("namesAndAlphabeticalNumbers: %r\n", alphanums[0:10])
	score := scoreFromNameNumbers(alphanums)
	fmt.Println("Result: ", score)
}
