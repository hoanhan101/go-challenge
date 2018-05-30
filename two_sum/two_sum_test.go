package two_sum

import (
	"reflect"
	"testing"
)

const (
	succeed = "\u2713"
	failed  = "\u2717"
)

func TestTwoSumI(t *testing.T) {
	tests := [][]int{
		[]int{2, 7, 11, 15},
		[]int{0, 9, 7, 3, 4, 10},
		[]int{1, 2, 8, 10},
	}

	targets := []int{
		9,
		19,
		100,
	}

	expectedResults := [][]int{
		[]int{0, 1},
		[]int{1, 5},
		[]int{0, 0},
	}

	for i := range tests {
		result := twoSumI(tests[i], targets[i])
		if reflect.DeepEqual(expectedResults[i], result) == true {
			t.Logf("\t%s\tCase: %v\tTarget: %v\tResult: %v", succeed, tests[i], targets[i], result)
		} else {
			t.Errorf("\t%s\tCase: %v\tTarget: %v\tResult: %v", failed, tests[i], targets[i], result)
		}
	}
}

func TestTwoSumII(t *testing.T) {
	tests := [][]int{
		[]int{2, 7, 11, 15},
		[]int{0, 3, 4, 7, 9, 10},
		[]int{1, 2, 8, 10},
	}

	targets := []int{
		9,
		19,
		100,
	}

	expectedResults := [][]int{
		[]int{0, 1},
		[]int{4, 5},
		[]int{0, 0},
	}

	for i := range tests {
		result := twoSumII(tests[i], targets[i])
		if reflect.DeepEqual(expectedResults[i], result) == true {
			t.Logf("\t%s\tCase: %v\tTarget: %v\tResult: %v", succeed, tests[i], targets[i], result)
		} else {
			t.Errorf("\t%s\tCase: %v\tTarget: %v\tResult: %v", failed, tests[i], targets[i], result)
		}
	}
}
