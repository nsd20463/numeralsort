/*
  Numeral-aware text sort

    "file1" < "file01" < "file2" < "file10"

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
		var x, y string
		x, a = extractNumeral(a)
		y, b = extractNumeral(b)
		if x != y {
			return lessNumeral(x, y)
		}
		// numeral section matched; return to matching text
	}
}

// extractNumeral extracts the numeral prefix of a.
// It returns the numeral and the remaining non-numeral part of a.
func extractNumeral(a string) (string, string) {
	for i, r := range a {
		if r < '0' || '9' < r {
			// split at this non-numeric rune
			return a[:i], a[i:]
		}
	}
	// a is entirely a numeral
	return a, ""
}

// lessNumeral compares two numerals in text form, sorting them as numbers
func lessNumeral(x, y string) bool {
	// the trick is that x and y only contain [0-9], so they can be treated as slices of bytes
	lx := len(x)
	ly := len(y)
	i := lx
	if i < ly {
		i = ly
	}
	for i > 0 {
		var xr, yr byte
		if i <= lx {
			xr = x[lx-i]
		}
		if i <= ly {
			yr = y[ly-i]
		}
		// xr,yr are the corresponding digts from x and y, or 0 (which is less than any rune in the '0'-'9' range)
		if xr != yr {
			return xr < yr
		}
		i--
	}
	// x and y are identical. thus they are not less-then
	// (note this case is not really reached, since the caller already checked that x != y, but in case this code
	// gets reused elsewhere it might as well be right)
	return false
}

// StringSlice is a slice of strings sortable in numeral-aware order
// It implements sort.Interface. The name matches the equivalent type in the standard sort package.
type StringSlice []string

func (s StringSlice) Len() int           { return len(s) }
func (s StringSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s StringSlice) Less(i, j int) bool { return Less(s[i], s[j]) }

// Strings is a utility function to sort a slice of strings using numeral-aware sort order.
// The name matched the name of the equivalent function in the standard sort package.
func Strings(strs []string) {
	sort.Sort(StringSlice(strs))
}
