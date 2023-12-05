package day2

import (
	utils "aot"
	"aot/day_2/internal"
	"strings"
)

func solution2(fp string) int {
	file, scanner := utils.GetScanner(fp)
	defer file.Close()

	totalScore := 0
	parser := internal.NewParser()
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Parse the game line
		game, err := parser.ParseLine(line)
		if err != nil {
			panic(err)
		}

		if game != nil {
			gameScore := 1
			cubesUsed := game.GetCubesUsed()
			for _, count := range cubesUsed {
				gameScore *= count
			}
			totalScore += gameScore
		}
	}
	return totalScore
}
