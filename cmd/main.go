package main

import (
	"hw/parser-calc/internal/tokenize"
	"log"
)

func main()  {
	log.Println("hello")
	str := "(1+(25/60)/*2)-3"
	tokens, err := tokenize.Tokenize(str)
	if err != nil {
		log.Println(err)
	}

	log.Println(tokens)
}