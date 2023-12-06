package day6

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask1(t *testing.T) {
	result := solution1("../day_6/dataset.txt")
	assert.Equal(t, 4811940, result)
}

func TestTask2(t *testing.T) {
	result := solution2("../day_6/dataset.txt")
	assert.Equal(t, 30077773, result)
}
