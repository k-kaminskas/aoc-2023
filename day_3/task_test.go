package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask1(t *testing.T) {
	result := solution1("../day_3/dataset.txt")
	assert.Equal(t, 527364, result)
}

func TestTask2(t *testing.T) {
	result := solution2("../day_3/dataset.txt")
	assert.Equal(t, 79026871, result)
}
