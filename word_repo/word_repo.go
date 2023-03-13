package word_repo

import (
	word_checker "wordleWizard/word-checker"
)

type WordRepo interface {
	FindWord(c word_checker.Checker) ([]string, error)
}
