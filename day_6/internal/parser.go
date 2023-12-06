package internal

import (
	utils "aot"
	"bufio"
	"regexp"
	"strings"
)

func ParseFile(scanner *bufio.Scanner) Races {
	var times, distances []int
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(times) == 0 {
			times = parseNumbers(line)
			continue
		}
		distances = parseNumbers(line)
	}
	if len(times) != len(distances) {
		panic("Failed to parse the file")
	}

	var races []*Race
	for i := range times {
		races = append(races, &Race{
			Time: times[i], Distance: distances[i],
		})
	}
	return races
}

func parseNumbers(line string) (nums []int) {
	re := regexp.MustCompile(`\d+`)
	for _, match := range re.FindAllString(line, -1) {
		nums = append(nums, utils.StrToInt(match))
	}
	return nums
}
