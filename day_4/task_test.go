package day4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask1(t *testing.T) {
	result := solution1("../day_4/dataset.txt")
	assert.Equal(t, 20407, result)
}

func TestTask2(t *testing.T) {
	result := solution2("../day_4/dataset.txt")
	assert.Equal(t, 23806951, result)
}
