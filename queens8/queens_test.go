package main

import (
	"testing"
)

type testBoardValid struct {
	board  [][]string
	result bool
	name   string
}

func TestBoardValid(t *testing.T) {
	testCases := []testBoardValid{
		{
			board: [][]string{
				{queen, "", "", "", "", "", "", queen},
				{"", "", "", "", "", "", "", ""},
				{"", "", "", "", "", "", "", ""},
				{"", "", "", "", "", "", "", ""},
				{"", "", "", "", "", "", "", ""},
				{"", "", "", "", "", "", "", ""},
				{"", "", "", "", "", "", "", ""},
				{"", "", "", "", "", "", "", ""},
			},
			result: false,
			name:   "row wrong",
		},
		{
			board: [][]string{
				{queen, "", "", "", "", "", "", ""},
				{"", "", "", "", "", "", "", ""},
				{"", "", "", "", "", "", "", ""},
				{"", "", "", "", "", "", "", ""},
				{"", "", "", "", "", "", "", ""},
				{"", "", "", "", "", "", "", ""},
				{"", "", "", "", "", "", "", ""},
				{queen, "", "", "", "", "", "", ""},
			},
			result: false,
			name:   "col wrong",
		},
		{
			board: [][]string{
				{queen, "", "", "", "", "", "", ""},
				{"", "", "", "", "", "", "", ""},
				{"", "", "", "", "", "", "", ""},
				{"", "", "", queen, "", "", "", ""},
				{"", "", "", "", "", "", "", ""},
				{"", "", "", "", "", "", "", ""},
				{"", "", "", "", "", "", "", ""},
				{"", "", "", "", "", "", "", ""},
			},
			result: false,
			name:   "diag top left wrong",
		},
		{
			board: [][]string{
				{"", "", "", "", "", "", "", ""},
				{"", "", "", "", "", "", "", ""},
				{"", "", "", "", "", "", "", ""},
				{"", "", "", queen, "", "", "", ""},
				{"", "", "", "", "", "", "", ""},
				{"", "", "", "", "", "", "", ""},
				{"", "", "", "", "", "", "", ""},
				{"", "", "", "", "", "", "", queen},
			},
			result: false,
			name:   "diag bottom right wrong",
		},
		{
			board: [][]string{
				{"", "", "", "", "", "", "", ""},
				{"", "", "", "", "", "", "", ""},
				{"", "", "", "", "", "", "", ""},
				{"", "", "", queen, "", "", "", ""},
				{"", "", "", "", "", "", "", ""},
				{"", queen, "", "", "", "", "", ""},
				{"", "", "", "", "", "", "", ""},
				{"", "", "", "", "", "", "", ""},
			},
			result: false,
			name:   "diag bottom left wrong",
		},
		{
			board: [][]string{
				{"", "", "", "", "", "", "", ""},
				{"", "", "", "", "", queen, "", ""},
				{"", "", "", "", "", "", "", ""},
				{"", "", "", queen, "", "", "", ""},
				{"", "", "", "", "", "", "", ""},
				{"", "", "", "", "", "", "", ""},
				{"", "", "", "", "", "", "", ""},
				{"", "", "", "", "", "", "", ""},
			},
			result: false,
			name:   "diag top right wrong",
		},
		{
			board: [][]string{
				{"", "", queen, "", "", "", "", ""},
				{"", "", "", "", "", "", "", queen},
				{"", "", "", queen, "", "", "", ""},
				{"", "", "", "", "", "", queen, ""},
				{queen, "", "", "", "", "", "", ""},
				{"", "", "", "", "", queen, "", ""},
				{"", queen, "", "", "", "", "", ""},
				{"", "", "", "", queen, "", "", ""},
			},
			result: true,
			name:   "right board 1",
		},
		{
			board: [][]string{
				{"", "", "", queen, "", "", "", ""},
				{"", queen, "", "", "", "", "", ""},
				{"", "", "", "", queen, "", "", ""},
				{"", "", "", "", "", "", "", queen},
				{"", "", "", "", "", queen, "", ""},
				{queen, "", "", "", "", "", "", ""},
				{"", "", queen, "", "", "", "", ""},
				{"", "", "", "", "", "", queen, ""},
			},
			result: true,
			name:   "right board 2",
		},
	}

	for _, tc := range testCases {
		result := boardValid(tc.board)
		if result != tc.result {
			t.Error(
				"testcase", tc.name,
				"expected", tc.result,
				"got", result,
			)
		}
	}
}

func TestPlaceQueens(t *testing.T) {
	board := placeQueens()
	if !boardValid(board) {
		t.Error("place queens failed")
	}
}