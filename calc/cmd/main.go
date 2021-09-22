package main

import (
	"bufio"
	"fmt"
	"hw/parser-calc/calc/internal/rpn"
	"io"
	"os"
)

func main()  {
	scan := bufio.NewScanner(io.Reader(os.Stdin))
	for scan.Scan() {
		str := scan.Text()

		if str == "exit" {
			os.Exit(0)
		}

		calc := rpn.RPN{}
		calc.CalculateExpression(str)
		fmt.Println(calc.GetResult())
	}
}

//package main
//
//import (
//	"fmt"
//	"hw/parser-calc/calc/internal/rpn"
//	"math"
//	"strconv"
//	"strings"
//)
//// flags for convert to RPN
//var popLoop bool
//var firstItem bool
//
//type RpnGo struct {
//	expression      string
//	expressionStack []string
//	operatorStack []string
//	rpnExpression string
//	rpnStack      []string
//	result       float64
//	resultString string
//}
//
//func (r *RpnGo) SetExpression(expression string) {
//	expression = strings.ReplaceAll(expression, " ", "")
//	expression = strings.ReplaceAll(expression, " ", "")
//	expression = "(" + expression + ")"
//	expression = strings.TrimSpace(expression)
//	r.expression = expression
//}
//
//func (r *RpnGo) GetExpression() string {
//	return r.expression
//}
//
//func (r *RpnGo) GetResult() float64 {
//	return r.result
//}
//
//func (r *RpnGo) AppendRPNItem(item string) {
//	if item != "(" && item != ")" {
//		r.rpnExpression = r.rpnExpression + item + " "
//		r.rpnStack = append(r.rpnStack, item)
//	}
//}
//
//func (r *RpnGo) GetRPNExpression() string {
//	return r.rpnExpression
//}
//func (r *RpnGo) GetRPNStack() []string {
//	return r.rpnStack
//}
//
//func (r *RpnGo) AppendRPNOperatorItem(item string) {
//	r.operatorStack = append(r.operatorStack, item)
//}
//
//func (r *RpnGo) GetLastOperatorFromStack() string {
//	if len(r.operatorStack) > 0 {
//		return r.operatorStack[len(r.operatorStack)-1]
//	}
//	return ""
//}
//
//func (r *RpnGo) PopOperatorFromStack() []string {
//	if len(r.operatorStack) > 0 {
//		r.operatorStack = r.operatorStack[:len(r.operatorStack)-1]
//	}
//	return r.operatorStack
//}
//
//func (r *RpnGo) GetOperatorStackLength() int {
//	return len(r.operatorStack)
//}
//
//func (r *RpnGo) ConvertExpressionToStack() []string {
//	expression := r.expression
//	var list []string
//	tempStr := ""
//	isLastCharNumeric := false
//
//	for i := 0; i < len(expression); i++ {
//		tempChar := fmt.Sprintf("%c", expression[i])
//		if r.IsNumericString(tempChar) {
//			if isLastCharNumeric || len(list) == 0 {
//				tempStr = tempStr + tempChar
//			} else {
//				tempStr = tempStr + tempChar
//			}
//			isLastCharNumeric = true
//		} else {
//			if isLastCharNumeric {
//				list = append(list, tempStr)
//			}
//			tempStr = ""
//			list = append(list, tempChar)
//			isLastCharNumeric = false
//		}
//
//		if i == (len(expression) - 1) {
//			if r.IsNumericString(tempChar) {
//				list = append(list, tempStr)
//			} else {
//				list = append(list, tempChar)
//			}
//		}
//	}
//	r.expressionStack = list
//	return list
//}
//
//func (r *RpnGo) ConvertToRPN(expStack []string) string {
//	for i := range expStack {
//		item := expStack[i]
//		if r.IsOperator(item) {
//			if r.GetOperatorStackLength() == 0 || firstItem {
//				firstItem = false
//				r.AppendRPNOperatorItem(item)
//			} else {
//				if item == "(" || item == " " {
//					r.AppendRPNOperatorItem(item)
//					continue
//				}
//				if r.GetOperatorStackLength() > 0 && item == ")" {
//					for r.GetOperatorStackLength() > 0 && r.GetLastOperatorFromStack() != "(" {
//						r.AppendRPNItem(r.GetLastOperatorFromStack())
//						r.PopOperatorFromStack()
//					}
//
//					if r.GetOperatorStackLength() > 0 && r.GetLastOperatorFromStack() == "(" {
//						r.PopOperatorFromStack()
//					}
//					continue
//				}
//
//				for len(r.operatorStack) > 0 && (r.CheckPrecedence(item) <= r.CheckPrecedence(r.GetLastOperatorFromStack())) {
//					r.AppendRPNItem(r.GetLastOperatorFromStack())
//					r.PopOperatorFromStack()
//					popLoop = true
//				}
//				if popLoop {
//					r.operatorStack = append(r.operatorStack, item)
//					popLoop = false
//				} else if len(r.operatorStack) > 0 && (r.CheckPrecedence(item) > r.CheckPrecedence(r.GetLastOperatorFromStack())) {
//					r.operatorStack = append(r.operatorStack, item)
//				}
//			}
//		} else {
//			r.AppendRPNItem(item)
//		}
//	}
//
//	for len(r.operatorStack) > 0 {
//		r.AppendRPNItem(r.GetLastOperatorFromStack())
//		r.PopOperatorFromStack()
//	}
//	return r.rpnExpression
//}
//
//func (r *RpnGo) CheckPrecedence(item string) int {
//	switch item {
//	case "^":
//		return 40
//	case "*":
//		return 30
//	case "/":
//		return 30
//	case "+":
//		return 20
//	case "-":
//		return 20
//	}
//	return 0
//}
//
//func (r *RpnGo) GetIndexOfStringList(stringList []string, search string) int {
//	for i := 0; i < len(stringList); i++ {
//		if stringList[i] == search {
//			return i
//		}
//	}
//	return -1
//}
//
//func (r *RpnGo) IsNumericString(value string) bool {
//	if value == "0" || value == "1" || value == "2" || value == "3" || value == "4" || value == "5" || value == "6" || value == "7" || value == "8" || value == "9" || value == "." {
//		return true
//	}
//	return false
//}
//
//func (r *RpnGo) IsOperator(value string) bool {
//	if value == "^" || value == "*" || value == "/" || value == "+" || value == "-" || value == "=" || value == ")" || value == "(" {
//		return true
//	}
//	return false
//}
//
//func (r *RpnGo) SimpleCalculate(value1 float64, value2 float64, operator string) float64 {
//	auxResult := 0.0
//	switch operator {
//	case "^":
//		auxResult = math.Pow(value1, value2)
//	case "*":
//		auxResult = value1 * value2
//	case "/":
//		auxResult = value1 / value2
//	case "+":
//		auxResult = value1 + value2
//	case "-":
//		auxResult = value1 - value2
//	}
//	return auxResult
//}
//
//func (r *RpnGo) CalculateRPN() float64 {
//	auxStack := r.rpnStack
//	for len(auxStack) > 1 {
//		for i := 0; i < len(auxStack); i++ {
//			item := auxStack[i]
//			if r.IsOperator(item) {
//				value1, err := strconv.ParseFloat(auxStack[i-2], 64)
//				if err != nil {
//					fmt.Printf("Error value1 as %s", auxStack[i-2])
//				}
//				value2, err := strconv.ParseFloat(auxStack[i-1], 64)
//				if err != nil {
//					fmt.Printf("Error value1 as %s", auxStack[i-1])
//				}
//				resultCalc := r.SimpleCalculate(value1, value2, item)
//				auxResult := fmt.Sprintf("%f", resultCalc)
//
//				auxStack[i] = auxResult
//				auxStack = r.RemoveFromStackByIndex(auxStack, i-1)
//				auxStack = r.RemoveFromStackByIndex(auxStack, i-2)
//				i = 0
//			}
//		}
//	}
//	if len(auxStack) == 1 {
//		auxResult, err := strconv.ParseFloat(auxStack[0], 64)
//		if err != nil {
//			fmt.Printf("Error value1 as %s", auxStack[0])
//		}
//		r.resultString = auxStack[0]
//		r.result = auxResult
//		return r.result
//	}
//	return -1
//}
//
//func (r *RpnGo) CalculateExpression(str string) float64 {
//	tokens, err := rpn.Tokenize(str)
//	if err != nil {
//		fmt.Println(err)
//		return -1
//	}
//	r.ConvertToRPN(tokens)
//	r.CalculateRPN()
//	return r.result
//}
//
//func (r *RpnGo) RemoveFromStackByIndex(list []string, index int) []string {
//	return append(list[:index], list[index+1:]...)
//}
//
//func main() {
//	expression := "2*2*(1+1)"
//
//	rpn := RpnGo{}
//	rpn.CalculateExpression(expression)
//	fmt.Println(rpn.GetResult())
//}