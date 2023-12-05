package internal

import "slices"

/* Matrix row */

type Row struct {
	ID              int
	DetectedNumbers DataIndex
	DetectedSymbols DataIndex
}

// GetAdjacentSymbolColumn - Gets adjacent symbol column for a given number & length.
// Symbols can be filtered using optional parameter.
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
