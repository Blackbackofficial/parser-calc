package internal

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"strings"
)

func Uniq(params Params)  {
	var str string

	if len(params.InputFile) == 0 {
		// TODO: переделать
		scan := bufio.NewScanner(io.Reader(os.Stdin))
		for scan.Scan() {
			str = StartFlags(scan.Text(), params)
		}
	} else {
		in, err := os.Open(params.InputFile)
		if err != nil {
			log.Fatalln(err)
		}

		defer in.Close()

		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(in)
		if err != nil {
			log.Fatalln(err)
		}
		str = StartFlags(buf.String(), params)
	}

	if len(params.OutputFile) == 0 {
		_, err := io.WriteString(os.Stdout, str)
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		out, err := os.Create(params.OutputFile)

		if err != nil {
			log.Fatalln(err)
		}

		defer out.Close()

		_, err = out.WriteString(str)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func StartFlags(str string, params Params) string {
	var line []string
	line = append(line, strings.Split(str, "\n")...)

	if params.C {
		cUniq := cUnique(line, params)
		return convert(cUniq)
	} else if params.D {
		uniq := dRepeated(line, params)
		return convert(uniq)
	} else if params.U {
		uUniq := uUnique(line, params)
		return convert(uUniq)
	} else {
		def := Default(line, params)
		return convert(def)
	}
}

func convert(arr []string) string {
	var out string
	for _, v := range arr {
		out += v + "\n"
	}
	return out
}