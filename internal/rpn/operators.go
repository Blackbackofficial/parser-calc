package rpn

// Operators precedence for Reverse Polish Notation.
var Operators = map[string]int {
	"^": 3,
	"*": 2,
	"/": 2,
	"+": 1,
	"-": 1,
	"(": 1,
	")": 1,
}