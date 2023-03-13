package memory_repo

import (
	"fmt"
	"io/ioutil"
	"strings"
	word_checker "wordleWizard/word-checker"
	"wordleWizard/word_repo"
)

var _ word_repo.WordRepo = (*memoryRepo)(nil)

type memoryRepo struct {
	words []string
}

func NewMemoryRepo() (word_repo.WordRepo, error) {
	m := memoryRepo{}
	err := m.init()
	return &m, err
}

func (m *memoryRepo) FindWord(c word_checker.Checker) ([]string, error) {
	validWords := make([]string, 0)
	for _, w := range m.words {
		if c.Valid(w) {
			validWords = append(validWords, w)
		}
	}

	return validWords, nil
}

func (m *memoryRepo) init() error {
	bs, err := ioutil.ReadFile("data/wordle-la")
	if err != nil {
		return err
	}

	s := string(bs)
	m.words = strings.Split(s, "\n")

	return nil
}

func (m *memoryRepo) print() error {
	for _, word := range m.words {
		fmt.Println(word)
	}
	return nil
}
