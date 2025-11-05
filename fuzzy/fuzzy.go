package fuzzy

import (
	"bufio"
	"io"
	"strings"
)

func FuzzyFind(w io.Reader, searchTerm string) (int, []int) {
	var linesFounded []int
	totalOccurence := 0
	currentLine := 0

	input := bufio.NewScanner(w)
	for input.Scan() {
		currentLine++

		line := input.Text()

		countOnLine := strings.Count(line, searchTerm)
		if countOnLine > 0 {
			linesFounded = append(linesFounded, currentLine)
			totalOccurence += countOnLine
		}
	}
	return totalOccurence, linesFounded
}
