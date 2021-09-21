package rpn

import (
	"github.com/stretchr/testify/assert"
	"math"
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

func TestConvertToRPNAndCalculateRPN(t *testing.T) {
	testCases := []struct {
		name        		string
		tokensCheck		 	bool
		input     			string
		tokens				[]string
		errTokens 			calcError
		rpnExpression 		string	// The ConvertToRPN function returns no errors, only CalculateRPN returns!
		result 				float64
	} {
		{
			name:        	"Valid token (-1.8+5.9)-1.2",
			tokensCheck: 	false,
			input:       	"(-1.8+5.9)-1.2",
			tokens: 		[]string{"(", "-1.8", "+", "5.9", ")", "-", "1.2"},
		}, {
			name:        	"Valid token -1+5*-10",
			tokensCheck: 	false,
			input:       	"-1+5*-10",
			tokens: 		[]string{"-1", "+", "5", "*", "-10"},
		}, {
			name:        	"Valid token (-1)*-1",
			tokensCheck: 	false,
			input:       	"(-1)*-1",
			tokens: 		[]string{"(", "-1", ")", "*", "-1"},
		}, {
			name:        	"Valid token -9-8-7",
			tokensCheck: 	false,
			input:       	"-9-8-7",
			tokens: 		[]string{"-9", "-", "8", "-", "7"},
		}, {
			name:        	"Valid token (-1-1*-1)-1",
			tokensCheck: 	false,
			input:       	"(-1-1*-1)-1",
			tokens: 		[]string{"(", "-1", "-", "1", "*", "-1", ")", "-", "1"},
		}, {
			name:        	"Invalid space",
			tokensCheck: 	false,
			input:       	"       ",
			tokens: 		[]string(nil),
		}, {
			name:        	"Invalid (-e5)",
			tokensCheck: 	false,
			input:       	"(-e5)",
			tokens: 		[]string{"(", "-"},
			errTokens: 		calcError{err: "error regx 'e'", expression: "(-e5)", location: 2},
		}, {
			name:        	"valid [4-6]*5",
			tokensCheck: 	false,
			input:       	"[4-6]*5",
			tokens: 		[]string(nil),
			errTokens: 		calcError {err: "error regx '['", expression: "[4-6]*5", location: 0},
		}, {
			name:        	"valid 89*34",
			tokensCheck: 	true,
			input:       	"89*34",
			rpnExpression: 	"89 34 * ",
			result: 		3026.0,
		}, {
			name:        	"valid (-1.8+5.9)-1.2",
			tokensCheck: 	true,
			input:       	"(-1.8+5.9)-1.2",
			rpnExpression: 	"-1.8 5.9 + 1.2 - ",
			result: 		2.9,
		}, {
			name:        	"valid 4 * 21) - (4 / 2)",
			tokensCheck: 	true,
			input:       	"(4 * 21) - (4 / 2)",
			rpnExpression: 	"4 21 * 4 2 / - ",
			result: 		82.0,
		}, {
			name:        	"invalid -----1++555",
			tokensCheck: 	true,
			input:       	"-----1++555",
			rpnExpression: 	"- - - - 1 - + 555 + ",
			result: 		0.0,
		}, {
			name:        	"valid 60-50*2^2",
			tokensCheck: 	true,
			input:       	"60-50*2^2",
			rpnExpression: 	"60 50 2 2 ^ * - ",
			result: 		-140.0,
		}, {
			name:       	"valid (-1-1*-1)-1",
			tokensCheck: 	true,
			input:       	"(-1-1*-1)-1",
			rpnExpression: 	"-1 1 -1 * - 1 - ",
			result: 		-1.0,
		}, {
			name:        	"invalid (-1*2)^2",
			tokensCheck: 	true,
			input:       	"(-1*2)^2",
			rpnExpression: 	"-1 2 * 2 ^ ",
			result: 		4.0,
		},
		{
			name:        	"invalid (-1*2-60-345)*0",
			tokensCheck: 	true,
			input:       	"(-1*2-60-345)*0",
			rpnExpression:	"-1 2 * 60 - 345 - 0 * ",
			result: 		0.0,
		},
		{
			name:        	"invalid (1 + )",
			tokensCheck: 	true,
			input:       	"(1 + )",
			rpnExpression: 	"1 + ",
			result: 		0.0,
		},
		{
			name:        	"invalid (6*2)/0",
			tokensCheck: 	true,
			input:       	"(6*2)/0", // (+Inf)
			rpnExpression: 	"6 2 * 0 / ",
			result: 		math.Inf(1),
		},
		{
			name:        	"valid (-1.111111*2-600000.435873987-345.849303)*0.9485833",
			tokensCheck: 	true,
			input:       	"(-1.111111*2-600000.435873987-345.849303)*0.9485833",
			rpnExpression: 	"-1.111111 2 * 600000.435873987 - 345.849303 - 0.9485833 * ",
			result: 		-569480.568299,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.tokensCheck {
				c := RPN{}
				token, _ := Tokenize(tc.input)
				c.ConvertToRPN(token)
				c.CalculateRPN()
				assert.Equal(t, c.rpnExpression, tc.rpnExpression)
				assert.Equal(t, c.GetResult(), tc.result)
			} else {
				tokens, err := Tokenize(tc.input)
				assert.Equal(t, tokens, tc.tokens)
				if err != nil {
					assert.Equal(t, err, tc.errTokens)
				}
			}
		})
	}
}