package internal

import (
	utils "aot"
	"bufio"
	"strconv"
	"strings"
)

type Parser struct {
	seedLineSig          string
	functionSig          string
	functionLineValCount int
}

func NewParser() *Parser {
	return &Parser{
		seedLineSig:          "seeds: ",
		functionSig:          "map:",
		functionLineValCount: 3,
	}
}

// ParseFile - Parse input file & return transformation function map
func (p *Parser) ParseFile(scanner *bufio.Scanner) (tfMap FunctionSet, seeds []int) {
	seeds, tfMap = make([]int, 0), make(FunctionSet)

	// Current row & batch start indexes
	var cInd, bInd int

	for scanner.Scan() {
		cInd++

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// Parse seeds
		if len(seeds) == 0 {
			lineSplit := strings.Split(line, p.seedLineSig)
			if lineSplit == nil {
				panic("Failed to parse seed line...")
			}
			for _, strNum := range strings.Fields(lineSplit[1]) {
				seeds = append(seeds, utils.StrToInt(strNum))
			}
			continue
		}

		// Parse transformations
		if strings.Contains(line, p.functionSig) {
			bInd = cInd
			continue
		}
		// Parse line into a Transformation
		transformation := p.parseLine(line)
		if transformation != nil {
			tfMap[bInd] = append(tfMap[bInd], transformation)
		}
	}
	return tfMap, seeds
}

// parseLine - Parses transformation line
func (p *Parser) parseLine(line string) *Transformation {
	numbers := strings.Fields(line)
	if len(numbers) != p.functionLineValCount {
		return nil
	}

	destSt, _ := strconv.Atoi(numbers[0])
	sourceSt, _ := strconv.Atoi(numbers[1])
	rangeLength, _ := strconv.Atoi(numbers[2])

	return &Transformation{
		destSt:      destSt,
		sourceSt:    sourceSt,
		rangeLength: rangeLength,
	}
}
