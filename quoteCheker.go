package main

import "fmt"

func Check(s string) error {

	var dq rune = '"'
	var sq rune = '\''
	var startsWithSigleQuote bool = false
	var startsWithDoubleQuote bool = false

	if s[0] == byte(dq) {
		startsWithDoubleQuote = true
	}

	if s[0] == byte(sq) {
		startsWithSigleQuote = true
	}

	if (s[0] == byte(dq) && s[len(s)-1] == byte(dq)) || (s[0] == byte(sq) && s[len(s)-1] == byte(sq)) {
		s = s[1 : len(s)-1]
	}

	for i := 0; i < len(s); i++ {
		c := rune(s[i])

		switch true {
		case c == dq && startsWithDoubleQuote:

			if i > 0 && s[i-1] != '\\' {
				return fmt.Errorf("syntax error, invalid character [\"] on index %v", i)
			}

		case c == sq && startsWithSigleQuote:
			if i > 0 && s[i-1] != '\\' {
				return fmt.Errorf("syntax error, invalid character ['] on index %v", i)
			}
		}
	}

	return nil
}
