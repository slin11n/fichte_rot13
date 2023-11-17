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
			{ testnumber: 11, alphabet: latin_alphabet, text: "Ich gehe morgen zum Zahnarzt", shift: 3, want: "lfk jhkh prujhq cxp cdkqducw",},
	}		

	for _, tc := range tests {
		got := rot(tc.text, tc.alphabet, tc.shift)
		if tc.want != got {
			t.Fatalf("expected: %v, got: %v in test: %d", tc.want, got, tc.testnumber)
		}
	}
}


func TestFindMaxElementInArray(t *testing.T) {
	type test struct {
		testnumber int
		a []int
		wantIdx int
		wantValue int
	}

	tests := []test{
			{ testnumber: 1, a: []int{1,2,3}, wantIdx: 2, wantValue: 3},
			{ testnumber: 2, a: []int{1,2,3,99,77,0,11}, wantIdx: 3, wantValue: 99},
			{ testnumber: 3, a: []int{1}, wantIdx: 0, wantValue: 1},
			{ testnumber: 4, a: []int{99,3,2,1,99}, wantIdx: 0, wantValue: 99},
			{ testnumber: 5, a: []int{99,4,2,1}, wantIdx: 0, wantValue: 99},
	}		

	for _, tc := range tests {
		gotIdx, gotValue := findMaxElementInArray(tc.a)
		if tc.wantIdx != gotIdx || tc.wantValue != gotValue {
			t.Fatalf("expected: idx %d val %d, got: idx %d val %d in test: %d", tc.wantIdx, tc.wantValue, gotIdx, gotValue, tc.testnumber)
		}
	}
}


