package main

import (
	"bufio"
	"fmt"
	"hw/parser-calc/calc/internal/rpn"
	"io"
	"log"
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
		_, err := calc.CalculateExpression(str)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(calc.GetResult())
	}
}