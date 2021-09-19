package tokenize

// Operator precedence for Reverse Polish Notation.
var operators = map[token]int {
	operator('*'): 2,
	operator('/'): 2,
	operator('+'): 1,
	operator('-'): 1,
}