package internal

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strconv"
	"strings"
)

// Uniq starting the utility
func Uniq(params Params) error {
	var str, read string
	var err error

	if len(params.InputFile) == 0 {
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() { // stop CTRL+D || control+D
			read += sc.Text() +"\n"
		}
		str, err = startFlags(read, params)
		if err != nil {
			return err
		}
	} else {
		in, err := os.Open(params.InputFile)
		if err != nil {
			return err
		}

		defer in.Close()

		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(in)
		if err != nil {
			return err
		}
		str, err = startFlags(buf.String(), params)
		if err != nil {
			return err
		}
	}

	if len(params.OutputFile) == 0 {
		_, err := io.WriteString(os.Stdout, str)
		if err != nil {
			return err
		}
	} else {
		out, err := os.Create(params.OutputFile)

		if err != nil {
			return err
		}

		defer out.Close()

		_, err = out.WriteString(str)
		if err != nil {
			return err
		}
	}
	return nil
}

// See which flags are active
func startFlags(str string, params Params) (string, error) {
	var line, newSlice []string
	line = append(line, strings.Split(str, "\n")...)

	// for tab \n in scanner
	if len(params.InputFile) == 0 {
		line = line[:len(line)-1]
	}

	// flags -f, -s
	if params.NumFields != 0 {
		newSlice = cutStrF(line, params)
	}
	if params.NumChars != 0 {
		if params.NumFields != 0 {
			newSlice = cutCharS(newSlice, params)
		} else {
			newSlice = cutCharS(line, params)
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
		return convertCount(line, cUniq), nil
	} else if params.D {
		var uniq []int
		if len(newSlice) == 0 {
			uniq = dRepeated(line, params)
		} else {
			uniq = dRepeated(newSlice, params)
		}
		return convert(line, uniq), nil
	} else if params.U {
		var uUniq []int
		if len(newSlice) == 0 {
			uUniq = uUnique(line, params)
		} else {
			uUniq = uUnique(newSlice, params)
		}
		return convert(line, uUniq), nil
	} else {
		var def []int
		if len(newSlice) == 0 {
			def = defaultF(line, params)
		} else {
			def = defaultF(newSlice, params)
		}
		return convert(line, def), nil
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