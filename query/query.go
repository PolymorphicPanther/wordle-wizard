package query

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidChar = errors.New("unsupported character")
var ErrUnexpectedChar = errors.New("unexpected character")
var ErrNoPosition = errors.New("expected position")

type Query struct {
	Include []byte
	Exclude []byte
	Placed  map[byte][]int
}

func (q *Query) String() string {
	sb := strings.Builder{}
	if len(q.Exclude) > 0 {
		sb.Grow(len(q.Exclude)*2 + 14)
		sb.WriteString("Exclusions: ")
		byteSliceToString(&sb, q.Exclude)
		sb.WriteString("\n")
	}

	if len(q.Exclude) > 0 {
		sb.Grow(len(q.Include)*2 + 14)
		sb.WriteString("Inclusions: ")
		byteSliceToString(&sb, q.Include)
		sb.WriteString("\n")
	}

	if len(q.Placed) > 0 {
		sb.Grow(len(q.Placed)*3 + 14)
		sb.WriteString("Placed: ")

		for b, ints := range q.Placed {
			sb.WriteString(fmt.Sprintf("(%q | ", b))

			for idx, i := range ints {
				sb.WriteString(strconv.Itoa(i))
				if idx != len(ints)-1 {
					sb.WriteByte(',')
				}
			}
			sb.WriteString(") ")
		}
	}

	return sb.String()
}

func byteSliceToString(sb *strings.Builder, bs []byte) {
	for i, b := range bs {
		sb.WriteByte(b)
		if i != len(bs)-1 {
			sb.WriteByte(',')
		}
	}
}

func parseExclusions(e string) []byte {
	return parseAsciiChars(e)
}

func parseInclusions(i string) []byte {
	return parseAsciiChars(i)
}

func parseAsciiChars(s string) []byte {
	b := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			continue
		}
		b = append(b, s[i])
	}
	return b
}

func parsePlacedCharacters(placd string) (map[byte][]int, error) {
	pos := map[byte][]int{}
	var currentLetter byte

	// TODO: this format will not work for idx > 9 :eyes: luckily Wordle only has 5 chars
	for i := 0; i < len(placd); {
		b := placd[i]
		if b > unicode.MaxASCII {
			return nil, ErrInvalidChar
		}

		if !(b >= 'a' && b <= 'z') && !(b >= 'A' && b <= 'Z') {
			return nil, fmt.Errorf("pos %d: %w", i, ErrUnexpectedChar)
		}

		currentLetter = b
		i++

		idx := make([]int, 0)
		for {
			if i == len(placd) {
				break
			}

			pos := placd[i]
			if pos > '9' || pos < '0' {
				break
			}
			idx = append(idx, int(pos-'0'))
			i++
		}

		if len(idx) < 1 {
			return nil, fmt.Errorf("ch: %s pos: %d %w", string(currentLetter), i-1, ErrNoPosition)
		}

		pos[currentLetter] = idx
	}
	return pos, nil
}

func CreateQuery(excl, incl, placd string) (Query, error) {
	posMap, err := parsePlacedCharacters(placd)
	if err != nil {
		return Query{}, err
	}

	return Query{
		Include: parseInclusions(incl),
		Exclude: parseExclusions(excl),
		Placed:  posMap,
	}, nil
}
