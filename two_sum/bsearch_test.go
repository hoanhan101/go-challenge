package two_sum

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := [][]int{
		[]int{1, 2, 3, 4, 5, 6},
		[]int{36, 48, 50, 62, 80},
		[]int{1, 2, 8, 10},
	}

	targets := []int{
		4,
		50,
		100,
	}

	expectedResults := []bool{
		true,
		true,
		false,
	}

	for i := range tests {
		result := binarySearch(tests[i], targets[i])
		if expectedResults[i] == result {
			t.Logf("\t%s\tCase: %v\tTarget: %v\tResult: %v", succeed, tests[i], targets[i], result)
		} else {
			t.Errorf("\t%s\tCase: %v\tTarget: %v\tResult: %v", failed, tests[i], targets[i], result)
		}
	}
}
