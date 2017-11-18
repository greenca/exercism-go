// Package scrabble implements a scoring function for a scrabble string
package scrabble

import "strings"

type letterScore struct {
	letters string
	value int
}

var (
	scorevalues = []letterScore{
		{"A, E, I, O, U, L, N, R, S, T", 1},
		{"D, G", 2}, 
		{"B, C, M, P", 3}, 
		{"F, H, V, W, Y", 4}, 
		{"K", 5}, 
		{"J, X", 8}, 
		{"Q, Z", 10}, 
	}
	scoremap = mapLetterValues()
)

// mapLetterValues converts the list of letter groups and scores to a map from any letter to its score
func mapLetterValues() map[string]int {
	lettermap := map[string]int{}
	for _, group := range scorevalues {
		for _, letter := range strings.Split(group.letters, ", ") {
			lettermap[letter] = group.value
		}
	}
	return lettermap
}

// Score takes a string and returns it's score if it were played in a game of scrabble
func Score(word string) int {
	score := 0
	for _, letter := range strings.Split(word, "") {
		score += scoremap[strings.ToUpper(letter)]
	}
	return score
}
