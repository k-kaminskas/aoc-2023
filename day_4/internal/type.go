package internal

import (
	"slices"
)

type Card struct {
	cardID      int
	cardNums    []int
	winningNums []int
}

func (c *Card) getWinningNumberCount() (count int) {
	for _, n := range c.winningNums {
		if slices.Contains(c.cardNums, n) {
			count++
		}
	}
	return count
}

func (c *Card) GetCardPoints() int {
	count := c.getWinningNumberCount()
	if count > 0 {
		return 1 << (c.getWinningNumberCount() - 1)
	}
	return count
}

/* Card List -------------------------------------------------------------------------------------------------------- */

type CardList []*Card

func (cl CardList) GetPoints() int {
	sum := 0
	for _, c := range cl {
		sum += c.GetCardPoints()
	}
	return sum
}

func (cl CardList) GetRecursivePoints() int {
	instances := make(map[int]int)
	for _, card := range cl {
		instances[card.cardID]++
	}
	for _, card := range cl {
		for offset := 1; offset <= card.getWinningNumberCount(); offset++ {
			instances[card.cardID+offset] += instances[card.cardID]
		}
	}
	sum := 0
	for _, value := range instances {
		sum += value
	}
	return sum
}
