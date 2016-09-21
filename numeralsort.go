/*
  Numeral-aware text sort

    "file1" < "file2" < "file10"

  I wrote this same idea years ago in 68000 assembler for the custom open-file dialog on an Amiga computer.
  It's still useful now, and will be as long as people name files like humans.

  Copyright 2016 Nicolas S. Dade
*/

package numeralsort

import (
	"sort"
	"strings"
)

// Less returns true if x < y in a numeral-aware comparison.
// It is suitable for use with Go's standard sort.Interface.
func Less(a, b string) bool {
	// the idea is to scan along a and b rune-by-rune (might as well the UTF-8 ready),
	// until a numeric [0-9] rune is reached in both strings. Then the numbers are
	// decoded and compared. If equal the text comparison continues... .
	// numbers are assumed to be in unsigned decimal (which is common. hex typically has 0000 prefixes)
	digits := "0123456789"

	for {
		i := strings.IndexAny(a, digits)
		j := strings.IndexAny(b, digits)
		if i < 0 || j < 0 {
			// no numeral to compare. finish up with a straight string comparison
			return a < b
		}
		if i != j || a[:i] != b[:j] {
			// text differs. finish by comparing the text
			return a[:i] < b[:j]
		}
		// a and b match up to i (which equals j at this point), and then that is the start of a numeral
		a = a[i:]
		b = b[j:]

		// decode the numeral. since small numbers are common I check if it might fit in uint64 before resorting to
		// bignum
		// first find the end of each numeral. There is no strings.IndexNotAny() (at least not in go1.7), so I have to write my own
		// function to find the first non-numeral rune
		x, a := extractNumeral(a)
		y, b := extractNumeral(b)

		// numeral section matched; return to matching text
	}
}

// extractNumeral extracts the numeral prefix of a.
// It returns the numeral decoded as a uint64, or if that doesn't fit, a big.Int
func extractNumeral(a string) (string, string) {
}

// Strings is a slice of strings sortable in numeral-aware order
// It implements sort.Interface
type Strings []string

func (s Strings) Len() int           { return len(s) }
func (s Strings) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Strings) Less(i, j int) bool { return Less(s[i], s[j]) }

// Sort is a utility function to sort a slice of strings using numeral-aware sort order
func Sort(strs []string) {
	sort.Sort(Strings(strs))
}
