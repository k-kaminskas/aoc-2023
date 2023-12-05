package day2

import (
	utils "aot"
	"aot/day_2/internal"
	"strings"
)

func solution1(fp string) int {
	file, scanner := utils.GetScanner(fp)
	defer file.Close()

	totalSum := 0
	parser := internal.NewParser()
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Parse the game line
		game, err := parser.ParseLine(line)
		if err != nil {
			panic(err)
		}
		if game != nil {
			if game.IsPossible() {
				totalSum += game.GetID()
			}
		}
	}
	return totalSum
}
