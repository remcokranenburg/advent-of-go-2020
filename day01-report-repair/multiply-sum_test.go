package main

import (
    "testing"
)

func isConsecutive(list []int) bool {
    if len(list) < 2 {
        return true
    }

    for i := range list {
        if i > 0 && list[i] != list[i - 1] + 1 {
            return false
        }
    }

    return true
}

func TestBuildRange(t *testing.T) {
    testCases := []int{-1, 0, 1, 2, 3}

    for testCase := range testCases {
        output := buildRange(testCase)

        if testCase >= 0 && len(output) != testCase {
            t.Fatalf("number of indices must be equal to range limit")
        }

        if testCase < 0 && len(output) != 0 {
            t.Fatalf("negative ranges must return no indices")
        }

        if !isConsecutive(output) {
            t.Fatalf("range elements must be consecutive")
        }

        if len(output) > 0 && output[0] != 0 {
            t.Fatalf("range must start with 0")
        }
    }
}
