/*
coords calculates coordinates for spreadsheets.
*/
package coords

import (
	"errors"
	"fmt"
)

const (
	maxCols      = 16384
	maxAlphabets = (maxCols / 26) - 1
)

// ColAlpha returns an alpha representation of a column index.
// index is an integer - ColAlpha(0) returns "A", etc.
func ColIndexToAlpha(index int) (string, error) {
	max := len(Colstream) - 1
	if index <= max {
		return Colstream[index], nil
	} else {
		msg := fmt.Sprintf("cannot have more than %d columns", max)
		return "", errors.New(msg)
	}
}

var Colstream = cols(maxAlphabets)

// ColLettersToIndex converts an alpha column
// reference to a zero-based numeric column identifier.
func ColAlphaToIndex(letters string) (int, error) {
	max := len(Colstream) - 1
	for i, v := range Colstream {
		if i > max {
			msg := fmt.Sprintf("Cannot exceed maximum of %d", max)
			return 0, errors.New(msg)
		}
		if v == letters {
			return i, nil
		}
	}
	return 0, errors.New("Cannot find requested string.")
}

//alphabet generates all the letters of the alphabet.
func alphabet() []string {
	letters := make([]string, 26)
	for idx := range letters {
		letters[idx] = string('A' + byte(idx))
	}
	return letters
}

//cols generates the alpha column component of Excel cell references
//Adds n alphabets to the first (A..Z) alphabet.
func cols(n int) []string {
	out := alphabet()
	alen := len(out)
	tmp := make([]string, alen)
	copy(tmp, out)
	for cycle := 0; cycle < n; cycle++ {
		for y := 0; y < alen; y++ {
			out = append(out, out[(cycle+2)-2]+tmp[y])
		}
	}
	return out
}
