package goadd

import "golang.org/x/exp/constraints"

// Number provides an interface to anything that makes sense to add in common
// programs (excluding complex analysis of course).
type Number interface {
	constraints.Float | constraints.Integer
}

// Add takes two addables and adds them together.
func Add[T Number](a, b T) T {
	return a + b
}
