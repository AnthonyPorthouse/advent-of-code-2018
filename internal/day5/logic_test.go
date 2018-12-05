package day5

import "testing"

func TestReduceString(t *testing.T) {
	data := []struct{
		in []rune
		out []rune
	}{
		{[]rune("aA"), []rune("")},
		{[]rune("abBA"), []rune("")},
		{[]rune("abAB"), []rune("abAB")},
		{[]rune("aabAAB"), []rune("aabAAB")},
		{[]rune("dabAcCaCBAcCcaDA"), []rune("dabCBAcaDA")},
	}

	for _, tt := range data {
		t.Run(string(tt.in), func(t *testing.T) {
			s := ReduceString(tt.in)

			if string(s) != string(tt.out) {
				t.Errorf("got %q, want %q", s, tt.out)
			}
		})
	}
}

func TestFindSmallest(t *testing.T) {
	data := []struct{
		in []rune
		out int
	}{
		{[]rune("dabAcCaCBAcCcaDA"), 4},
	}

	for _, tt := range data {
		t.Run(string(tt.in), func(t *testing.T) {
			s := FindSmallest(tt.in)

			if s != tt.out {
				t.Errorf("got %q, want %q", s, tt.out)
			}
		})
	}
}