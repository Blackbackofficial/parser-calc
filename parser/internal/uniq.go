package internal

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// Uniq starting the utility
func Uniq(params Params)  {
	var str, read string

	if len(params.InputFile) == 0 {
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			read += sc.Text() +"\n"
		}
		str = startFlags(read, params)
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
		str = startFlags(buf.String(), params)
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

// See which flags are active
func startFlags(str string, params Params) string {
	var line []string
	line = append(line, strings.Split(str, "\n")...)

	if params.C {
		cUniq := cUnique(line, params)
		return convertCount(line, cUniq)
	} else if params.D {
		uniq := dRepeated(line, params)
		return convert(line, uniq)
	} else if params.U {
		uUniq := uUnique(line, params)
		return convert(line, uUniq)
	} else {
		def := defaultF(line, params)
		return convert(line, def)
	}
}

// String summation
func convert(str []string, arr []int) string {
	var out string
	for _, v := range arr {
		out += str[v] + "\n"
	}
	return out
}

// String summation with position
func convertCount(str []string, arr []CountU) string {
	var out string
	for _, v := range arr {
		out += strconv.Itoa(v.count)+" "+str[v.num]+"\n"
	}
	return out
}