package query

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateQueryExclusions(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		expectedOutput []byte
	}{
		{
			name:           "Ascii chars",
			input:          "barov3891s",
			expectedOutput: []byte{'b', 'a', 'r', 'o', 'v', '3', '8', '9', '1', 's'},
		}, {
			name:           "Unicode chars",
			input:          "↋ξↈ⟱⟴∰⫷",
			expectedOutput: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			q, err := CreateQuery(testCase.input, "", "")
			assert.Nil(t, err)
			assert.ElementsMatch(t, testCase.expectedOutput, q.Exclude)
		})
	}
}

func TestCreateQueryInclusions(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		expectedOutput []byte
	}{
		{
			name:           "Ascii chars",
			input:          "barov3891s",
			expectedOutput: []byte{'b', 'a', 'r', 'o', 'v', '3', '8', '9', '1', 's'},
		}, {
			name:           "Unicode chars",
			input:          "↋ξↈ⟱⟴∰⫷",
			expectedOutput: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			q, err := CreateQuery("", testCase.input, "")
			assert.Nil(t, err)
			assert.ElementsMatch(t, testCase.expectedOutput, q.Contains)
		})
	}
}

func TestCreateQueryPlacements(t *testing.T) {
	testCases := []struct {
		name              string
		input             string
		expectedError     error
		expectedOutputMap map[byte][]int
	}{
		{
			name:  "Ascii chars valid sequence",
			input: "l34a1q7",
			expectedOutputMap: map[byte][]int{
				'l': {3, 4},
				'a': {1},
				'q': {7},
			},
		},
		{
			name:              "Unicode chars valid sequence",
			input:             "l34⟴1∰7",
			expectedError:     ErrInvalidChar,
			expectedOutputMap: nil,
		},
		{
			name:              "Ascii chars valid sequence but negative index",
			input:             "l-34a1q7",
			expectedOutputMap: nil,
			expectedError:     ErrNoPosition,
		},
		{
			name:              "Ascii chars invalid sequence",
			input:             "abc1f2n3",
			expectedError:     ErrNoPosition,
			expectedOutputMap: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			q, err := CreateQuery("", "", testCase.input)
			assert.ErrorIs(t, err, testCase.expectedError)
			assert.Equal(t, testCase.expectedOutputMap, q.Placed)
		})
	}

}
