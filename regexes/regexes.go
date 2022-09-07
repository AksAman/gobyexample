package regexes

import (
	"fmt"
	"regexp"
	"strings"
)

func Run() {

	pattern1 := "p([a-z]+)ch"
	pattern2 := "p([a-z]+)ch ([0-9]+)"

	// check for matches
	if match, err := regexp.MatchString(pattern1, "peach"); err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("match: %v\n", match)
	}

	r1, _ := regexp.Compile(pattern1)
	r2, _ := regexp.Compile(pattern2)
	fmt.Println("match", r1.MatchString("peach"))

	// find
	testString1 := "a small peach punch"
	testString2 := "a small peach 1234AB punch"
	result := r1.FindString(testString1)
	fmt.Printf("result: %v\n", result)

	indices := r1.FindStringIndex(testString1)
	fmt.Printf("indices: %v\n", indices)

	allIndices := r1.FindAllStringIndex(testString1, 2)
	fmt.Printf("allIndices: %v\n", allIndices)

	subMatch := r2.FindStringSubmatch(testString2)
	fmt.Printf("subMatch: %v\n", subMatch)

	byteMatch := r1.Match([]byte(testString1))
	fmt.Printf("byteMatch: %v\n", byteMatch)

	replaced := r2.ReplaceAllStringFunc(testString2, strings.ToUpper)
	fmt.Printf("replaced: %v\n", replaced)
}
