package solutions

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bradisbon/AoC-2022/reader"
)

func Day4A() {
	scanner := reader.Reader("inputs/day4.txt")

	overlaps := 0
	for scanner.Scan() {
		pairs := strings.Split(scanner.Text(),",")
		pairA := strings.Split(pairs[0],"-")
		pairB := strings.Split(pairs[1],"-")
		startA, _ := strconv.Atoi(pairA[0])
		endA, _ := strconv.Atoi(pairA[1])
		startB, _ := strconv.Atoi(pairB[0])
		endB, _ := strconv.Atoi(pairB[1])
		if (startA <= startB && endA >= endB) || (startB <= startA && endB >= endA) {
			overlaps++
		}
	}

	fmt.Printf("overlaps: %v\n",overlaps)
}
