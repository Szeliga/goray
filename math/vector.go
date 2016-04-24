package math

// Vector - struct holding X Y Z values of a 3D vector
type Vector struct {
	X int
	Y int
	Z int
}

// Add - adds two vectors together
func (a Vector) Add(b Vector) Vector {
	return Vector{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

// Sub - subtracts b Vector from a Vector
func (a Vector) Sub(b Vector) Vector {
	return Vector{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}
