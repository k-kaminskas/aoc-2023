package day3

import (
	utils "aot"
	"aot/day_3/internal"
)

const AdjacentNumberCount = 2

func solution2(fp string) int {
	file, scanner := utils.GetScanner(fp)
	defer file.Close()

	matrix := internal.NewMatrix(scanner)
	aggregatedMap := make(internal.PartNumMap)
	for rowID := range matrix {
		partNumsMap := matrix.GetGearPartNumbersMap(rowID)
		for symID, nums := range partNumsMap {
			aggregatedMap[symID] = append(aggregatedMap[symID], nums...)
		}
	}

	totalSum := 0
	for _, nums := range aggregatedMap {
		// A gear is any * symbol that is adjacent to exactly two part numbers
		if len(nums) == AdjacentNumberCount {
			totalSum += nums[0] * nums[1]
		}
	}

	return totalSum
}
