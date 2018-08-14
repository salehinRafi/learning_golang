package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	// Find a pattern matches a string directly (applied to MatchString task only)
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println("MatchString: ", match)

	// For other regexp task, we need to `compile` an optimized Regexp struct
	r, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println("Plain Compiler: ", r)

	// Finds the pattern matches a string.
	fmt.Println("MatchString: ", r.MatchString("peach"))

	// Finds the match for the regexp. (the first it found)
	fmt.Println("FindString: ", r.FindString("peace pinch peach"))

	// Finds the start and end indexes of matches (the first it found)
	fmt.Println("FindStringIndex: ", r.FindStringIndex("punch pinch peach"))

	// Finds the whole-pattern matches and submatch withon those matches (return p([a-z]+)ch and ([a-z]+).)
	fmt.Println("FindStringSubmatch: ", r.FindStringSubmatch("peach"))

	// Finds the start and end indexes of matches and submatches (p([a-z]+)ch and ([a-z]+).)
	fmt.Println("FindStringSubmatchIndex: ", r.FindStringSubmatchIndex("peach"))

	// Finds all matches in the input (not just the first)
	fmt.Println("FindAllString: ", r.FindAllString("peach punch pinch", -1))

	// Finds the whole start and indexes of all matches
	fmt.Println("FindAllStringSubmatchIndex: ", r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	// Finds the limit number of matches in the input
	fmt.Println("FindAllString: ", r.FindAllString("peach punch pinch", 2))

	fmt.Println("Match and []byte arg: ", r.Match([]byte("peach")))

	// Creating constants with regular expressions. A plain Compile won’t work for constants because it has 2 return values.
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("Constant Compile: ", r)

	// replace subsets of strings with other values.
	fmt.Println("ReplaceAllString: ", r.ReplaceAllString("a peach", "<fruit>"))

	// Func variant allows to transform matched text with a given function.
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println("Transform Matched: ", string(out))
}
