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
		for sc.Scan() { // stop CTRL+D || control+D
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
	var line, newSlice []string
	line = append(line, strings.Split(str, "\n")...)

	// for tab \n in scanner
	if len(params.InputFile) == 0 {
		line = line[:len(line)-2]
	}

	// flags -f, -s
	if params.NumFields != 0 {
		newSlice = cutStrF(line, params.NumFields)
	}
	if params.NumChars != 0 {
		if params.NumFields != 0 {
			newSlice = cutCharS(newSlice, params.NumChars)
		} else {
			newSlice = cutCharS(line, params.NumChars)
		}
	}

	// flags -c, -d, -u
	if params.C {
		var cUniq []CountU
		if len(newSlice) == 0 {
			cUniq = cUnique(line, params)
		} else {
			cUniq = cUnique(newSlice, params)
		}
		return convertCount(line, cUniq)
	} else if params.D {
		var uniq []int
		if len(newSlice) == 0 {
			uniq = dRepeated(line, params)
		} else {
			uniq = dRepeated(newSlice, params)
		}
		return convert(line, uniq)
	} else if params.U {
		var uUniq []int
		if len(newSlice) == 0 {
			uUniq = uUnique(line, params)
		} else {
			uUniq = uUnique(newSlice, params)
		}
		return convert(line, uUniq)
	} else {
		var def []int
		if len(newSlice) == 0 {
			def = defaultF(line, params)
		} else {
			def = defaultF(newSlice, params)
		}
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