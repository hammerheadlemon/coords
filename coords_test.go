package coords

import (
	"testing"
)

// Test AlphaStream
func TestAlphaStream(t *testing.T) {
	if colstream[26] != "AA" {
		t.Errorf("The test expected AA, got %v.", colstream[26])
	}
	if len(colstream) > maxCols {
		t.Errorf(`Number of columns in alphastream exceeds Excel maximum.
		alphastream contains %d, maxCols is %d`, len(colstream), maxCols)
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
		s, err := ColIndexToAlpha(c.index)
		if err != nil {
			t.Fatal(err)
		}
		if s != c.alpha {
			t.Errorf("Expected %s, got %s\n", c.alpha, s)
		}
	}
}

func TestColAlphaToIndex(t *testing.T) {
	type args struct {
		letters string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"TestColAlphaToIndex", args{"XEZ"}, 16379, false},
		{"TestColAlphaToIndex", args{"XEY"}, 16378, false},
		{"TestColAlphaToIndex", args{"XEX"}, 16377, false},
		{"TestColAlphaToIndex", args{"A"}, 0, false},
		{"TestColAlphaToIndex", args{"B"}, 1, false},
		{"TestColAlphaToIndex", args{"AA"}, 26, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ColAlphaToIndex(tt.args.letters)
			if (err != nil) != tt.wantErr {
				t.Errorf("ColAlphaToIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ColAlphaToIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
