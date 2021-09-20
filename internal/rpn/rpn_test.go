package rpn

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	exp1 = "89*34"
	exp2 = "238+95"
	exp3 = "(36 + 1)/ 10"
	exp4 = "60-50*2^2"
	exp5 = "((3*34+5)*1)/1"
	exp6 = "(-1-1*-1)-1"
	exp7 = "-9-8-7"
	exp8 = "(-1)*-1"
	exp9 = "-1+5*-10"
	exp10 = "(-1.8+5.9)-1.2"
	exp11 = "(4 * 21) - (4 / 2)"
	exp12 = "(-1*2)^2"
	exp13 = "(-1*2-60-345)*0"
	exp14 = "(6*2)/0" // (+Inf)
	exp15 = "(-1.111111*2-600000.435873987-345.849303)*0.9485833"
	expErr1 = "(-e5)"
	expErr2 = "[4-6]*5"
	expErr3 = "       "
	expErr4 = "-----1++555" // defer will catch the error
	expErr5 = "(1 + )" // defer will catch the error
)

func TestTokenizeStandartOperations(t *testing.T) {
	tokens1, _ := Tokenize(exp1)
	tokens2, _ := Tokenize(exp2)
	tokens3, _ := Tokenize(exp3)
	tokens4, _ := Tokenize(exp4)
	tokens5, _ := Tokenize(exp5)

	res1 := []string{"89", "*", "34"}
	res2 := []string{"238", "+", "95"}
	res3 := []string{"(", "36", "+", "1", ")", "/", "10"}
	res4 := []string{"60", "-", "50", "*", "2", "^", "2"}
	res5 := []string{"(", "(", "3", "*", "34", "+", "5", ")", "*", "1", ")", "/", "1"}

	assert.Equal(t, res1, tokens1)
	assert.Equal(t, res2, tokens2)
	assert.Equal(t, res3, tokens3)
	assert.Equal(t, res4, tokens4)
	assert.Equal(t, res5, tokens5)
}

func TestTokenizeUnaryOperations(t *testing.T) {
	tokens6, _ := Tokenize(exp6)
	tokens7, _ := Tokenize(exp7)
	tokens8, _ := Tokenize(exp8)
	tokens9, _ := Tokenize(exp9)
	tokens10, _ := Tokenize(exp10)

	res6 := []string{"(", "-1", "-", "1", "*", "-1", ")", "-", "1"}
	res7 := []string{"-9", "-", "8", "-", "7"}
	res8 := []string{"(", "-1", ")", "*", "-1"}
	res9 := []string{"-1", "+", "5", "*", "-10"}
	res10 := []string{"(", "-1.8", "+", "5.9", ")", "-", "1.2"}

	assert.Equal(t, res6, tokens6)
	assert.Equal(t, res7, tokens7)
	assert.Equal(t, res8, tokens8)
	assert.Equal(t, res9, tokens9)
	assert.Equal(t, res10, tokens10)
}

func TestTokenizeErrors(t *testing.T) {
	tokens1, err1 := Tokenize(expErr1)
	tokens2, err2 := Tokenize(expErr2)
	tokens3, _ := Tokenize(expErr3)

	st1 := calcError {
		err: "error regx 'e'",
		expression: "(-e5)",
		location: 2,
	}

	st2 := calcError {
		err: "error regx '['",
		expression: "[4-6]*5",
		location: 0,
	}

	res1 := []string{"(", "-"}
	res2 := []string(nil)

	assert.Equal(t, res1, tokens1)
	assert.Equal(t, err1, st1)
	assert.Equal(t, res2, tokens2)
	assert.Equal(t, res2, tokens3)
	assert.Equal(t, err2, st2)
}

// The ConvertToRPN function returns no errors, only CalculateRPN returns!
func TestConvertToRPN(t *testing.T) {
	c1 := RPN{}
	tokens1, _ := Tokenize(exp1)
	c1.ConvertToRPN(tokens1)
	res1 := "89 34 * "

	c2 := RPN{}
	tokens2, _ := Tokenize(exp10)
	c2.ConvertToRPN(tokens2)
	res2 := "-1.8 5.9 + 1.2 - "

	c3 := RPN{}
	tokens3, _ := Tokenize(exp11)
	c3.ConvertToRPN(tokens3)
	res3 := "4 21 * 4 2 / - "

	c4 := RPN{}
	tokens4, _ := Tokenize(expErr4)
	c4.ConvertToRPN(tokens4)
	res4 := "- - - - 1 - + 555 + " // as it should

	assert.Equal(t, c1.rpnExpression, res1)
	assert.Equal(t, c2.rpnExpression, res2)
	assert.Equal(t, c3.rpnExpression, res3)
	assert.Equal(t, c4.rpnExpression, res4)
}

func TestCalculateRPN(t *testing.T) {
	c1 := RPN{}
	tokens1, _ := Tokenize(exp1)
	c1.ConvertToRPN(tokens1)
	c1.CalculateRPN()

	c2 := RPN{}
	tokens2, _ := Tokenize(exp10)
	c2.ConvertToRPN(tokens2)
	c2.CalculateRPN()

	c3 := RPN{}
	tokens3, _ := Tokenize(exp11)
	c3.ConvertToRPN(tokens3)
	c3.CalculateRPN()

	c4 := RPN{}
	tokens4, _ := Tokenize(expErr4)
	c4.ConvertToRPN(tokens4)
	c4.CalculateRPN()

	c5 := RPN{}
	tokens5, _ := Tokenize(exp4)
	c5.ConvertToRPN(tokens5)
	c5.CalculateRPN()

	c6 := RPN{}
	tokens6, _ := Tokenize(exp6)
	c6.ConvertToRPN(tokens6)
	c6.CalculateRPN()

	c7 := RPN{}
	tokens7, _ := Tokenize(exp12)
	c7.ConvertToRPN(tokens7)
	c7.CalculateRPN()

	c8 := RPN{}
	tokens8, _ := Tokenize(exp13)
	c8.ConvertToRPN(tokens8)
	c8.CalculateRPN()

	c9 := RPN{}
	tokens9, _ := Tokenize(expErr5)
	c9.ConvertToRPN(tokens9)
	c9.CalculateRPN()

	c10 := RPN{}
	tokens10, _ := Tokenize(exp14)
	c10.ConvertToRPN(tokens10)
	c10.CalculateRPN()

	c11 := RPN{}
	tokens11, _ := Tokenize(exp15)
	c11.ConvertToRPN(tokens11)
	c11.CalculateRPN()

	assert.Equal(t, c1.GetResult(), 3026.0)
	assert.Equal(t, c2.GetResult(), 2.9)
	assert.Equal(t, c3.GetResult(), 82.0)
	assert.Equal(t, c4.GetResult(), 0.0)
	assert.Equal(t, c5.GetResult(), -140.0)
	assert.Equal(t, c6.GetResult(), -1.0)
	assert.Equal(t, c7.GetResult(), 4.0)
	assert.Equal(t, c8.GetResult(), 0.0)
	assert.Equal(t, c9.GetResult(), 0.0) // defer will catch the error and return 0
	assert.Equal(t, fmt.Sprint(c10.GetResult()), fmt.Sprint("+Inf")) // defer will catch the error and return 0 (+Inf Equal 0.0)
	assert.Equal(t, c11.GetResult(), -569480.568299)
}