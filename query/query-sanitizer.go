package query

import (
	"errors"
	"wordleWizard/config"
)

var ErrInclusionCount = errors.New("inclusions exceed word length")

func SanitizeQuery(cfg config.Config, q *Query) error {
	q.Exclude = sanitizeBytes(q.Exclude)
	q.Include = sanitizeBytes(q.Include)
	q.Placed = sanitizeMap(cfg, q.Placed)

	err := checkInclusionsLength(cfg, q.Include)
	if err != nil {
		return err
	}

	return nil
}

func checkInclusionsLength(cfg config.Config, b []byte) error {
	if len(b) < cfg.LetterCount {
		return nil
	}

	return ErrInclusionCount
}

func sanitizeMap(cfg config.Config, m map[byte][]int) map[byte][]int {
	sM := map[byte][]int{}

	for b, ints := range m {
		bs := sanitizeBytes([]byte{b})
		if len(bs) < 1 {
			continue
		}

		s := sanitizePositions(cfg, ints)
		if len(s) > 0 {
			sM[bs[0]] = s
		}

	}
	return sM
}

func sanitizePositions(cfg config.Config, bs []int) []int {
	valid := make([]int, 0, len(bs))
	for _, b := range bs {
		if b < 0 || b > (cfg.LetterCount-1) {
			continue
		}
		valid = append(valid, b)
	}
	return valid
}

func sanitizeBytes(bs []byte) []byte {
	set := map[byte]struct{}{}
	for _, b := range bs {
		if b >= 'a' && b <= 'z' {
			set[b] = struct{}{}
		} else if b >= 'A' && b <= 'Z' {
			set[b+('a'-'A')] = struct{}{}
		}
	}

	validB := make([]byte, 0, len(set))
	for b := range set {
		validB = append(validB, b)
	}
	return validB
}
