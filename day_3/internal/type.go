package internal

import (
	utils "aot"
	"fmt"
	"slices"
)

const GearSymbolSignifier = "*"
const GearNumberCount = 2

// DataIndex - K: starting column ID | V: entry value
type DataIndex map[int]string

// PartNumMap - K: gear symbol location in the matrix | V: adjacent part numbers
type PartNumMap map[string][]int

func (p PartNumMap) GetGearSum() (sum int) {
	for _, nums := range p {
		// A gear is any * symbol that is adjacent to exactly two part numbers
		if len(nums) == GearNumberCount {
			sum += nums[0] * nums[1]
		}
	}
	return sum
}

/* Matrix ----------------------------------------------------------------------------------------------------------- */

type Matrix []*Row

// Builds a partial matrix around a current row
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

// GetRowPartNumberSum - Collects all part numbers from a row
// based on ID & sums the numbers into a total amount.
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

// GetGearPartNumbersMap - Produces a symbol location ID for any
// matched "*" & aggregates all part numbers from all adjacent rows
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

/* Matrix row ------------------------------------------------------------------------------------------------------- */

type Row struct {
	ID              int
	DetectedNumbers DataIndex
	DetectedSymbols DataIndex
}

// GetAdjacentSymbolColumn - Gets adjacent symbol column for a given number & length
func (r Row) GetAdjacentSymbolColumn(colID, length int, filter ...string) (found bool, fColID int) {
	startCol, endCol := colID-1, colID+length
	for i := startCol; i <= endCol; i++ {
		if symbol, ok := r.DetectedSymbols[i]; ok {
			if len(filter) == 0 || slices.Contains(filter, symbol) {
				return true, i
			}
		}
	}
	return false, 0
}
