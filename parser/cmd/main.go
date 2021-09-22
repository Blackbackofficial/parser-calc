package main

import (
	"hw/parser-calc/parser/internal"
	"log"
)

func main() {
	params, err := internal.SearchParams()
	if err != nil {
		log.Fatalln(err)
	}
	internal.Uniq(params)
}