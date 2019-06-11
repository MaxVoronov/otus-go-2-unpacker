package unpacker

import "testing"

func TestIsDigit(t *testing.T) {
	validDigitRunes := []rune("0123456789")
	for _, digitRune := range validDigitRunes {
		if !isDigit(digitRune) {
			t.Errorf("Symbol \"%s\" [rune %d] is not digit", string(digitRune), digitRune)
		}
	}

	invalidDigitRunes := []rune("abcXYZ!@#")
	for _, digitRune := range invalidDigitRunes {
		if isDigit(digitRune) {
			t.Errorf("Symbol \"%s\" [rune %d] is valid digit", string(digitRune), digitRune)
		}
	}
}

func TestIsEscapeSymbol(t *testing.T) {
	validEscapeRunes := []rune(`\`)
	for _, escapeRune := range validEscapeRunes {
		if !isEscapeSymbol(escapeRune) {
			t.Errorf("Symbol \"%s\" [rune %d] is not escape symbol", string(escapeRune), escapeRune)
		}
	}

	invalidEscapeRunes := []rune(`|/#@!`)
	for _, escapeRune := range invalidEscapeRunes {
		if isEscapeSymbol(escapeRune) {
			t.Errorf("Symbol \"%s\" [rune %d] is valid escape symbol", string(escapeRune), escapeRune)
		}
	}
}

func TestCaseRegular(t *testing.T) {
	data := "a4bc2d5e"
	expect := "aaaabccddddde"

	if result, _ := Unpack(data); result != expect {
		t.Errorf("Invalid result: expect \"%s\", got \"%s\"", expect, result)
	}
}

func TestCaseSimple(t *testing.T) {
	data := "abcd"
	expect := "abcd"

	if result, _ := Unpack(data); result != expect {
		t.Errorf("Invalid result: expect \"%s\", got \"%s\"", expect, result)
	}
}

func TestCaseIncorrectOnlyNumbers(t *testing.T) {
	data := "45"

	if result, err := Unpack(data); err == nil {
		t.Errorf("Passed invalid data with result \"%s\", expect error", result)
	}
}

func TestCaseEmptyString(t *testing.T) {
	if result, _ := Unpack(""); result != "" {
		t.Errorf("Invalid result: expect empty string, got \"%s\"", result)
	}
}

func TestCaseIncorrectMissedLetter(t *testing.T) {
	data := "a2b37"

	if result, err := Unpack(data); err == nil {
		t.Errorf("Passed invalid data with result \"%s\", expect error", result)
	}
}

func TestCaseEscapedNumbers(t *testing.T) {
	data := `qwe\4\5`
	expect := "qwe45"

	if result, _ := Unpack(data); result != expect {
		t.Errorf("Invalid result: expect \"%s\", got \"%s\"", expect, result)
	}
}

func TestCasePackedNumber(t *testing.T) {
	data := `qwe\45`
	expect := "qwe44444"

	if result, _ := Unpack(data); result != expect {
		t.Errorf("Invalid result: expect \"%s\", got \"%s\"", expect, result)
	}
}

func TestCasePackedSlash(t *testing.T) {
	data := `qwe\\5`
	expect := `qwe\\\\\`

	if result, _ := Unpack(data); result != expect {
		t.Errorf("Invalid result: expect \"%s\", got \"%s\"", expect, result)
	}
}

func TestCaseComplexSymbols(t *testing.T) {
	data := "€3¢2𐍈4"
	expect := "€€€¢¢𐍈𐍈𐍈𐍈"

	if result, _ := Unpack(data); result != expect {
		t.Errorf("Invalid result: expect \"%s\", got \"%s\"", expect, result)
	}
}

func TestCaseEmoji(t *testing.T) {
	data := "😀2😍3🦇"
	expect := "😀😀😍😍😍🦇"

	if result, _ := Unpack(data); result != expect {
		t.Errorf("Invalid result: expect \"%s\", got \"%s\"", expect, result)
	}
}
