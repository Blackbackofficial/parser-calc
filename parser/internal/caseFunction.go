package internal

import (
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

var r = regexp.MustCompile(`^[\d]+`)

func Unique(str []string, params Params) []string {
	var count int
	var uniq []string
	var last string
	for _, v := range str {
		if count == 0 {
			count++
			uniq = append(uniq, strconv.Itoa(count)+" "+v)
		} else {
			if (params.I && strings.EqualFold(last, v)) || (!params.I && last == v) { // flag -i
				count++
				i := r.FindStringSubmatch(uniq[len(uniq)-1])[0]
				split := uniq[len(uniq)-1][utf8.RuneCountInString(i)+1:]
				uniq[len(uniq)-1] = strconv.Itoa(count)+" "+split
			} else {
				count = 1
				uniq = append(uniq, strconv.Itoa(count)+" "+v)
			}
		}
		last = v
	}
	return uniq
}