package internal

import (
	utils "aot"
	"bufio"
	"fmt"
	"strings"
	"unicode"
)

const NonDataSymbol rune = '.'
const GearSymbolSignifier = "*"

// DataIndex - K: starting column ID | V: entry value
type DataIndex map[int]string

// PartNumMap - K: gear symbol location in the matrix | V: adjacent part numbers
type PartNumMap map[string][]int

type Matrix []*Row

// NewMatrix - Reads the input file and initializes an indexed matrix with relevant data
func NewMatrix(scanner *bufio.Scanner) Matrix {
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
			if !isDigit && char != NonDataSymbol {
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

// buildAdjacentRowMatrix - Builds a partial matrix around
// the current row, including the above & below rows
func (m Matrix) buildAdjacentRowMatrix(rowID int, row *Row) Matrix {
	ma := Matrix{row}
	if rowID-1 >= 0 {
		ma = append(ma, m[rowID-1])
	}
	if rowID+1 < len(m) {
		ma = append(ma, m[rowID+1])
	}
	return ma
}

/* Entrypoints */

// GetRowPartNumberSum - Collects all part numbers from a row
// based on ID & sums the numbers into a total amount.
// A *part* number is a number with any adjacent symbols
func (m Matrix) GetRowPartNumberSum(cRowID int) int {
	cRow, sum := m[cRowID], 0
	for numColID, number := range cRow.DetectedNumbers {
		for _, row := range m.buildAdjacentRowMatrix(cRowID, cRow) {
			adjacent, _ := row.GetAdjacentSymbolColumn(numColID, len(number))
			if adjacent {
				sum += utils.StrToInt(number)
			}
		}
	}
	return sum
}

// GetGearPartNumbersMap - Loops through all the numbers in the row &
// tries to find any adjacent gear symbols. Produces a symbol location ID
// for any matched case & adds up all part numbers from all adjacent rows
func (m Matrix) GetGearPartNumbersMap(cRowID int) PartNumMap {
	cRow, adjacentNums := m[cRowID], make(PartNumMap)
	for numColID, number := range cRow.DetectedNumbers {
		for _, row := range m.buildAdjacentRowMatrix(cRowID, cRow) {
			adjacent, matchedColID := row.GetAdjacentSymbolColumn(
				numColID, len(number), GearSymbolSignifier,
			)
			if adjacent {
				symID := fmt.Sprintf("%v-%v", row.ID, matchedColID)
				adjacentNums[symID] = append(adjacentNums[symID], utils.StrToInt(number))
			}
		}
	}
	return adjacentNums
}
