package day3

import (
	utils "aot"
	"aot/day_3/internal"
)

func solution1(fp string) int {
	file, scanner := utils.GetScanner(fp)
	defer file.Close()

	totalSum := 0
	matrix := internal.NewMatrix(scanner)
	for rowID := range matrix {
		totalSum += matrix.GetRowPartNumberSum(rowID)
	}

	return totalSum
}
