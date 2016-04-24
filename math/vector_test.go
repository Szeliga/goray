package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	result := Vector{1, 1, 1}.Add(Vector{2, 2, 2})
	assert.Equal(t, Vector{3, 3, 3}, result, "should add correctly")
}

func TestSub(t *testing.T) {
	result := Vector{3, 3, 3}.Sub(Vector{1, 1, 1})
	assert.Equal(t, Vector{2, 2, 2}, result, "should subtract correcly")
}
