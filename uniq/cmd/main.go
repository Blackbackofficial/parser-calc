package main

import (
	"hw/parser-calc/uniq/internal"
	"log"
)

func main() {
	params, err := internal.SearchParams()
	if err != nil {
		log.Fatalln(err)
	}
	if params.NumChars < 0 || params.NumFields < 0 {
		log.Fatalln("Incorrect num/char field")
	}
	err = internal.Uniq(params)
	if err != nil {
		log.Fatalln(err)
	}
}