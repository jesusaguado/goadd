package goadd

import "golang.org/x/exp/constraints"

// Addable provides an interface to anything that makes sense to add in common
// programs (excluding complex analysis of course).
type Addable interface {
	constraints.Float | constraints.Integer
}

// Add takes two addables and adds them together.
func Add[T Addable](a, b T) T {
	return a + b
}
