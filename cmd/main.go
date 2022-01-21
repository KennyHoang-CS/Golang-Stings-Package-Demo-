package main

import (
	"fmt"
	"strings"
)

func main() {

	// compare two strings
	fmt.Println("strings.Compare('Kenny', 'Hoang'): ", strings.Compare("Kenny", "Hoang"))
	fmt.Println("strings.Compare('Kenny', 'Kenny'): ", strings.Compare("Kenny", "Kenny"))

	// checks if substr is within a string
	fmt.Println("strings.Contains('Kenny', 'Ken'): ", strings.Contains("Kenny", "Ken"))
	fmt.Println("strings.Contains('Kenny', 'asdKn'):", strings.Contains("Kenny", "asdKn"))

	// check if any unicode chars are within a string
	fmt.Println("strings.ContainsAny('Kenny', 't'): ", strings.ContainsAny("Kenny", "t"))

	// Count the number of non-overlapping instances of substr
	fmt.Println("strings.Count('Kenny', 'n'): ", strings.Count("Kenny", "n"))

	// ToLower returns s with all Unicode letters mapped to their lower case.
	fmt.Println("strings.ToLower('KENNY'): ", strings.ToLower("KENNY"))

	// ToUpper returns s with all Unicode letters mapped to their upper case.
	fmt.Println("strings.ToUpper('kenny'): ", strings.ToUpper("kenny"))

	// TrimSpace returns a slice of the string s, with all leading and trailing white space removed, as defined by Unicode.
	fmt.Println("strings.TrimSpace('\t\n Hello, Kenny \n\t\r\n'): ", strings.TrimSpace("\t\n Hello, Kenny \n\t\r\n"))

	// Join concatenates the elements of its first argument to create a single string. The separator string sep is placed between elements in the resulting string.
	fmt.Println("strings.Join(['I', 'Love', 'Cookies'], ', '): ", strings.Join([]string{"I", "love", "cookies"}, ", "))
}
