package day4

import (
	utils "aot"
	"aot/day_4/internal"
)

func solution1(fp string) (points int) {
	file, scanner := utils.GetScanner(fp)
	defer file.Close()

	cards := internal.NewParser().ParseFile(scanner)
	if cards != nil {
		return cards.GetPoints()
	}
	return points
}

func solution2(fp string) (rPoints int) {
	file, scanner := utils.GetScanner(fp)
	defer file.Close()

	cards := internal.NewParser().ParseFile(scanner)
	if cards != nil {
		return cards.GetRecursivePoints()
	}
	return rPoints
}
