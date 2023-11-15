package main

import (
	"testing"
)

func TestRot(t *testing.T) {
	type test struct {
		testnumber int
		alphabet	[]string
		text string
		shift int
		want	string
	}

	var latin_alphabet = []string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}
	tests := []test{
			{ testnumber: 1, alphabet: []string{"a", "b", "c"}, text: "a", shift: 1, want: "b",},
			{ testnumber: 2, alphabet: []string{"a", "b", "c"}, text: "a", shift: 2, want: "c",},
			{ testnumber: 3, alphabet: []string{"a", "b", "c"}, text: "a", shift: 3, want: "a",},
			{ testnumber: 4, alphabet: []string{"a", "b", "c"}, text: "a", shift: 0, want: "a",},
			{ testnumber: 5, alphabet: []string{"a", "b", "c"}, text: "a", shift: 6, want: "a",},
			{ testnumber: 6, alphabet: []string{"a", "b", "c"}, text: "a", shift: 37704, want: "a",},
			{ testnumber: 7, alphabet: []string{"a", "b", "c"}, text: "a", shift: 37703, want: "c",},
			{ testnumber: 8, alphabet: []string{}, text: "a", shift: 1, want: "",},
			{ testnumber: 9, alphabet: latin_alphabet, text: "Nils", shift: 1, want: "ojmt",},
			{ testnumber: 10, alphabet: latin_alphabet, text: "Nils", shift: 26, want: "nils",},
			//{ testnumber: 11, alphabet: latin_alphabet, text: "Jörn", shift: 1, want: "k",},
	}

	for _, tc := range tests {
		got := rot(tc.text, tc.alphabet, tc.shift)
		if tc.want != got {
			t.Fatalf("expected: %v, got: %v in test: %d", tc.want, got, tc.testnumber)
		}
	}
}


