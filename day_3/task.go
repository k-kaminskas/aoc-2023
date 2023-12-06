package day3

import (
	utils "aot"
	"aot/day_3/internal"
)

func solution1(fp string) int {
	file, scanner := utils.GetScanner(fp)
	defer file.Close()

	totalSum := 0
	matrix := internal.NewParser().ParseFile(scanner)

	for rowID := range matrix {
		totalSum += matrix.GetRowPartNumberSum(rowID)
	}
	return totalSum
}

func solution2(fp string) int {
	file, scanner := utils.GetScanner(fp)
	defer file.Close()

	aggregatedMap := make(internal.PartNumMap)
	matrix := internal.NewParser().ParseFile(scanner)

	for rowID := range matrix {
		partNumsMap := matrix.GetGearPartNumbersMap(rowID)
		for symID, nums := range partNumsMap {
			aggregatedMap[symID] = append(aggregatedMap[symID], nums...)
		}
	}

	return aggregatedMap.GetGearSum()
}
