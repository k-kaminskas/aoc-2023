package day1

import (
	utils "aot"
	"strconv"
	"strings"
)

var numMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
	"1":     "1",
	"2":     "2",
	"3":     "3",
	"4":     "4",
	"5":     "5",
	"6":     "6",
	"7":     "7",
	"8":     "8",
	"9":     "9",
}

// Loops through the number map & returns first
// & last found digits based on indexes
func getDigits(s string) (firstDigit, lastDigit string) {
	var minIndex, maxIndex int
	var minKey, maxKey string

	for key := range numMap {
		ind := strings.Index(s, key)
		if ind != -1 && (minKey == "" || ind < minIndex) {
			minIndex, minKey = ind, key
		}
		ind = strings.LastIndex(s, key)
		if ind != -1 && (maxKey == "" || ind > maxIndex) {
			maxIndex, maxKey = ind, key
		}
	}
	return numMap[minKey], numMap[maxKey]
}

// Solution for the second task
func solution2(fp string) int {
	file, scanner := utils.GetScanner(fp)
	defer file.Close()

	totalSum := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		firstDigit, lastDigit := getDigits(line)

		// Compose a number & add to the final sum
		number, err := strconv.Atoi(firstDigit + lastDigit)
		if err != nil {
			panic(err)
		}
		totalSum += number
	}
	return totalSum
}
