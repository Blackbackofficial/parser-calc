package internal

import (
	"errors"
	"flag"
)

type Params struct {
	C, D, U, I		      bool
	NumFields, NumChars   int64
	InputFile, OutputFile string
}

func SearchParams() (Params, error) {
	f := Params{}
	flag.BoolVar(&f.C,"c", false, "Count the number of occurrences of lines in the input. Print this number before the line, separated by a space.")
	flag.BoolVar(&f.D, "d", false, "Print only those lines that were repeated in the input data.")
	flag.BoolVar(&f.U, "u", false, "Print only those lines that have not been repeated in the input data.")
	flag.BoolVar(&f.I, "i", false, "Not case sensitive.")
	flag.Int64Var(&f.NumFields, "f", 0, "Ignore the first num_fields fields in the line. A field in a string is a non-empty character set separated by a space.")
	flag.Int64Var(&f.NumChars, "s", 0, "Ignore the first num_chars characters in the string. When used in conjunction with " +
		"the -f option, the first characters after the num_fields fields are counted (excluding the space separator after the last field).")
	flag.Parse()

	if !((f.C && !f.D && !f.U) || (!f.C && f.D && !f.U) || (!f.C && !f.D && f.U)) {
		flag.Usage()
		return f, errors.New("incorrect flag")
	}

	f.InputFile = flag.Arg(0) // Input file
	f.OutputFile = flag.Arg(1) // Output file
	return f, nil
}