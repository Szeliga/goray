package core

import "math"

// Vector - struct holding X Y Z values of a 3D vector
type Vector struct {
	X, Y, Z float64
}

// Add - adds two vectors together
func (a Vector) Add(b Vector) Vector {
	return Vector{
		X: a.X + b.X,
		Y: a.Y + b.Y,
		Z: a.Z + b.Z,
	}
}

// Sub - subtracts b Vector from a Vector
func (a Vector) Sub(b Vector) Vector {
	return Vector{
		X: a.X - b.X,
		Y: a.Y - b.Y,
		Z: a.Z - b.Z,
	}
}

// MultiplyByScalar - multiplies a Vector by s float64
func (a Vector) MultiplyByScalar(s float64) Vector {
	return Vector{
		X: a.X * s,
		Y: a.Y * s,
		Z: a.Z * s,
	}
}

// Length - calculates the length(magnitude) of the Vector
func (a Vector) Length() float64 {
	return math.Sqrt(a.Dot(a))
}

// Dot - calculates the dot product of two Vectors
func (a Vector) Dot(b Vector) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

// Cross - calculates the cross product of two Vectors
func (a Vector) Cross(b Vector) Vector {
	return Vector{
		X: a.Y*b.Z - a.Z*b.Y,
		Y: a.Z*b.X - a.X*b.Z,
		Z: a.X*b.Y - a.Y*b.X,
	}
}

// Normalize - returns a versor created from the given vector
func (a Vector) Normalize() Vector {
	return a.MultiplyByScalar(1. / a.Length())
}
