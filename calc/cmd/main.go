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
			return
		}

		calc := rpn.RPN{}
		calc.CalculateExpression(str)
		fmt.Println(calc.GetResult())
	}
}