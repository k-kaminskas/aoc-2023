package day1

import (
	utils "aot"
	"aot/day_1/internal"
	"strconv"
	"strings"
)

func solution(fp string, fn internal.DigitFunction) int {
	file, scanner := utils.GetScanner(fp)
	defer file.Close()

	totalSum := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		firstDigitR, lastDigitR := fn(line)

		// Compose a number & add to the final sum
		number, err := strconv.Atoi(firstDigitR + lastDigitR)
		if err != nil {
			panic(err)
		}
		totalSum += number
	}
	return totalSum
}

func solution1(fp string) int {
	return solution(fp, internal.GetDigits)
}

func solution2(fp string) int {
	return solution(fp, internal.ParseDigits)
}
