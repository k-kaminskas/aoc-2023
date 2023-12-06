package day5

import (
	utils "aot"
	"aot/day_5/internal"
	"slices"
)

func solution1(fp string) int {
	file, scanner := utils.GetScanner(fp)
	defer file.Close()

	tfMap, seeds := internal.NewParser().ParseFile(scanner)

	trSeeds := make([]int, 0)
	for _, seed := range seeds {
		for _, key := range tfMap.GetSortedKeys() {
			function, ok := tfMap[key]
			if !ok {
				continue
			}
			seed = function.PerformTransformations(seed)
		}
		trSeeds = append(trSeeds, seed)
	}
	return slices.Min(trSeeds)
}

func solution2(fp string) int {
	file, scanner := utils.GetScanner(fp)
	defer file.Close()

	ranges := make([]*internal.Range, 0)
	tfMap, seeds := internal.NewParser().ParseFile(scanner)

	for _, seed := range internal.ChunkIntSliceToPairs(seeds) {
		range_ := []*internal.Range{{Start: seed[0], End: seed[0] + seed[1]}}
		for _, key := range tfMap.GetSortedKeys() {
			function, ok := tfMap[key]
			if !ok {
				continue
			}
			range_ = function.PerformRangeTransformations(range_...)
		}
		ranges = append(ranges, range_...)
	}

	minValue := 0
	for _, range_ := range ranges {
		if minValue == 0 || range_.Start < minValue {
			minValue = range_.Start
		}
	}
	return minValue
}
