package simple_checker

import (
	"regexp"
	"strings"
	"wordleWizard/config"
	"wordleWizard/ds_utils"
	"wordleWizard/query"
	word_checker "wordleWizard/word-checker"
)

var _ word_checker.Checker = (*simpleChecker)(nil)

func NewSimpleChecker(cfg config.Config, q query.Query) word_checker.Checker {
	reStr := createRegexString(q, cfg.LetterCount)
	re := regexp.MustCompile(reStr)
	return simpleChecker{
		q:  q,
		re: re,
	}
}

type simpleChecker struct {
	re *regexp.Regexp
	q  query.Query
}

func (s simpleChecker) Valid(w string) bool {
	// If Go had a PCRE regex engine everything could be checked in one go
	// with something like: "^(?=[a-z]*i)(?![a-z]*[outyas])(pr[a-z]{2}e[a-z]{1}cl)$" but alas :( ....
	match := s.re.MatchString(w)
	return match && hasAllChars(s.q.Contains, w)
}

func hasAllChars(inc []byte, s string) bool {
	set := ds_utils.SliceToSet[byte, byte, struct{}](inc,
		func(b byte) byte {
			return b
		},
		func(b byte) struct{} {
			return struct{}{}
		})

	for _, c := range s {
		delete(set, byte(c))
	}
	return len(set) == 0
}

func createRegexString(q query.Query, count int) string {
	placeHolder := "#"
	tmpStr := strings.Repeat(placeHolder, count)
	regex := processPlacings(q.Placed, tmpStr)
	bounds := processCharacterBounds(q.Exclude)

	chClass := convertToCharacterClass(bounds)

	return strings.Replace(regex, placeHolder, chClass, -1)
}

func processPlacings(placings map[byte][]int, regex string) string {
	for ch, ints := range placings {
		if ch < 'a' || ch > 'z' {
			continue
		}

		for _, i := range ints {
			if i < 0 || i > 4 {
				continue
			}

			regex = regex[:i] + string(ch) + regex[i+1:]
		}
	}
	return regex
}

func processCharacterBounds(excl []byte) []bound {
	if len(excl) <= 0 {
		return []bound{{
			start: 'a',
			end:   'z',
		}}
	}

	ds_utils.SortSlice(excl)

	arr := make([]bool, 26)

	for _, b := range excl {
		arr[b-'a'] = true
	}

	start := 0

	bounds := make([]bound, 0, 2)
	for i := 0; i < len(arr); {
		if start == 0 && !arr[i] {
			start = i + 'a'

		} else if start != 0 && arr[i] {
			bounds = append(bounds, bound{
				start: byte(start),
				end:   byte(i + 'a' - 1),
			})
			start = 0
		}

		i++
	}

	if start != 0 {
		bounds = append(bounds, bound{
			start: byte(start),
			end:   'z',
		})
	}
	return bounds
}

func convertToCharacterClass(bs []bound) string {
	if len(bs) == 0 {
		return ""
	}

	sb := strings.Builder{}
	sb.WriteByte('[')
	for _, b := range bs {
		if b.start == b.end {
			sb.WriteByte(b.start)
			continue
		}

		sb.WriteByte(b.start)
		sb.WriteByte('-')
		sb.WriteByte(b.end)
	}

	sb.WriteByte(']')
	return sb.String()
}
