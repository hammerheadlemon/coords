package coords

import (
	"testing"
)

func TestAlphaStream(t *testing.T) {
	if Colstream[26] != "AA" {
		t.Errorf("The test expected AA, got %v.", Colstream[26])
	}
	if len(Colstream) > maxCols {
		t.Errorf(`Number of columns in alphastream exceeds Excel maximum.
		alphastream contains %d, maxCols is %d`, len(Colstream), maxCols)
	}
}

func TestAlphaSingle(t *testing.T) {
	ab := alphabet()
	if ab[0] != "A" {
		t.Errorf("The test expected A, got %v.", ab[0])
	}
	if ab[1] != "B" {
		t.Errorf("The test expected B, got %v.", ab[1])
	}
	if ab[25] != "Z" {
		t.Errorf("The test expected Z, got %v.", ab[25])
	}
}

func TestAlphas(t *testing.T) {
	a := 2 // two alphabets long
	ecs := cols(a)
	cases := []struct {
		col int
		val string
	}{
		{0, "A"},
		{25, "Z"},
		{26, "AA"},
		{52, "BA"},
	}
	for _, c := range cases {
		// we're making sure we can pass that index
		r := 26 * a
		if c.col > r {
			t.Fatalf("Cannot use %d as index to array of %d", c.col, r)
		}
		if got := ecs[c.col]; got != c.val {
			t.Errorf("The test expected ecs[%d] to be %s - got %s.",
				c.col, c.val, ecs[c.col])
		}
	}
}

func TestCollectLetters(t *testing.T) {
	cases := []struct {
		alpha string
		index int
	}{
		{"XEZ", 16379},
		{"XEY", 16378},
		{"XEX", 16377},
		{"A", 0},
		{"B", 1},
		{"C", 2},
		{"AA", 26},
	}
	for _, c := range cases {
		s, err := colAlphaToIndex(c.alpha)
		if err != nil {
			t.Fatal(err)
		}
		if s != c.index {
			t.Errorf("Expected %d to return %s, instead it returned %d", c.index, c.alpha, s)
		}

	}
}

func TestGetColApha(t *testing.T) {
	cases := []struct {
		index int
		alpha string
	}{
		{0, "A"},
		{1, "B"},
		{2, "C"},
		{16377, "XEX"},
	}
	for _, c := range cases {
		s, err := colIndexToAlpha(c.index)
		if err != nil {
			t.Fatal(err)
		}
		if s != c.alpha {
			t.Errorf("Expected %s, got %s\n", c.alpha, s)
		}
	}
}
