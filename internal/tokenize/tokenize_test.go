package tokenize

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	expr1 = "(10 + 5) + 10"
	expr2 = "1 + 1 * 5"
	expr3 = "1 - 8 / 4"
	expr4 = "(10 + 2)/6 + 1"
	expr5 = "(2 * 15) - (6 / 3)"
	expr6 = "-12 + 1 * 22 * 2 - 5 / 5 - 53 + 12 * 2 - 2 * -6"
	exp1 = "89*34"
	exp2 = "238+95"
	exp3 = "60-50"
	exp4 = "89*34"

	res1 = 25
	res2 = 6
	res3 = -1
	res4 = 3
	res5 = 28
	res6 = 14

	inCorrect1 = "(+)"
	inCorrect2 = "(1 + )"
	inCorrect3 = "*3-2"
	inCorrect4 = "/"
	inCorrect5 = "[55 + 4]"
	inCorrect6 = "(5^2 + 4)"
	inCorrect7 = "Фыва"
	inCorrect8 = "1 * 1)"
)

//func TestCheckSuccess(t *testing.T) {
//	var sortedInput = []string{"1", "2", "3", "4", "5"}
//
//	tokens, err := Tokenize(str)
//	if err != nil {
//		t.Fatalf("Check failed on sorted input: %s", err)
//	}
//}
func TestTokenize(t *testing.T) {

	var a string = "Hello"
	var b string = "Hello"

	assert.Equal(t, a, b, "The two words should be the same.")

}