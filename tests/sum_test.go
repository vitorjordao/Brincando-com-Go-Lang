package test

import (
	"testing"
)

type SumTable struct {
	a      int
	b      int
	result int
}

func TestSoma(t *testing.T) {

	data := []SumTable{
		{a: 1, b: 2, result: 3},
		{a: 2, b: 2, result: 4},
		{a: 3, b: 2, result: 5},
	}

	for _, value := range data {
		total := Sum(value.a, value.b)

		if total != value.result {
			t.Errorf("The result is invalid")
		}
	}

}
