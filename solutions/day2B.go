package solutions

import (
	"fmt"
	"strings"

	"github.com/bradisbon/AoC-2022/reader"
)

var rPSOutcomes = map[string]map[string]int {
	"A": {
		"X": 3,
		"Y": 1,
		"Z": 2,
	},
	"B": {
		"X": 1,
		"Y": 2,
		"Z": 3,
	},
	"C": {
		"X": 2,
		"Y": 3,
		"Z": 1,
	},
}

var secondLetterScores= map[string]int {
	"X": 0,
	"Y": 3,
	"Z": 6,
}

// score for second letter of input + score for hand you pick
func Day2A() {
	scanner := reader.Reader("inputs/day2A.txt")
	totalScore := 0
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), " ")
		totalScore += rPSOutcomes[row[0]][row[1]] + secondLetterScores[row[1]]
	}

	fmt.Printf("Total score: %v\n", totalScore)
}
