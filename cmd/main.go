package main

import (
	"fmt"
	"strings"
)

func main() {
		text := "Nutella ohne Butter ist besser als Nutella mit Butter"
		shift := 17
		result := rot(text, alph, shift)
		fmt.Printf("%s rotiert um %d ergibt %s\n", text, shift, result)
}

var alph = []string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}

// rot takes a string, an alphabet and a rotation (shift) value
// returns the string rotated by rotation value on the alphabet
// If alphabet is empty rot returns an empty string
// If an character in the input text does not exists in the alphabet
// the character is omitted in the result
func rot(text string, alphabet []string, shift int) string {
	if len(alphabet) == 0 {
		return ""
	}
	text = strings.ToLower(text)
	var result string
	var idx int 
	for _, t := range text {
		if string(t) == " " {
			result = result + " "
			continue
		}
		for j, alpha := range alphabet {
			if alpha == string(t) {
				idx = j
			}
		}
		tmpShift := shift
		for idx + tmpShift > len(alphabet) - 1 {
			tmpShift = tmpShift - len(alphabet)
		}
		result = result + alphabet[idx + tmpShift]
	}
	return result
}