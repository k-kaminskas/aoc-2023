package internal

import (
	utils "aot"
	"fmt"
	"regexp"
	"strings"
)

type Parser struct {
	gameIDRegex    *regexp.Regexp
	batchDelimiter string
	itemDelimiter  string
	attrDelimiter  string
}

func NewParser() *Parser {
	return &Parser{
		gameIDRegex:    regexp.MustCompile(`Game (\d+):`),
		batchDelimiter: ";",
		itemDelimiter:  ",",
		attrDelimiter:  " ",
	}
}

func (p *Parser) ParseLine(line string) *Game {
	match := p.gameIDRegex.FindStringSubmatch(line)
	if match != nil {
		gameData := strings.TrimPrefix(line, match[0])
		if gameData != "" {
			return &Game{
				gameID:      utils.StrToInt(match[1]),
				cubeBatches: p.parseCubeBatch(gameData),
			}
		}
	}
	panic(fmt.Sprintf("Failed to parse the line - %s", line))
}

func (p *Parser) parseCubeBatch(gameData string) []*CubeBatch {
	var batches []*CubeBatch
	for en, batch := range strings.Split(gameData, p.batchDelimiter) {
		cubes := p.parseCubes(strings.TrimSpace(batch))
		batches = append(batches, &CubeBatch{
			batchID: en,
			cubes:   cubes,
		})
	}
	return batches
}

// Parses cubes data & extracts attributes
func (p *Parser) parseCubes(batchData string) Cubes {
	cubes := make(Cubes)
	for _, cube := range strings.Split(batchData, p.itemDelimiter) {
		cubeData := strings.Split(strings.TrimSpace(cube), p.attrDelimiter)
		if cubeData == nil {
			panic(fmt.Sprintf("Failed to parse cube data... Line - %s", cube))
		}
		countStr, color := cubeData[0], cubeData[1]
		cubes[color] = utils.StrToInt(countStr)
	}
	return cubes
}
