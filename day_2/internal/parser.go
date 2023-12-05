package internal

import (
	"fmt"
	"regexp"
	"strconv"
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
// Parses game data line & creates datasets based on delimiter
// returns a parsed game struct pointer
func (p *Parser) ParseLine(line string) (*Game, error) {
	if line == "" {
		return nil, fmt.Errorf("line is empty")
	}

	gameMatch := p.gameIDRegex.FindStringSubmatch(line)
	if gameMatch == nil {
		return nil, fmt.Errorf("no game number found in line: %s", line)
	}
	gameID, err := strconv.Atoi(gameMatch[1])
	if err != nil {
		return nil, fmt.Errorf("failed to parse game number: %v", err)
	}
	gameData := strings.TrimPrefix(line, gameMatch[0])
	if gameData == "" {
		return nil, fmt.Errorf("failed to extract game data from line: %s", line)
	}

	return &Game{
		gameID:      gameID,
		cubeBatches: p.parseCubeBatchLine(gameData),
	}, nil
}

// Parses game data line & extracts batches based on delimiter
// returns a slice of parsed cube batch slice
func (p *Parser) parseCubeBatchLine(gameData string) []*CubeBatch {
	var batches []*CubeBatch
	for en, batch := range strings.Split(gameData, p.batchDelimiter) {
		cubes := p.parseCubeLine(strings.TrimSpace(batch))
		batches = append(batches, &CubeBatch{
			batchID: en,
			cubes:   cubes,
		})
	}
	return batches
}

// Parses cube line & extracts attributes based on delimiter
// returns a slice of parsed cube map
func (p *Parser) parseCubeLine(batchData string) Cubes {
	cubes := make(Cubes)
	for _, cube := range strings.Split(batchData, p.itemDelimiter) {
		cubeData := strings.Split(strings.TrimSpace(cube), p.attrDelimiter)
		if cubeData == nil {
			panic(fmt.Sprintf("Failed to parse cube data... Line - %s", cube))
		}
		countStr, color := cubeData[0], cubeData[1]

		count, err := strconv.Atoi(countStr)
		if err != nil {
			panic(err)
		}
		cubes[color] = count
	}
	return cubes
}
