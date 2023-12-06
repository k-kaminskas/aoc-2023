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

/* Game ------------------------------------------------------------------------------------------------------------- */

type Game struct {
	gameID      int
	cubeBatches []*CubeBatch
}

// IsPossible - Checks whether current game is possible
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

/* Solution Functions ---------------------------------------------------------------------------------------------- */

type SFunction func(g *Game) int

func GetID(game *Game) (id int) {
	if game.IsPossible() {
		return game.gameID
	}
	return id
}

func GetScore(game *Game) (score int) {
	score = 1
	cubesUsed := game.GetCubesUsed()
	for _, count := range cubesUsed {
		score *= count
	}
	return score
}
