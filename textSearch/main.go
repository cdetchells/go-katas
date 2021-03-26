package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	textToSearchPtr := flag.String("textToSearch", "", "The text to search")
	subtextPtr := flag.String("subtext", "", "The text that you're looking for")
	flag.Parse()
	output := textSearch(*textToSearchPtr, *subtextPtr)
	fmt.Println(output)
}

func textSearch(textToSearch, subtext string) string {
	// If there's no text to search or no subtext then we can just return an empty output
	if len(textToSearch) == 0 || len(subtext) == 0 {
		return ""
	}
	// Used to store the positions of the matched strings
	var matches []string
	// Used to count up the number of matched characters
	j := 0
	for i := 0; i < len(textToSearch); i++ {
		// We have a case sensitive character match
		if textToSearch[i] == subtext[j] {
			j++
		} else {
			// We have a case insensitive character match
			if unicode.ToLower(rune(textToSearch[i])) == unicode.ToLower(rune(subtext[j])) {
				j++
			} else {
				// No character match so reset back to 0
				j = 0
			}
		}
		// We have a match
		if j == len(subtext) {
			// Store the start position
			matches = append(matches, strconv.Itoa((i+1)-(j-1)))
			// Reset back to 0
			j = 0
		}
	}
	// Return the array as a comma separated string
	return strings.Join(matches, ",")
}
