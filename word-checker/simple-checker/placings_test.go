package simple_checker

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
	"wordleWizard/query"
)

func TestOverlappingPlacings(t *testing.T) {
	q := query.Query{
		Placed: map[byte][]int{
			'b': {3},
			'z': {3},
			'q': {3},
		},
	}

	actualString := processPlacings(q.Placed, "#####")
	match, err := regexp.Match(`###([bzq])#`, []byte(actualString))
	assert.Nil(t, err)
	assert.True(t, match)
}

func TestPlacings(t *testing.T) {

	testCases := []struct {
		name     string
		q        query.Query
		expected string
	}{
		{
			name: "One placed",
			q: query.Query{
				Placed: map[byte][]int{
					'a': {1},
				},
			},
			expected: "#a###",
		},
		{
			name: "Same char many places",
			q: query.Query{
				Placed: map[byte][]int{
					'a': {0, 2, 4},
				},
			},
			expected: "a#a#a",
		},
		{
			name: "Different chars different places",
			q: query.Query{
				Placed: map[byte][]int{
					'b': {0},
					'z': {1},
					'q': {4},
				},
			},
			expected: "bz##q",
		},
		{
			name:     "No placings",
			q:        query.Query{},
			expected: "#####",
		},
		{
			name: "Out of range indices",
			q: query.Query{
				Placed: map[byte][]int{
					'b': {-1},
					'z': {5},
					'q': {69},
				},
			},
			expected: "#####",
		},
		{
			name: "Out of range chars",
			q: query.Query{
				Placed: map[byte][]int{
					'A':  {0},
					'!':  {1},
					'\n': {4},
				},
			},
			expected: "#####",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actualString := processPlacings(testCase.q.Placed, "#####")

			assert.Equal(t, testCase.expected, actualString, "unexpected regex generated")
		})
	}
}
