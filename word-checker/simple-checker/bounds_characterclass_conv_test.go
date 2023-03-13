package simple_checker

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertToCharacterClass(t *testing.T) {
	testCases := []struct {
		name                   string
		bounds                 []bound
		expectedCharacterClass string
	}{
		{
			name: "a to z",
			bounds: []bound{
				{
					start: 'a',
					end:   'z',
				},
			},
			expectedCharacterClass: "[a-z]",
		}, {

			name: "Exclude mixed",
			bounds: []bound{
				{'a', 'e'},
				{'g', 'o'},
				{'q', 'u'},
				{'w', 'w'},
				{'y', 'z'},
			},
			expectedCharacterClass: "[a-eg-oq-uwy-z]",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := convertToCharacterClass(testCase.bounds)
			assert.Equal(t, testCase.expectedCharacterClass, actual)
		})
	}
}
