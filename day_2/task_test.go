package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask1(t *testing.T) {
	result := solution1("../day_2/dataset.txt")
	assert.Equal(t, 2207, result)
}

func TestTask2(t *testing.T) {
	result := solution2("../day_2/dataset.txt")
	assert.Equal(t, 62241, result)
}
