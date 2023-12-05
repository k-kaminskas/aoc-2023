package day5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask1(t *testing.T) {
	result := solution1("../day_5/dataset.txt")
	assert.Equal(t, 226172555, result)
}

func TestTask2(t *testing.T) {
	result := solution2("../day_5/dataset.txt")
	assert.Equal(t, 47909639, result)
}
