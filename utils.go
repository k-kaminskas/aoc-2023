package aot

import (
	"bufio"
	"os"
	"strconv"
)

// GetScanner - Opens a file based on a relative path &
// returns a scanner object. Returned file is expected to be closed with defer
func GetScanner(relativePath string) (*os.File, *bufio.Scanner) {
	// Open the file & get the scanner
	file, err := os.Open(relativePath)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	return file, scanner
}

// StrToInt - Converts string number to integer
func StrToInt(number string) int {
	value, err := strconv.Atoi(number)
	if err != nil {
		panic(err)
	}
	return value
}
