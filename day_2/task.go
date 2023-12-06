package day2

import (
	utils "aot"
	"aot/day_2/internal"
	"strings"
)

func solution(fp string, fn internal.SFunction) int {
	file, scanner := utils.GetScanner(fp)
	defer file.Close()

	totalSum := 0
	parser := internal.NewParser()
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Parse the game line
		game := parser.ParseLine(line)
		if game != nil {
			totalSum += fn(game)
		}
	}
	return totalSum
}

func solution1(fp string) int {
	return solution(fp, internal.GetID)
}

func solution2(fp string) int {
	return solution(fp, internal.GetScore)
}
