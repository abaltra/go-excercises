package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	candidates := []int{2, 3, 6, 7}
	target := 7

	expectedMap := map[string]int{
		"223": 0,
		"7":   0,
	}

	toMap := func(a []int) string {
		sb := strings.Builder{}
		for _, n := range a {
			sb.WriteString(fmt.Sprint(n))
		}

		return sb.String()
	}

	results := combinationSum(candidates, target)

	for _, result := range results {
		m := toMap(result)
		if _, ok := expectedMap[m]; !ok {
			t.Errorf("Result %s not found in expected set", m)
		} else {
			expectedMap[m] = 1
		}
	}

	for key, val := range expectedMap {
		if val == 0 {
			t.Errorf("Set %s not found in results", key)
		}
	}

}
