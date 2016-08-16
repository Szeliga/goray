package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	result := Vector{1., 1., 1.}.Add(Vector{2., 2., 2.})
	assert.Equal(t, Vector{3., 3., 3.}, result, "should add correctly")
}

func TestSub(t *testing.T) {
	result := Vector{3., 3., 3.}.Sub(Vector{1., 1., 1.})
	assert.Equal(t, Vector{2., 2., 2.}, result, "should subtract correcly")
}

func TestMultiplyByScalar(t *testing.T) {
	result := Vector{3., 3., 3.}.MultiplyByScalar(2.)
	assert.Equal(t, Vector{6., 6., 6.}, result, "should multiply each component by given scalar")
}

func TestLength(t *testing.T) {
	assert.Equal(t, 3., Vector{3., 0., 0.}.Length(), "calculates the length(magnitude) correctly for Vector{3., 0., 0.}")
	assert.Equal(t, 6., Vector{6., 0., 0.}.Length(), "calculates the length(magnitude) correctly for Vector{6., 0., 0.}")
	assert.Equal(t, 6.324555320336759, Vector{6., 2., 0.}.Length(), "calculates the length(magnitude) correctly for Vector{6., 2., 0.}")
}

func TestDotProductOfTwoPerpendicularVectors(t *testing.T) {
	result := Vector{1., 0., 0.}.Dot(Vector{0., 1., 0.})
	assert.Equal(t, 0., result, "dot product of two perpendicular vectors is 0")
}

func TestDotProductOfTwoParallelVectors(t *testing.T) {
	result := Vector{1., 0., 0.}.Dot(Vector{1., 0., 0.})
	assert.Equal(t, 1., result, "dot product of two parallel vectors is 1")
}

func TestCrossProduct(t *testing.T) {
	result := Vector{1., 0., 0.}.Cross(Vector{0., 1., 0.})
	assert.Equal(t, Vector{0., 0., 1.}, result, "cross product returns a vector perpendicular to the other two")
}

func TestNormalize(t *testing.T) {
	assert.Equal(t, Vector{1., 0., 0.}, Vector{10., 0., 0.}.Normalize(), "returns a unit vector (versor) from the given vector")
	assert.Equal(t, Vector{0., 1., 0.}, Vector{0., 10., 0.}.Normalize(), "returns a unit vector (versor) from the given vector")
	assert.Equal(t, Vector{0., 0., 1.}, Vector{0., 0., 10.}.Normalize(), "returns a unit vector (versor) from the given vector")
}
