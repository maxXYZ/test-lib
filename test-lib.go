package test_lib

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

// Add two integers. See more on [google]
// [google]: https://google.com
func Add[I int](x I, y I) I {
	return x + y
}