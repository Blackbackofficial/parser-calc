package internal

import (
	"bufio"
	"io"
	"log"
	"os"
)

func Unic(params Params)  {

	var input, output *os.File
	defer input.Close()
	defer output.Close()

	if len(params.InputFile) == 0 {
		scan := bufio.NewScanner(io.Reader(os.Stdin))
		for scan.Scan() {
			str := scan.Text()

			// LOGIC

			log.Println(str)
			return
		}
		input = os.Stdin
	} else {
		input, err := os.Open(params.InputFile)
		log.Println(input)
		if err != nil {
			log.Fatalln(err)
		}
	}

}
