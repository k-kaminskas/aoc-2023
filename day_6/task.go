package day6

import (
	utils "aot"
	"aot/day_6/internal"
)

func solution1(fp string) (product int) {
	file, scanner := utils.GetScanner(fp)
	defer file.Close()

	races := internal.ParseFile(scanner)
	if races != nil {
		return races.Product()
	}
	return product
}

func solution2(fp string) (aProduct int) {
	file, scanner := utils.GetScanner(fp)
	defer file.Close()

	races := internal.ParseFile(scanner)
	if races != nil {
		return races.AggregatedProduct()
	}
	return aProduct
}
