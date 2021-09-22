package internal

import (
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

var r = regexp.MustCompile(`^[\d]+`)

func cUnique(str []string, params Params) []string {
	var count int
	var cUniq []string
	var last string
	for _, v := range str {
		if count == 0 {
			count++
			cUniq = append(cUniq, strconv.Itoa(count)+" "+v)
		} else {
			if (params.I && strings.EqualFold(last, v)) || (!params.I && last == v) { // flag -i
				count++
				i := r.FindStringSubmatch(cUniq[len(cUniq)-1])[0]
				split := cUniq[len(cUniq)-1][utf8.RuneCountInString(i)+1:]
				cUniq[len(cUniq)-1] = strconv.Itoa(count)+" "+split
			} else {
				count = 1
				cUniq = append(cUniq, strconv.Itoa(count)+" "+v)
			}
		}
		last = v
	}
	return cUniq
}

func dRepeated(str []string, params Params) []string {
	var dUniq []string
	var last string
	var repeat bool
	for _, v := range str {
		if (params.I && strings.EqualFold(last, v)) || (!params.I && last == v) { // flag -i
			if !repeat {
				dUniq = append(dUniq, v)
				repeat = true
			}
		} else {
			repeat = false
		}
		last = v
	}
	return dUniq
}

func uUnique(str []string, params Params) []string {
	var uUniq []string
	var last string
	var repeat bool
	var first bool
	for _, v := range str {
		if !((params.I && strings.EqualFold(last, v)) || (!params.I && last == v)) && first { // flag -i
			if !repeat {
				uUniq = append(uUniq, v)
				repeat = true
			}
		} else {
			repeat = false
		}
		last = v
		first = true
	}
	return uUniq
}

func Default(str []string, params Params) []string {
	var def []string
	var last string
	for _, v := range str {
		if (params.I && strings.EqualFold(last, v)) || (!params.I && last == v) { // flag -i
			continue
		}
		def = append(def, v)
		last = v
	}
	return def
}