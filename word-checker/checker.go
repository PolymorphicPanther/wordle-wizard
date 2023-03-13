package word_checker

type Checker interface {
	Valid(w string) bool
}
