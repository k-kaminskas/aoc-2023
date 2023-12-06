package internal

import (
	"strings"
	"unicode"
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

// GetDigits
// Loops through the string & returns first & last digit runes
func GetDigits(s string) (firstDigit, lastDigit rune) {
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

// ParseDigits
// Loops through the number map & parses first
// & last found digits based on indexes
func ParseDigits(s string) (firstDigit, lastDigit string) {
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
