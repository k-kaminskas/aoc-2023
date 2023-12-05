package day1

import (
	utils "aot"
	"strconv"
	"strings"
	"unicode"
)

// Loops through the string & returns first & last digit runes
func getDigitRunes(s string) (firstDigit, lastDigit rune) {
	for _, r := range s {
		if unicode.IsDigit(r) {
			if firstDigit == 0 {
				firstDigit = r
			}
			lastDigit = r
		}
	}
	return firstDigit, lastDigit
}

// Solution for the first task
func solution1(fp string) int {
	file, scanner := utils.GetScanner(fp)
	defer file.Close()

	totalSum := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		firstDigitR, lastDigitR := getDigitRunes(line)

		// Compose a number & add to the final sum
		number, err := strconv.Atoi(string(firstDigitR) + string(lastDigitR))
		if err != nil {
			panic(err)
		}
		totalSum += number
	}
	return totalSum
}
