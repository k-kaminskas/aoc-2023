package internal

import (
	utils "aot"
	"bufio"
	"fmt"
	"regexp"
	"strings"
)

type Parser struct {
	cardIDRegex    *regexp.Regexp
	batchDelimiter string
}

func NewParser() *Parser {
	return &Parser{
		cardIDRegex:    regexp.MustCompile(`Card\s*(\d+):`),
		batchDelimiter: "|",
	}
}

func (p *Parser) ParseFile(scanner *bufio.Scanner) (cards CardList) {
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		match := p.cardIDRegex.FindStringSubmatch(line)
		if match != nil {
			cardData := strings.TrimPrefix(line, match[0])
			parts := strings.Split(cardData, p.batchDelimiter)
			if parts != nil {
				cards = append(cards, &Card{
					cardID:      utils.StrToInt(match[1]),
					cardNums:    p.parseNumbers(parts[1]),
					winningNums: p.parseNumbers(parts[0]),
				})
				continue
			}
		}
		panic(fmt.Sprintf("Failed to parse the line - %s", line))
	}
	return cards
}

func (p *Parser) parseNumbers(s string) (numbers []int) {
	for _, strNum := range strings.Fields(strings.TrimSpace(s)) {
		numbers = append(numbers, utils.StrToInt(strNum))
	}
	return numbers
}
