package simple_checker

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBounds(t *testing.T) {
	testCases := []struct {
		name           string
		exclusions     []byte
		expectedBounds []bound
	}{
		{
			name:       "No exclusions",
			exclusions: nil,
			expectedBounds: []bound{
				{
					start: 'a',
					end:   'z',
				},
			},
		},
		{
			name:       "Exclude a",
			exclusions: []byte{'a'},
			expectedBounds: []bound{
				{
					start: 'b',
					end:   'z',
				},
			},
		},
		{
			name:       "Exclude z",
			exclusions: []byte{'z'},
			expectedBounds: []bound{
				{
					start: 'a',
					end:   'y',
				},
			},
		},
		{
			name:       "Exclude boundaries",
			exclusions: []byte{'a', 'z'},
			expectedBounds: []bound{
				{
					start: 'b',
					end:   'y',
				},
			},
		},
		{
			name:       "Exclude middle consecutive chars",
			exclusions: []byte{'m', 'n', 'o'},
			expectedBounds: []bound{
				{'a', 'l'},
				{'p', 'z'},
			},
		},
		{
			name:       "Exclude starting consecutive letters",
			exclusions: []byte{'a', 'b', 'c'},
			expectedBounds: []bound{
				{'d', 'z'},
			},
		},
		{
			name:       "Exclude ending consecutive letters",
			exclusions: []byte{'x', 'y', 'z'},
			expectedBounds: []bound{
				{'a', 'w'},
			},
		},
		{
			name:       "Exclude mixed",
			exclusions: []byte{'f', 'p', 'v', 'x'},
			expectedBounds: []bound{
				{'a', 'e'},
				{'g', 'o'},
				{'q', 'u'},
				{'w', 'w'},
				{'y', 'z'},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := processCharacterBounds(testCase.exclusions)

			assert.ElementsMatch(t, testCase.expectedBounds, actual)
		})
	}
}
