package rpn

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

type RPN struct {
	operatorStack	[]string
	rpnStack      	[]string
	rpnExpression 	string
	result       	float64
	resultString 	string
}

// flags for convert to RPN
var popLoop bool
var firstItem bool

// Operators precedence for Reverse Polish Notation.
var Operators = map[string]int {
	"^": 3,
	"*": 2,
	"/": 2,
	"+": 1,
	"-": 1,
	"(": 0,
	")": 0,
}

func (c *RPN) GetResult() float64 {
	return c.result
}

func (c *RPN) AppendRPNItem(item string) {
	if item != "(" && item != ")" {
		c.rpnExpression = c.rpnExpression + item + " "
		c.rpnStack = append(c.rpnStack, item)
	}
}

func (c *RPN) GetLastOperatorFromStack() string {
	if len(c.operatorStack) > 0 {
		return c.operatorStack[len(c.operatorStack)-1]
	}
	return ""
}

func (c *RPN) PopOperatorFromStack() []string {
	if len(c.operatorStack) > 0 {
		c.operatorStack = c.operatorStack[:len(c.operatorStack)-1]
	}
	return c.operatorStack
}

func (c *RPN) ConvertToRPN(expStack []string) {
	for i := range expStack {
		item := expStack[i]
		if _, found := Operators[item]; found {
			if len(c.operatorStack) == 0 || !firstItem {
				firstItem = true
				c.operatorStack = append(c.operatorStack, item)
			} else {
				if item == "(" {
					c.operatorStack = append(c.operatorStack, item)
					continue
				}

				if len(c.operatorStack) > 0 && item == ")" {
					for len(c.operatorStack) > 0 && c.GetLastOperatorFromStack() != "(" {
						c.AppendRPNItem(c.GetLastOperatorFromStack())
						c.PopOperatorFromStack()
					}
					if len(c.operatorStack) > 0 && c.GetLastOperatorFromStack() == "(" {
						c.PopOperatorFromStack()
					}
					continue
				}

				for len(c.operatorStack) > 0 && (Operators[item] <= Operators[c.GetLastOperatorFromStack()]) {
					c.AppendRPNItem(c.GetLastOperatorFromStack())
					c.PopOperatorFromStack()
					popLoop = true
				}

				if popLoop {
					c.operatorStack = append(c.operatorStack, item)
					popLoop = false
				} else if len(c.operatorStack) > 0 && (Operators[item] > Operators[c.GetLastOperatorFromStack()]) {
					c.operatorStack = append(c.operatorStack, item)
				}
			}
		} else {
			c.AppendRPNItem(item)
		}
	}

	for len(c.operatorStack) > 0 {
		c.AppendRPNItem(c.GetLastOperatorFromStack())
		c.PopOperatorFromStack()
	}
}

func (c *RPN) SimpleCalculate(val1 float64, val2 float64, operator string) float64 {
	auxResult := 0.0
	switch operator {
	case "^":
		auxResult = math.Pow(val1, val2)
	case "*":
		auxResult = val1 * val2
	case "/":
		auxResult = val1 / val2
	case "+":
		auxResult = val1 + val2
	case "-":
		auxResult = val1 - val2
	}
	return auxResult
}

func (c *RPN) CalculateRPN() (float64, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Invalid RPN")
		}
	}()
	auxStack := c.rpnStack
	for len(auxStack) > 1 {
		for i := 0; i < len(auxStack); i++ {
			if _, found := Operators[auxStack[len(auxStack) - 1]]; !found {
				return -1, errors.New("Invalid operator")
			}
			item := auxStack[i]
			if _, found := Operators[item]; found {
				value1, err := strconv.ParseFloat(auxStack[i-2], 64)
				if err != nil {
					fmt.Printf("Error value1 as %s", auxStack[i-1])
				}
				value2, err := strconv.ParseFloat(auxStack[i-1], 64)
				if err != nil {
					fmt.Printf("Error value1 as %s", auxStack[i-1])
				}

				resultCalc := c.SimpleCalculate(value1, value2, item)
				auxResult := fmt.Sprintf("%f", resultCalc)

				auxStack[i] = auxResult
				auxStack = append(auxStack[:i-1], auxStack[i-1+1:]...)
				auxStack = append(auxStack[:i-2], auxStack[i-2+1:]...)
				i = 0
			}
		}
	}

	if len(auxStack) == 1 {
		auxResult, err := strconv.ParseFloat(auxStack[0], 64)
		if err != nil {
			fmt.Printf("Error value1 as %s", auxStack[0])
		}
		c.resultString = auxStack[0]
		c.result = auxResult
		return c.result, nil
	}
	return -1, nil
}

func (c *RPN) CalculateExpression(str string) (float64, error) {
	tokens, err := Tokenize(str)
	if err != nil {
		return -1, err
	}
	c.ConvertToRPN(tokens)
	_, err = c.CalculateRPN()
	if err != nil {
		return -1, err
	}
	return c.result, nil
}