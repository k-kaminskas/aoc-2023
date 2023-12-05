package internal

import "fmt"

/* Max cube combinations for a game */

var maxColorCombos = Cubes{
	"red":   12,
	"green": 13,
	"blue":  14,
}

type Cubes map[string]int

type CubeBatch struct {
	batchID int
	cubes   Cubes
}

type Game struct {
	gameID      int
	cubeBatches []*CubeBatch
}

// IsPossible - Checks whether current game is possible
// by matching it against static max color combinations
func (g *Game) IsPossible() bool {
	for _, batch := range g.cubeBatches {
		for color, count := range batch.cubes {
			threshold, ok := maxColorCombos[color]
			if !ok {
				panic(fmt.Sprintf("Unrecognized color %s", color))
			}
			if count > threshold {
				return false
			}
		}
	}
	return true
}

// GetCubesUsed - Returns the amount of cubes used between all the batches of a game
// (Smallest amount of cubes for each color so all batch variations are possible)
func (g *Game) GetCubesUsed() Cubes {
	cubesUsed := make(Cubes)
	for _, batch := range g.cubeBatches {
		for color, count := range batch.cubes {
			if cubesUsed[color] < count {
				cubesUsed[color] = count
			}
		}
	}
	return cubesUsed
}

// GetID - returns ID for current game
func (g *Game) GetID() int {
	return g.gameID
}
