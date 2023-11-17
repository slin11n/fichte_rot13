package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
)



func main() {
	// func 1
	// 1. Das programm bekommt einen Text und einen int um wie viel es rotieren soll
	// 2. dann gibt es den rotierten Text aus
	// func 2
	// 3. bekommt einen rotierten Text denn es entschlüsseln muss mit hilfe der wordlist
	// 4. ratiert den Text so lang wie das Alphabet ist 
	// 5. vergleicht jedes wort mit der wordlist und zählt welche rotation am meisten matches bekommt
	// 6. was die warscheinlichste übersetzung ist  

	wordList := readWordList("german_word_list.txt")
	wordMap := convertListToMap(wordList)
	var hits = make([]int, len(alphabet))
	text := readEncryptedTextFromUser()
	text = strings.TrimSpace(text)
	for shift := 0; shift < len(alphabet); shift++ {
		result := rot(text, alphabet, shift)
		hits[shift] = countWordsInLanguageMap(result, wordMap)
	}
	idx, maxHits := findMaxElementInArray(hits)
	decryptedText := rot(text, alphabet, idx)
	decryptedShfit := len(alphabet)-idx
	fmt.Printf("Entschlüsselter Text:\n%s\nDer Text wurde um %d Stellen rotiert. Es gab %d Treffer in der Wortliste.\n", decryptedText, decryptedShfit, maxHits)
}

var alphabet = []string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}



// readWordList reads file of words and gives back an arrays of strings
func readWordList(filePath string ) []string {
	words, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} 
	return strings.Split(string(words), "\n")
}

// convertListToMap converts the word list to a map/table
func convertListToMap(elements []string) map[string]bool {
    elementMap := make(map[string]bool)

    for _, s := range elements {  
       elementMap[strings.ToLower(s)] = true
    }
    return elementMap
}

// readEncryptedTextFromUser asks for the rotated string from the user that wants to be correctly rotated
func readEncryptedTextFromUser() string {
	fmt.Println("Bitte einen verschlüsselteten Text eingeben")
    reader := bufio.NewReader(os.Stdin)
    line, err := reader.ReadString('\n')
    if err != nil {
        os.Exit(1)
    }
 	return line
}

// countWordsInLanguageMap counts the words of every rotation by matches of the word list
func countWordsInLanguageMap(rotatedText string, wordMap map[string]bool) int {
	var counter int
	for _, word := range strings.Split(rotatedText, " ") {
		if wordMap[word] {
			counter++
		}

	}
	return counter
}

// findMaxElementInArray scans a given array for its highest value
// returns the index and the value it was in
func findMaxElementInArray(a []int) (int, int) {
	var max int
	var idx int
	for i := 0; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
			idx = i
		}
	}
	return idx, max
}



// rot it just rotates a string
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


