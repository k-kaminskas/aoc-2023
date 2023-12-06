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

// ParseLine
// Parses game line & creates datasets
func (p *Parser) ParseLine(line string) (*Game, error) {
	if line == "" {
		return nil, fmt.Errorf("line is empty")
	}

	gameMatch := p.gameIDRegex.FindStringSubmatch(line)
	if gameMatch == nil {
		return nil, fmt.Errorf("no game number found in line: %s", line)
	}
	gameData := strings.TrimPrefix(line, gameMatch[0])
	if gameData == "" {
		return nil, fmt.Errorf("failed to extract game data from line: %s", line)
	}

	return &Game{
		gameID:      utils.StrToInt(gameMatch[1]),
		cubeBatches: p.parseCubeBatch(gameData),
	}, nil
}

// Parses game data line & extracts batches
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
