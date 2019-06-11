package unpacker

import (
	"errors"
	"strconv"
	"strings"
)

func Unpack(data string) (string, error) {
	var buf string
	var isEscaped bool
	var result strings.Builder

	for _, r := range data {
		s := string(r)

		if isDigit(r) {
			// First char is number
			if buf == "" {
				return "", errors.New("invalid input string")
			}

			// If number is counter then save buffered symbol to result
			if buf != "" && !isEscaped {
				n, _ := strconv.Atoi(s)
				result.WriteString(strings.Repeat(buf, n))
				buf = ""
				continue
			}
		}

		if buf != "" {
			// Escape symbol
			if isEscapeSymbol(r) && !isEscaped {
				isEscaped = true
				continue
			}

			// Save buffered symbol to result
			isEscaped = false
			result.WriteString(buf)
			buf = s
			continue

		}

		// Add symbol to buffer
		buf = s
	}

	// Read and save latest symbol from buffer
	result.WriteString(buf)

	return result.String(), nil
}

func isEscapeSymbol(r rune) bool {
	return r == 92
}

func isDigit(r rune) bool {
	return r >= 48 && r <= 57
}
