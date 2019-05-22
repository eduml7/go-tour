/*
Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite
against the provided function and prints success or failure. You might find strings.Fields helpful. 
*/

package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	monk := strings.Fields(s)
	m := make(map[string]int)
	for _, c := range monk {
		m[c] = m[c] + 1
	}
	return m //map[string]int{"x": 1}
}

func main() {
	wc.Test(WordCount)
}

/*
PASS
 f("I am learning Go!") = 
  map[string]int{"Go!":1, "I":1, "am":1, "learning":1}
PASS
 f("The quick brown fox jumped over the lazy dog.") = 
  map[string]int{"The":1, "brown":1, "dog.":1, "fox":1, "jumped":1, "lazy":1, "over":1, "quick":1, "the":1}
PASS
 f("I ate a donut. Then I ate another donut.") = 
  map[string]int{"I":2, "Then":1, "a":1, "another":1, "ate":2, "donut.":2}
PASS
 f("A man a plan a canal panama.") = 
  map[string]int{"A":1, "a":2, "canal":1, "man":1, "panama.":1, "plan":1}
*/
