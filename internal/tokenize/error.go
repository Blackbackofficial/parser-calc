package tokenize

import "fmt"

type calcError struct {
	err string
	expression string
	location int
}

// Nice output like in Python
func (e calcError) Error() string {
	return fmt.Sprintf("calc: %s\n       %s\n%*s", e.err, e.expression, e.location+8, "^") // 8 отступов
}

// Print error
func makeError(expression string, loc int, desc string, args ...interface{}) calcError {
	return calcError{
		err: fmt.Sprintf(desc, args...),
		expression: expression,
		location: loc,
	}
}