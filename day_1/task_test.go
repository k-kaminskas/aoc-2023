package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask1(t *testing.T) {
	result := solution1("../day_1/dataset.txt")
	assert.Equal(t, 56042, result)
}

func TestTask2(t *testing.T) {
	result := solution2("../day_1/dataset.txt")
	assert.Equal(t, 55358, result)
}
