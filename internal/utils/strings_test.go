package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizingString(t *testing.T) {
	normalized := NormalizeString("Hello WoRld")
	assert.Equal(t, "hello_world", normalized)
}

func TestCalculatingLevenshteinDistance(t *testing.T) {
	cost := CalculateLevenshteinDistance("capital one", "capitalone")
	assert.Equal(t, 1, cost)
}
