package main

import (
	"reflect"
	"testing"
)

type testMerge struct {
	a      []int
	b      []int
	result []int
}

func TestMerge(t *testing.T) {
	testCases := []testMerge{
		{
			a:      []int{1, 2, 5},
			b:      []int{1, 1, 3},
			result: []int{1, 1, 1, 2, 3, 5},
		},
		{
			a:      []int{2, 5},
			b:      []int{3},
			result: []int{2, 3, 5},
		},
	}

	for _, tc := range testCases {
		result := merge(tc.a, tc.b)
		if !reflect.DeepEqual(result, tc.result) {
			t.Error(
				"a", tc.a,
				"b", tc.b,
				"expected", tc.result,
				"got", result,
			)
		}
	}
}

type testMergeSort struct {
	a      []int
	result []int
}

func TestMergeSort(t *testing.T) {
	testCases := []testMergeSort{
		{
			a:      []int{1},
			result: []int{1},
		},
		{
			a:      []int{10, 2, 5},
			result: []int{2, 5, 10},
		},
		{
			a:      []int{1, 10, 2, 5},
			result: []int{1, 2, 5, 10},
		},
	}

	for _, tc := range testCases {
		result := mergeSort(tc.a)
		if !reflect.DeepEqual(result, tc.result) {
			t.Error(
				"a", tc.a,
				"expected", tc.result,
				"got", result,
			)
		}
	}
}
