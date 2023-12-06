package internal

import (
	"bufio"
	"strings"
	"unicode"
)

type Parser struct {
	nonDataSymbol rune
}

func NewParser() *Parser {
	return &Parser{
		nonDataSymbol: '.',
	}
}

func (p *Parser) ParseFile(scanner *bufio.Scanner) Matrix {
	matrix := make(Matrix, 0)
	rowNum := 0
	for scanner.Scan() {
		rowNum++
		line := strings.TrimSpace(scanner.Text())

		// Gather numbers & symbols from the current row
		nums, symbols := make(DataIndex), make(DataIndex)
		startColID, fragments := 0, make([]rune, 0)
		for i, char := range line {
			colID := i + 1
			isDigit, eol := unicode.IsDigit(char), colID == len(line)
			if isDigit {
				if len(fragments) == 0 {
					startColID = colID
				}
				fragments = append(fragments, char)
			}
			if len(fragments) != 0 && (eol || !isDigit) {
				nums[startColID] = string(fragments)
				startColID, fragments = 0, make([]rune, 0)
			}
			if !isDigit && char != p.nonDataSymbol {
				symbols[colID] = string(char)
			}
		}
		matrix = append(matrix, &Row{
			ID:              rowNum,
			DetectedNumbers: nums,
			DetectedSymbols: symbols,
		})
	}
	return matrix
}
