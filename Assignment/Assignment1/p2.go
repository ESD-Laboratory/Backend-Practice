package Assignment1

import (
	"strings"
	"unicode"
)

/*
Clean up sentences. [WEIGHT = 1]

Given a string, you need to clean up and remove unnecessary characters.
The only characters that are allowed are:
    - letters (a-z or A-Z)
    - numbers (0-9)
    - space character (" ")

Replace every other characters with a single space. And then, remove all spaces
in the beginning and at the end of the string (leading/trailing spaces).

Finally, return the resulting string.

Examples
------------------------------------
Input: "I am learning Go???"
Output: "I am learning Go"
Explanation: Remove the question marks at the end
------------------------------------
Input: "Everything is good"
Output: "Everything is good"
Explanation: Nothing changes so return string as is
------------------------------------
Input: "    Independence day: August 17th, 1945 "
Output: "Independence day August 17th 1945"
Explanation: Replace the : and , and then remove all trailing/leading spaces
------------------------------------
*/

func CleanUp(s string) string {
	// TODO: Your code here
	var cleanedStr strings.Builder
	prevCharWasSpace := false
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) || r == ' ' {
			if r == ' ' {
				if !prevCharWasSpace {
					cleanedStr.WriteRune(r)
				}
				prevCharWasSpace = true
			} else {
				cleanedStr.WriteRune(r)
				prevCharWasSpace = false
			}
		} else {
			if !prevCharWasSpace {
				cleanedStr.WriteRune(' ')
			}
			prevCharWasSpace = true
		}
	}

	return strings.TrimSpace(cleanedStr.String())
}
