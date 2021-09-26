package main

import (
	"hw/parser-calc/uniq/internal/uniq"
	"log"
)

func main() {
	params, err := uniq.SearchParams()
	if err != nil {
		log.Fatalln(err)
	}
	if params.NumChars < 0 || params.NumFields < 0 {
		log.Fatalln("Incorrect num/char field")
	}
	err = uniq.Uniq(params)
	if err != nil {
		log.Fatalln(err)
	}
}