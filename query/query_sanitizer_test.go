package query

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"wordleWizard/config"
)

func TestSanitizeExclusions(t *testing.T) {
	testCases := []struct {
		name               string
		inputExclusions    []byte
		expectedExclusions []byte
	}{
		{
			name:               "Only alpha chars included",
			inputExclusions:    []byte{'1', 'a', '2', 'b', '!', 'c', '#', 'd'},
			expectedExclusions: []byte{'a', 'b', 'c', 'd'},
		},
		{
			name:               "Upper case converted to lower",
			inputExclusions:    []byte{'A', 'B', 'C', 'D'},
			expectedExclusions: []byte{'a', 'b', 'c', 'd'},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			q := Query{Exclude: testCase.inputExclusions}
			err := SanitizeQuery(config.Config{LetterCount: 5}, &q)

			assert.Nil(t, err)
			assert.ElementsMatch(t, testCase.expectedExclusions, q.Exclude)
		})
	}
}

func TestSanitizeInclusions(t *testing.T) {
	testCases := []struct {
		name               string
		inputInclusions    []byte
		expectedInclusions []byte
		expectedErr        error
		cfg                config.Config
	}{
		{
			name:               "Only alpha chars included",
			inputInclusions:    []byte{'1', 'a', '2', 'b', '!', 'c', '#', 'd'},
			expectedInclusions: []byte{'a', 'b', 'c', 'd'},
			cfg:                config.Config{LetterCount: 5},
		},

		{
			name:               "Cannot include more chars than total available places",
			inputInclusions:    []byte{'a', 'b', 'c', 'd', 'e', 'f'},
			expectedInclusions: []byte{'a', 'b', 'c', 'd', 'e', 'f'},
			expectedErr:        ErrInclusionCount,
			cfg:                config.Config{LetterCount: 5},
		},
		{

			name:               "Upper case converted to lower",
			inputInclusions:    []byte{'A', 'B', 'C', 'D'},
			expectedInclusions: []byte{'a', 'b', 'c', 'd'},
			cfg:                config.Config{LetterCount: 5},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			q := Query{Contains: testCase.inputInclusions}
			err := SanitizeQuery(config.Config{LetterCount: 5}, &q)

			assert.ErrorIs(t, err, testCase.expectedErr)
			assert.ElementsMatch(t, testCase.expectedInclusions, q.Contains)
		})
	}
}

func TestSanitizePlaced(t *testing.T) {
	testCases := []struct {
		name        string
		cfg         config.Config
		inputMap    map[byte][]int
		expectedMap map[byte][]int
	}{
		{
			name: "Invalid positions excluded",
			cfg:  config.Config{LetterCount: 5},
			inputMap: map[byte][]int{
				'a': {1, 2, 7, 8},
			},
			expectedMap: map[byte][]int{
				'a': {1, 2},
			},
		},
		{
			name: "Invalid chars excluded",
			cfg:  config.Config{LetterCount: 5},
			inputMap: map[byte][]int{
				'!': {1},
				'5': {2},
				'a': {3},
			},
			expectedMap: map[byte][]int{
				'a': {3},
			},
		},
		{
			name: "Uppercase changed to lowercase",
			cfg:  config.Config{LetterCount: 5},
			inputMap: map[byte][]int{
				'B': {0},
				'D': {2},
				'W': {3},
			},
			expectedMap: map[byte][]int{
				'b': {0},
				'd': {2},
				'w': {3},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			q := Query{Placed: testCase.inputMap}
			err := SanitizeQuery(testCase.cfg, &q)

			assert.Nil(t, err)
			assert.Equal(t, testCase.expectedMap, q.Placed)
		})
	}
}
