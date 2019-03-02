package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	input         int
	desiredResult [][]int
}

var td []Test

func TestTheQuestion(t *testing.T) {
	x := make([][]int, 0)
	x = append(x, []int{1, 1, 1, 1})
	x = append(x, []int{2, 1, 1})
	x = append(x, []int{1, 2, 1})
	x = append(x, []int{1, 1, 2})
	x = append(x, []int{2, 2})

	td = append(td, Test{
		4,
		x,
	})

	for _, v := range td {
		result := theQuestion(v.input)

		assert.Equalf(
			t,
			v.desiredResult,
			result,
			"Expected %v, got %v\n", v.desiredResult, result,
		)
	}
}
