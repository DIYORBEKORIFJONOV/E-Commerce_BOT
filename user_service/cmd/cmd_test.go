package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Sum(a, b int) int {
	return a + b
}

func TestSum(t *testing.T) {
	result := Sum(2, 3)
	assert.Equal(t, 5, result, "Сумма должна быть 5")
}
