package rpn

import (
	"regexp"
	"unicode"
)

var unary int
var start bool

// This regex matches the number and identifier tokens, respectively.
var numberPattern = regexp.MustCompile(`^(\d+(\.\d*)?|\.\d+?)([eE][-+]?\d+)?`)

// Tokenize breaks an expression into tokens. WELCOME TO THOSE COMPILERS WITH THEIR LEXICAL ANALYSIS!
func Tokenize(exp string) ([]string, error) {
	skip := 0
	start = true
	unary = 0 // MUST!
	var tokens []string
	for i, r := range exp {
		// Previously checked runes
		if skip > 0 {
			skip--
			continue
		}

		// Space
		if unicode.IsSpace(r) {
			continue
		}

		// If first symbol ex: -1-1
		if start && string(r) == "-" {
			unary++
		}
		start = false

		// Operators
		if _, found := Operators[string(r)]; found {
			if string(r) != ")" {
				unary++
			}
			tokens = append(tokens, string(r))
			continue
		}

		// Number
		if  r == '.' || r >= '0' && r <= '9' {
			m := numberPattern.FindString(exp[i:])
			if unary == 2 && tokens[len(tokens)-1] == "-" {
				tokens = tokens[:len(tokens)-1]
				tokens = append(tokens, "-" + m)
			} else {
				tokens = append(tokens, m)
			}
			unary = 0
			skip = len([]rune(m)) - 1
			continue
		}

		// Parenthesis
		switch r {
		case '(':
			tokens = append(tokens, "(" )
			continue
		case ')':
			tokens = append(tokens, ")")
			continue
		default:
			return tokens, makeError(exp, i, "error regx '%c'", r)
		}
	}
	return tokens, nil
}