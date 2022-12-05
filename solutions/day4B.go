package solutions

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bradisbon/AoC-2022/reader"
)

func parsePairs(pairS string) [][]int {
	pairs := strings.Split(pairS,",")
	pairA := strings.Split(pairs[0],"-")
	pairB := strings.Split(pairs[1],"-")
	startA, _ := strconv.Atoi(pairA[0])
	endA, _ := strconv.Atoi(pairA[1])
	startB, _ := strconv.Atoi(pairB[0])
	endB, _ := strconv.Atoi(pairB[1])

	return [][]int{{startA,endA},{startB,endB}}
}

func Day4B() {
	scanner := reader.Reader("inputs/day4.txt")

	overlaps := 0
	for scanner.Scan() {
		pairs := parsePairs(scanner.Text())
		startA := pairs[0][0]
		endA := pairs[0][1]
		startB := pairs[1][0]
		endB := pairs[1][1]
		if (startB >= startA && startB <= endA) || (startA >= startB && startA <= endB) {
			overlaps++
		}
	}

	fmt.Printf("overlaps: %v\n",overlaps)
}
