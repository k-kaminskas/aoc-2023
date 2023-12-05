package day1

import (
	utils "aot"
	"aot/day_1/internal"
	"strconv"
	"strings"
)

func solution1(fp string) int {
	file, scanner := utils.GetScanner(fp)
	defer file.Close()

	totalSum := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		firstDigitR, lastDigitR := internal.GetDigits(line)

		// Compose a number & add to the final sum
		number, err := strconv.Atoi(string(firstDigitR) + string(lastDigitR))
		if err != nil {
			panic(err)
		}
		totalSum += number
	}
	return totalSum
}

func solution2(fp string) int {
	file, scanner := utils.GetScanner(fp)
	defer file.Close()

	totalSum := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		firstDigit, lastDigit := internal.ParseDigits(line)

		// Compose a number & add to the final sum
		number, err := strconv.Atoi(firstDigit + lastDigit)
		if err != nil {
			panic(err)
		}
		totalSum += number
	}
	return totalSum
}
