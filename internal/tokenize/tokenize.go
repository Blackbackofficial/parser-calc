package tokenize

import (
	"regexp"
	"strconv"
	"unicode"
)

// WELCOME TO THOSE COMPILERS WITH THEIR LEXICAL ANALYSIS!
type token interface {}

type number float64

type operator string

// This regex matches the number and identifier tokens, respectively.
var numberPattern = regexp.MustCompile(`^(\d+(\.\d*)?|\.\d+?)([eE][-+]?\d+)?`)

// Tokenize breaks an expression into tokens
func Tokenize(exp string) ([]token, error) {
	skip := 0
	var tokens []token
	for i, ru := range exp {
		// previously checked runes
		if skip > 0 {
			skip--
			continue
		}

		// Space
		if unicode.IsSpace(ru) {
			continue
		}

		// Number
		if  ru == '.' || ru >= '0' && ru <= '9' {
			m := numberPattern.FindString(exp[i:])
			x, err := strconv.ParseFloat(m, 64)
			if err != nil {
				return nil, makeError(exp, i, "error in number (%v)", err)
			}
			tokens = append(tokens, number(x))
			skip = len([]rune(m)) - 1
			continue
		}

		// Operators
		if _, found := operators[operator(ru)]; found {
			if ru == '-' || ru == '+' {
				if len(tokens) == 0 {
					tokens = append(tokens, operator(ru))
					continue
				}
				switch tokens[len(tokens)-1].(type) {
				case operator:
					tokens = append(tokens, operator(ru))
					continue
				}
			}
			tokens = append(tokens, operator(ru))
			continue
		}

		// parenthesis
		switch ru {
		case '(':
			tokens = append(tokens, "(")
			continue
		case ')':
			tokens = append(tokens, ")")
			continue
		default:
			return nil, makeError(exp, i, "error regx '%c'", ru)
		}
	}
	return tokens, nil
}